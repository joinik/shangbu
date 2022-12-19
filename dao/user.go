package dao

import (
	"context"
	"go_ctry/model"

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

// GetProfileById 根据id 查询用户详情
func (dao *UserDao) GetUserProfileById(uId uint) (userProfile *model.UserProfile,  err error) {
	err = dao.DB.Model(&model.UserProfile{}).Where("user_id=?", uId).First(&userProfile).Error
    return 
}


// UpdateUserProfileById 根据id 更新用户详情
func (dao *UserDao) UpdateUserProfileById(uId uint, user *map[string]interface{}) (err error) {
	return dao.DB.Model(&model.UserProfile{}).Where("user_id=?", uId).Updates(&user).Error

}





// ExitOrNotByUserName 根据手机号 查询用户是否存在
func (dao *UserDao) ExitOrNotByUserName(name string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name=? or mobile=?", name,name).
		Count(&count).Error

    if count == 0{
        return nil, false, err
    }

    err = dao.DB.Model(&model.User{}).Where("user_name=? or mobile=?", name,name).
		First(&user).Error
	if err != nil {
		return nil, false, err
	}
	return user, true, nil
}

// ExitOrNotByPhone 根据手机号 查询用户是否存在
func (dao *UserDao) ExitOrNotByPhone(mobile string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("mobile=?", mobile).
		Count(&count).Error

	if count == 0 {
		return nil, false, err
	}

	err = dao.DB.Model(&model.User{}).Where("mobile=?", mobile).
		First(&user).Error
	if err != nil {
		return nil, false, err
	}
	return user, true, nil
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).Create(&user).Error
	return
}


