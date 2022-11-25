package api

import (
	"go_ctry/service"

	"github.com/gin-gonic/gin"
)

// 注册用户 api
func RegisterUser(c *gin.Context) {
	userService := service.Userservice{}
	if err := c.ShouldBind(&userService); err == nil {
		res := userService.Register(c.Request.Context())
		c.JSON(200, res)
	} else {

	}

}

// 查询用户详情
func GetUserProfile(c *gin.Context) {
	userService := service.Userservice{}
	if err := c.ShouldBind(&userService); err == nil {
		res := userService.GetUserProfile(c.Request.Context(), c.Param("id") )
		c.JSON(200, res)
	}
}
