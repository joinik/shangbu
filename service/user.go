package service

import (
	"context"
	"go_ctry/model"
	"go_ctry/pkg/e"
	util "go_ctry/pkg/utils"
	"go_ctry/serializer"
	"time"

	"github.com/google/uuid"

	// "time"
	logging "github.com/sirupsen/logrus"
)

// 定义 Userservice

type Userservice struct {
	UserName  string `form:"user_name" json:"user_name"` // 用户名
	Mobile    string `form:"mobile" json:"mobile"`       // 手机号
	Introduce string `form:"introduce" json:"introduce"`
	Business  uint   `form:"business" json:"business"` // 商家标识
	Gender    string `form:"gender" json:"gender"`
	Age       uint   `form:"age" json:"age"`
	Email     string `form:"email" json:"email"` // 邮箱
	Vcode     string `form:"vcode" json:"vcode"` // 手机验证码
}

// 用户注册服务
func (service *Userservice) Register(ctx context.Context) serializer.Response {
	code := e.SUCCESS

	var count int64
	dbClient := model.NewDBClient(ctx)

	// 1. 判断短信 验证码
	// if service.Vcode

	// 2. 根据手机号查询 用户 是否存在
	err := dbClient.Model(&model.User{}).Where("mobile=?", service.Mobile).
		Count(&count).Error

	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	var user = &model.User{}
	// 如果 没有用户名 则进行数据库 创建
	if count == 0 {

		// 生成 性别
		var gender model.MyGender
		gender, err = model.NewGender("WOMAN")
		if err != nil {
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    "性别错误",
			}
		}

		// var userProfile  model.UserProfile{}

		uuid := uuid.New()
		key := uuid.String()
		// 定义model 变量
		user = &model.User{
			UserName:   service.Mobile[:4] + key[:6],
			Mobile:     service.Mobile,
			Introduce:  service.Introduce,
			Business:   service.Business,
			Last_login: time.Now(),
			UserProfile: model.UserProfile{
				UserGender: gender,
				Age:        0,
				Email:      "",
			},
		}

		// 创建用户数据
		err := dbClient.Model(&model.User{}).Create(&user).Error

		if err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}

	} else {
		// 用户 已存在的情况下

		err = dbClient.Model(&model.User{}).Where("mobile=?", service.Mobile).
			First(&user).Error

		if err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}

	token, refreshToken, err := util.MyGenerateToken(user.ID, user.UserName, 0, true)
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token, RefreshToken: refreshToken},
		Msg:    e.GetMsg(code),
	}

}

// 修改 用户性情
func (service *Userservice) UpdateUserInfo(ctx context.Context, uid uint) serializer.Response {
	code := e.SUCCESS

	// var count int64
	dbClient := model.NewDBClient(ctx)
	var user *model.User
	// 判断用户个人数据是否为空

	if service.Mobile == "" || service.Age == 0 || service.Introduce == "" ||
		service.Gender == "" || service.Email == "" {
		return serializer.Response{
			Status: code,
			Msg:    "数据不完整",
		}
	}

	// 查询用户
	err := dbClient.Model(&model.User{}).Where("id=?", uid).
		First(&user).Error

	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	gender, err := model.NewGender(service.Gender)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    "性别错误",
		}
	}


	var mapUser = map[string]interface{}{"user_name": service.UserName,
		"mobile":    service.Mobile,
		"Introduce": service.Introduce,
	}

	// user.UserName = service.UserName
	// user.Mobile = service.Mobile
	// user.Introduce = service.Introduce
	// user.UserProfile.UserGender = gender
	// user.UserProfile.Age = service.Age
	// user.UserProfile.Email = service.Email

	err = dbClient.Model(&model.User{}).Where("id=?", uid).Updates(mapUser).Error

	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 更新用户详情

	userProfile := &model.UserProfile{
		UserGender: gender,
		Age:        service.Age,
		Email:      service.Email,
	}

	err = dbClient.Model(&model.UserProfile{}).Where("user_id=?", uid).Updates(&userProfile).Error

	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   nil,
	}

}

// 查询 用户详情
func (service *Userservice) GetUserProfile(ctx context.Context, userid uint) serializer.Response {
	code := e.SUCCESS

	dbClient := model.NewDBClient(ctx)

	var userProfile *model.UserProfile
	// 根据 用户id  查询  用户详情

	err := dbClient.Model(&model.UserProfile{}).Where("user_id=?", userid).
		First(&userProfile).Error

	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUserProfile(userProfile),
	}

}
