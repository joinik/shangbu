package api

import (
	"go_ctry/service"

	"github.com/gin-gonic/gin"
)

// 注册用户 api
func RegisterUser(c *gin.Context) {
	userService := service.Userservice{}
	if err := c.ShouldBind(&userService); err == nil {
		res := userService.Registers(c.Request.Context())
		c.JSON(200, res)
	} 
	// else {
	// 	c.JSON(400, ErrorResponse(err))
	// 	util.LogrusObj.Infoln(err)
	// }


}

// 查询用户详情
func GetUserProfile(c *gin.Context) {
	userService := service.Userservice{}
	if err := c.ShouldBind(&userService); err == nil {
		res := userService.GetUserProfile(c.Request.Context(), c.Param("id") )
		c.JSON(200, res)
	}
}
