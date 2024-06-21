package v1

// import (
// 	"go_ctry/service"

// 	"github.com/gin-gonic/gin"
// )

// // 获取地址api
// func GetArea(c *gin.Context) {
//     areaService := service.AreaService{}
   
//     if err := c.ShouldBind(&areaService); err == nil {
//          aid := c.Param("aid")
//         res := areaService.Get(c.Request.Context(), aid)
//         c.JSON(200, res)
//     }
// }