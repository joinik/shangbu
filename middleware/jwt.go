package middleware

import (
	"go_ctry/pkg/e"
	util "go_ctry/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// jwt  token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		var claims *util.Claims
		var err error
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.ErrorAuth
		} else {
			claims, err = util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("claim", claims)
		c.Next()
	}
}

// JWTAdmin token验证中间件
func JWTAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.InvalidParams
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			} else if claims.Authority == 0 {
				code = e.ErrorAuthInsufficientAuthority
			}
		}
		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
