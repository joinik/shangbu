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
		// v1.GET("area/:aid", api.GetArea)

		// 用户登录 即注册
		v1.POST("login", api.LoginUser)
		// 刷新业务token
		v1.GET("updateToken", api.UpdateToken)
		// 评论查询
		v1.POST("comments", api.GetComment)



		// 需要登录保护的
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			// // 修改用户详情
			// authed.PUT("userprofile", api.UpdateUserInfo)
			// // 上传用户头像
			// authed.POST("uploadAvatar", api.UploadAvatar)

			// --------文章相关api----------
			// 文章创建
			authed.POST("createArt", api.CreateArt)
			// 文章更新
			authed.PUT("updateArt", api.UpdateArt)
			// 文章点赞 点踩 记录
			authed.POST("likeArt", api.ArtRecord)

			// --------评论相关api----------
			// 评论创建 
			authed.POST("createCom", api.CreateComment)
			// 评论点赞 
			authed.POST("createComRecord", api.CreateComRecord)
		}

	}
	return r

}
