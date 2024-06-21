package dao

import (
	"context"
	"go_ctry/model"
	"time"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

// GetUserById 根据id 查询用户
func (dao *UserDao) GetUserById(uId uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", uId).First(&user).Error
	return
}

// UpdateUserById 根据 id 更新用户信息
func (dao *UserDao) UpdateUserById(uId uint, user *model.User) (err error) {
	return dao.DB.Model(&model.User{}).Where("id=?", uId).Updates(&user).Error

}

// ExitOrNotByUserName 根据手机号 查询用户是否存在
func (dao *UserDao) ExitOrNotByUserName(name string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name=? or mobile=?", name, name).
		Count(&count).Error

	if count == 0 {
		return nil, false, err
	}

	err = dao.DB.Model(&model.User{}).Where("user_name=? or mobile=?", name, name).
		First(&user).Error
	if err != nil {
		return nil, false, err
	}
	return user, true, nil
}

// ExitOrNotByPhone 根据手机号 查询用户是否存在
func (dao *UserDao) ExitOrNotByPhone(mobile string) (user *model.User, err error) {

	// DB获取用户
	err = dao.DB.Model(&model.User{}).Where("mobile=?", mobile).
		First(&user).Error
	// 判断用户是否存在
	if err == nil {
		// 更新用户登录时间
		dao.DB.Model(&user).Update("last_login", time.Now())
		return user, nil
	} else {
		// 用户不存在
		user = &model.User{
			Mobile:   mobile,
			UserName: mobile,
		}
		err = dao.DB.Model(&model.User{}).Create(&user).Error
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}
