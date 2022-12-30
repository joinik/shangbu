package v1

import (
	util "go_ctry/pkg/utils"
	"go_ctry/service"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var comService service.CommentService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&comService); err == nil {
		res := comService.Post(c.Request.Context(), claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		// util.LogrusObj.Infoln(err)
	}
}

func GetComment(c *gin.Context) {
	var comService service.CommentService
	if err := c.ShouldBind(&comService); err == nil {
		res := comService.Get(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		// util.LogrusObj.Infoln(err)
	}
}

func CreateComRecord(c *gin.Context) {
	var comRecordService service.ComRecordService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&comRecordService); err == nil {
		res := comRecordService.PostComRecord(c.Request.Context(), claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		// util.LogrusObj.Infoln(err)
	}
}
