package v1

import (
	util "go_ctry/pkg/utils"
	"go_ctry/service"

	"github.com/gin-gonic/gin"
)

// func UploadArtPhoto(c *gin.Context) {
// 	var artService service.ArticleService

// 	form, _ := c.MultipartForm()
// 	files := form.File["upload[]"]

// 	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
// 	if err := c.ShouldBind(&artService); err != nil {
// 		res := artService.UploadArtPhoto(c.Request.Context(), claim.ID, files)
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(400, ErrorResponse(err))
// 		// util.LogrusObj.Infoln(err)
// 	}

// }

func CreateArt(c *gin.Context) {
	artService := service.ArticleService{}
	form, _ := c.MultipartForm()
	files := form.File["upload"]

	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&artService); err == nil {
		res := artService.Create(c.Request.Context(), claim.ID, files)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		// util.LogrusObj.Infoln(err)
	}
}

func GetArt(c *gin.Context) {
	artService := service.ArticleService{}
	if err:= c.ShouldBind(&artService); err == nil {
		artid := c.Param("artid")
		res := artService.GetArtByArtID(c.Request.Context(), artid)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		// util.LogrusObj.Infoln(err)
	}
}

func GetArtContent(c *gin.Context) {
	artService := service.ArticleService{}
	if err:= c.ShouldBind(&artService); err == nil {
		artid := c.Param("artid")
		res := artService.GetContentByArtID(c.Request.Context(), artid)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		// util.LogrusObj.Infoln(err)
	}
}



