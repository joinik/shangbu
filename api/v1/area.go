package api

import (
	"go_ctry/service"

	"github.com/gin-gonic/gin"
)




// 获取地址api
func GetArea(c *gin.Context) {
    areaService := service.AreaService{}
    if err := c.ShouldBind(&areaService); err == nil {
        res := areaService.Get(c.Request.Context(), "788930")
        c.JSON(200, res)
    }
}