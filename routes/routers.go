package routes

import (
	api "go_ctry/api/v1"
	"go_ctry/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	// store := cookie.NewStore([]byte("something-very-sectet"))
	// r.Use(middleware.Cors())
	// r.Use(sessions.Session("mysession", store))

	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(middleware.Cors())
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		// 地址查询
		v1.GET("area/:aid", api.GetArea)
		// 用户注册
		v1.POST("register", api.RegisterUser)

		// 用户登录
		v1.POST("login", api.LoginUser)

		// 需要登录保护的
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			// 查询用户详情
			authed.GET("userProfile", api.GetUserProfile)
			// 修改用户详情
			authed.PUT("userprofile", api.UpdateUserInfo)
		}

	}
	return r

}
