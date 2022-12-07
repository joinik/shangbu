package routes

import (
	api "go_ctry/api/v1"
	"go_ctry/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go_ctry/docs" // 这里需要引入本地已生成文档
	


)

// 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	// store := cookie.NewStore([]byte("something-very-sectet"))
	// r.Use(middleware.Cors())
	// r.Use(sessions.Session("mysession", store))


	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(middleware.Cors())
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		// 地址查询
		v1.GET("area", api.GetArea)
		// 用户注册
		v1.POST("register", api.RegisterUser)

		v1.GET("userProfile/:id", api.GetUserProfile)
	}
	return r

}
