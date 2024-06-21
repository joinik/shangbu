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

/*
 * @Description: 插入收藏ID到用户收藏表
 * @param uId	用户ID
 * @param cid	收藏ID切片
 * @return err  错误信息
 */
func (dao *UserDao) InsertCollectById(cid []uint, uId uint) error {
	// var cols []map[string]interface{}
	// for _, v := range cid {
	// 	cols = append(cols, map[string]interface{}{"UserID": uId, "CollectID": v})
	// }

	var cols []model.UserCollect
	for _, v := range cid {
		cols = append(cols, model.UserCollect{UserID: uId, CollectID: v})
	}

	return dao.DB.Model(&model.UserCollect{}).Create(&cols).Error
}

// SelectCollectById 根据用户ID查询用户及其收藏信息。
// 该方法使用GORM库来查询数据库，首先根据uId获取用户信息，然后预加载用户的收藏信息。
// 参数:
//
//	uId - 用户的唯一标识符。
//
// 返回值:
//
//	*model.User - 查询到的用户对象。
//	error - 如果查询过程中出现错误，则返回错误对象；否则返回nil。
func (dao *UserDao) SelectCollectById(uId uint) (user *model.User, err error) {
	// 使用GORM的Model方法指定查询的模型，这里是model.User。
	// Preload方法用于预加载关联的数据，这里是用户的收藏信息。
	// First方法根据uId查询数据库，将查询结果存储到user变量中。
	err = dao.DB.Model(&model.User{}).Preload("UserCollects").First(&user, uId).Error
	// 如果查询过程中出现错误，返回nil和错误对象。
	if err != nil {
		return nil, err
	}
	// 如果查询成功，返回查询到的用户对象和nil。
	return user, nil
}
