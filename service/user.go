package service

import (
	"context"
	"go_ctry/dao"
	"go_ctry/model"
	"go_ctry/pkg/e"
	util "go_ctry/pkg/utils"
	"go_ctry/serializer"
	"mime/multipart"
	"regexp"
	"time"

	"github.com/google/uuid"

	// "time"
	logging "github.com/sirupsen/logrus"
)

// 定义 Userservice

type Userservice struct {
	UserName  string `form:"user_name" json:"user_name"` // 用户名
	Mobile    string `form:"mobile" json:"mobile"`       // 手机号
	Pwd       string `form:"pwd" json:"pwd"`
	Introduce string `form:"introduce" json:"introduce"` // 简介
	Business  uint   `form:"business" json:"business"`   // 商家标识
	Gender    string `form:"gender" json:"gender"`
	Age       uint   `form:"age" json:"age"`
	Email     string `form:"email" json:"email"`     // 邮箱
	Vcode     string `form:"vcode" json:"vcode"`     // 手机验证码
	options   string `form:"options" json:"options"` // 登录方式 1 手机登录的方式  2 账号登录方式
}

// 用户注册服务
func (service *Userservice) Register(ctx context.Context) serializer.Response {
	code := e.SUCCESS

	// 生成 userDao 用于数据库操作
	userDao := dao.NewUserDao(ctx)

	// 1. 判断短信 验证码
	// if service.Vcode

	// 2. 根据手机号查询 用户 是否存在

	// 2.1 正则判断  手机号
	if rest, _ := regexp.Match(`^1[3456789]\d{9}$`, []byte(service.Mobile)); !rest {
		return serializer.Response{
			Status: code,
			Msg:    "数据不正确",
		}
	}

	user, exit, err := userDao.ExitOrNotByPhone(service.Mobile)

	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 判断手机号是否已经注册
	if exit {
		logging.Info(err)
		code = e.ErrorExistPhone
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	} else {
		// 如果 没有用户名 则进行数据库 创建
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
		err := userDao.CreateUser(user)
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

// 用户登录
func (service *Userservice) Login(ctx context.Context) serializer.Response {
	code := e.SUCCESS

	userDao := dao.NewUserDao(ctx)

	// 1. 判断 登录 方式
	if service.options == "1" {
		// 手机号 + 验证码 登录

		// 1. 判断 手机号 合法
		if rest, _ := regexp.Match(`^1[3456789]\d{9}$`, []byte(service.Mobile)); !rest {
			return serializer.Response{
				Status: code,
				Msg:    "数据不正确",
			}
		}

		// 2. 验证 短信验证码

		// 3. 根据手机号 查询用户 信息

		user, exist, err := userDao.ExitOrNotByPhone(service.Mobile)

		if !exist {
			logging.Info(err)
			code = e.ErrorUserNotFound
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		// 4. 返回响应
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
	} else if service.options == "2" {
		// 账号 + 密码

		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}

	} else {
		// 非法方式登录
		return serializer.Response{
			Status: e.InvalidParams,
			Msg:    e.GetMsg(code),
		}
	}

}

// 设置密码
func (service *Userservice) UpdatePwd(ctx context.Context, uid uint) serializer.Response {

	code := e.SUCCESS

	userDao := dao.NewUserDao(ctx)

	// 1.判断密码 合法性
	if rest, _ := regexp.Match(`^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{6,20}$`, []byte(service.Pwd)); !rest {
		return serializer.Response{
			Status: code,
			Msg:    "数据不正确",
		}
	}
	// 2. 密码加密
	user, err := userDao.GetUserById(uid)

	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	user.SetPassword(service.Pwd)

	// 3. 保存数据库
	err = userDao.UpdateUserById(uid, user)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 返回响应
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}

// 修改 用户详情
func (service *Userservice) UpdateUserInfo(ctx context.Context, uid uint) serializer.Response {
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	// 判断用户个人数据是否为空

	if service.Mobile == "" || service.Age == 0 || service.Introduce == "" ||
		service.Gender == "" || service.Email == "" {
		return serializer.Response{
			Status: code,
			Msg:    "数据不完整",
		}
	}

	if rest, _ := regexp.Match(`^1[3456789]\d{9}$`, []byte(service.Mobile)); !rest {
		return serializer.Response{
			Status: code,
			Msg:    "数据不正确",
		}
	} else if rest, _ = regexp.Match(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`, []byte(service.Email)); !rest {
		return serializer.Response{
			Status: code,
			Msg:    "数据不正确",
		}
	}

	// 查询用户

	user, err := userDao.GetUserById(uid)

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

	// var mapUser = map[string]interface{}{"user_name": service.UserName,
	// 	"mobile":    service.Mobile,
	// 	"Introduce": service.Introduce,
	// }

	_, exist, err := userDao.ExitOrNotByPhone(service.Mobile)

	if exist {
		logging.Info(err)
		code = e.ErrorExistPhone
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	user.UserName = service.UserName
	user.Mobile = service.Mobile
	user.Introduce = service.Introduce
	// user.UserProfile.UserGender = gender
	// user.UserProfile.Age = service.Age
	// user.UserProfile.Email = service.Email

	err = userDao.UpdateUserById(uid, user)

	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 更新用户详情
	err = userDao.UpdateUserProfileById(uid, &map[string]interface{}{"user_gender": gender, "age": service.Age, "emial": service.Email})

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

	userDao := dao.NewUserDao(ctx)
	// 根据 用户id  查询  用户详情

	userProfile, err := userDao.GetUserProfileById(userid)

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

// 用户头像上传
func (service *Userservice) UploadHead(ctx context.Context, userid uint, file multipart.File, fileSize int64) serializer.Response {
	code := e.SUCCESS
	var err error
	
	path, err := UploadToQiNiu(file, fileSize)
	if err != nil {
		code = e.ErrorUploadFile
		return serializer.Response{
			Status: code,
			Error:  path,
			Data:   e.GetMsg(code),
		}
	}
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userid)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	user.Avatar = path
	err = userDao.UpdateUserById(userid, user)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}

}
