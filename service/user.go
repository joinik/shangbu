package service

import (
	"context"
	"go_ctry/dao"
	"go_ctry/pkg/e"
	util "go_ctry/pkg/utils"
	"go_ctry/serializer"
	"regexp"

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
	Options   string `form:"options" json:"options"` // 登录方式 1 手机登录的方式  2 账号登录方式
}

// 用户登录
func (service *Userservice) Login(ctx context.Context) serializer.Response {
	code := e.SUCCESS

	userDao := dao.NewUserDao(ctx)

	// 1. 判断 登录 方式
	if service.Options == "1" {
		// 手机号 + 验证码 登录

		// 1. 判断 手机号 合法
		if rest, _ := regexp.Match(`^1[3-9]\d{9}$`, []byte(service.Mobile)); !rest {
			return serializer.Response{
				Status: code,
				Msg:    "数据不正确",
			}
		}

		//TODO 2. 验证 短信验证码

		// 3. 根据手机号 查询用户 信息

		user, err := userDao.ExitOrNotByPhone(service.Mobile)
		if err != nil {
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}

		// 4. 返回响应
		token, refreshToken, err := util.MyGenerateToken(user.ID, user.Mobile, 0, true)
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
	} else {
		// 非法方式登录
		return serializer.Response{
			Status: e.InvalidParams,
			Msg:    e.GetMsg(code),
		}
	}

}

// 查询 用户详情
// func (service *Userservice) GetUserProfile(ctx context.Context, userid uint) serializer.Response {
// 	code := e.SUCCESS

// 	userDao := dao.NewUserDao(ctx)
// 	// 根据 用户id  查询  用户详情

// 	userProfile, err := userDao.GetUserProfileById(userid)

// 	if err != nil {
// 		code = e.ErrorDatabase
// 		return serializer.Response{
// 			Status: code,
// 			Msg:    e.GetMsg(code),
// 		}
// 	}

// 	return serializer.Response{
// 		Status: code,
// 		Msg:    e.GetMsg(code),
// 		Data:   serializer.BuildUserProfile(userProfile),
// 	}

// }

// 用户头像上传
// func (service *Userservice) UploadHead(ctx context.Context, userid uint, file multipart.File, fileSize int64) serializer.Response {
// 	code := e.SUCCESS
// 	var err error

// 	path, err := UploadToQiNiu(file, fileSize)
// 	if err != nil {
// 		code = e.ErrorUploadFile
// 		return serializer.Response{
// 			Status: code,
// 			Error:  path,
// 			Data:   e.GetMsg(code),
// 		}
// 	}
// 	userDao := dao.NewUserDao(ctx)
// 	user, err := userDao.GetUserById(userid)
// 	if err != nil {
// 		code = e.ErrorDatabase
// 		return serializer.Response{
// 			Status: code,
// 			Msg:    e.GetMsg(code),
// 			Error:  err.Error(),
// 		}
// 	}

// 	user.Avatar = path
// 	err = userDao.UpdateUserById(userid, user)
// 	if err != nil {
// 		logging.Info(err)
// 		code = e.ErrorDatabase
// 		return serializer.Response{
// 			Status: code,
// 			Msg:    e.GetMsg(code),
// 			Error:  err.Error(),
// 		}
// 	}

// 	return serializer.Response{
// 		Status: code,
// 		Data:   serializer.BuildUser(user),
// 		Msg:    e.GetMsg(code),
// 	}

<<<<<<< HEAD
}

// UpdateToken 更新token
func (service *Userservice) UpdateToken(token string) serializer.Response {

	code := e.SUCCESS

	// 根据refreshtoken 刷新业务token
	calim, err := util.ParseToken(token)

	if err != nil || !calim.Isrefresh {
		code = e.ErrorAuthCheckTokenFail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	token, _, err = util.MyGenerateToken(calim.ID, calim.Username, calim.Authority, false)

	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   token,
	}

}
=======
// }
>>>>>>> 95f4958 (refactor: 🚴重构用户登录模块)
