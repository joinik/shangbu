package v1

import (
	util "go_ctry/pkg/utils"
	"go_ctry/service"

	"github.com/gin-gonic/gin"
)

// 注册用户 api
func RegisterUser(c *gin.Context) {
	var userService service.Userservice
	if err := c.ShouldBind(&userService); err == nil {
		res := userService.Register(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		// util.LogrusObj.Infoln(err)
	}

}

// 登录用户 Api
func LoginUser(c *gin.Context) {
	var loginUserServer service.Userservice
	if err := c.ShouldBind(&loginUserServer); err == nil {
		res := loginUserServer.Login(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}


func UpdatePwd(c *gin.Context) {
	var updatePwdService service.Userservice
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&updatePwdService); err == nil {
		res := updatePwdService.UpdatePwd(c.Request.Context(), claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		// util.LogrusObj.Infoln(err)
	}
}



// 修改用户信息
func UpdateUserInfo(c *gin.Context) {
	var userService service.Userservice
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userService); err == nil {
		res := userService.UpdateUserInfo(c.Request.Context(), claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		// util.LogrusObj.Infoln(err)
	}

}

// 查询用户详情
func GetUserProfile(c *gin.Context) {
	var userService service.Userservice
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userService); err == nil {
		res := userService.GetUserProfile(c.Request.Context(), claims.ID)
		c.JSON(200, res)
	}
}

// 上传用户头像api
func UploadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	uploadHeadService := service.Userservice{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&uploadHeadService); err == nil {
		res := uploadHeadService.UploadHead(c.Request.Context(), claim.ID, file, fileSize)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}

}
