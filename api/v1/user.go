package v1

import (
	"go_ctry/service"

	"github.com/gin-gonic/gin"
)

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

// 修改用户信息

// 查询用户详情

// 上传用户头像api
// func UploadAvatar(c *gin.Context) {
// 	file, fileHeader, _ := c.Request.FormFile("file")
// 	fileSize := fileHeader.Size
// 	uploadHeadService := service.Userservice{}
// 	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
// 	if err := c.ShouldBind(&uploadHeadService); err == nil {
// 		res := uploadHeadService.UploadHead(c.Request.Context(), claim.ID, file, fileSize)
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(400, ErrorResponse(err))
// 		util.LogrusObj.Infoln(err)
// 	}

// }
