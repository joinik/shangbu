package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               // 请求方法
		origin := c.Request.Header.Get("Origin") // 请求头部
		var HeaderKeys []string
		for k := range c.Request.Header {
			HeaderKeys = append(HeaderKeys, k) // 添加到切片
		}
		headerStr := strings.Join(HeaderKeys, ", ") // str 拼接
		if headerStr != "" {
			headerStr = fmt.Sprintf("acess-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "acess-control-allow-origin, access-control-allow-headers"

		}

		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			// 这是允许访问所有域
			c.Header("Access-Control-Allow-Origin", "*")
			//  服务器支持所有跨域请求方法，为了避免浏览器次请求的多次“预检”请求
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")

			// header 类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRT-Token, token, session, X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许跨域设置
			// 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			c.Header("Access-Control-Max-Age", "172800")          // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false") //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")             // 设置返回格式是json

		}

        // 放行所有的OPTIONS方法
        if method == "OPTIONS" {
            c.JSON(http.StatusOK, "Options Request!")
        }

        // 放行
        c.Next() 

	}
}
