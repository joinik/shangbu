package service

import (
	"context"
	"fmt"
	"go_ctry/model"
	"go_ctry/pkg/e"
	"go_ctry/serializer"
	"strconv"
	"time"

	logging "github.com/sirupsen/logrus"
)

// 定义 Userservice
type Userservice struct {
	UserName  string `json:"username" ` // 用户名
	Mobile    string `json:"phone" `    // 手机号
	Introduce string `json:"introduce"`
	Business  uint   `json:"business"` // 商家标识
	Gender    string `json:"gender"`
	Age       uint   `json:"age"`
	Email     string `json:"email"` // 邮箱
	Vcode     string `json:"vcode"` // 手机验证码
}

// 根据用户名查询

// func (service *Userservice) ExistOrNotByUserName(ctx context.Context) serializer.Response {
// 	var count int64

// 	dbClient := model.NewDBClient(ctx)

// 	_ = dbClient.Model(&model.User{}).Where("user_name=?", service.UserName).
// 		Count(&count).Error

// 	return serializer.Response{
// 		Status: e.SUCCESS,
// 		Msg:    strconv.FormatInt(count, 10),
// 	}

// }

// 用户注册服务
func (service *Userservice) Register(ctx context.Context) serializer.Response {
	code := e.SUCCESS

	// 数据判断
	if service.Mobile ==“” && service.  {
		
	}

	var count int64
	dbClient := model.NewDBClient(ctx)

	// 查询用户名
	err := dbClient.Model(&model.User{}).Where("user_name=?", service.Mobile).
		Count(&count).Error

	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 如果 没有用户名 则进行数据库 创建
	if count == 0 {

		// 生成 性别
		var gender model.MyGender
		if service.Gender == "" {

			gender, _ = model.NewGender("Man")
		} else {
			gender, _ = model.NewGender("Man")
		}

		// var userProfile  model.UserProfile{}

		// 定义model 变量
		// user := &model.User{
		// 	UserName:   service.Mobile,
		// 	Mobile:     service.Mobile,
		// 	Introduce:  service.Introduce,
		// 	Business:   service.Business,
		// 	Last_login: time.Now(),
		// 	UserProfile: model.UserProfile{
		// 		UserGender: gender,
		// 		Age:        service.Age,
		// 		Email:      service.Email,
		// 	},
		// }

		user := &model.User{
			UserName:   service.Mobile,
			Mobile:     service.Mobile,
			Introduce:  service.Introduce,
			Business:   service.Business,
			Last_login: time.Now(),
		}

		// 创建用户数据
		err := dbClient.Model(&model.User{}).Create(&user).Error

		// fmt.Println(user.UserProfile.UserID)
		// fmt.Println("_________________>>>>>>>>>>>")

		if err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}

		//  用户详情 创建
		userProfile := &model.UserProfile{
			UserID:     user.ID,
			UserGender: gender,
			Age:        service.Age,
			Email:      service.Email,
		}

		// 创建用户详情
		err = dbClient.Model(&model.UserProfile{}).Create(&userProfile).Error

		if err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}

		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 用户 已存在的情况下

	user := &model.User{}
	err = dbClient.Model(&model.User{}).Where("mobile=?", service.Mobile).
		First(user).Error

	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// 查询 用户详情
func (service *Userservice) GetUserProfile(ctx context.Context, userid string) serializer.Response {
	code := e.SUCCESS

	dbClient := model.NewDBClient(ctx)

	uid, _ := strconv.Atoi(userid)

	var userProfile *model.UserProfile
	// 根据 用户id  查询  用户详情
	err := dbClient.Model(&model.UserProfile{}).Where("user_id=?", uint(uid)).
		First(&userProfile).Error

	fmt.Printf("%v", userProfile)
	fmt.Printf("序列化之前》》》》》》》》》")
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	fmt.Println("用户 没有问题")
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUserProfile(userProfile),
	}

}
