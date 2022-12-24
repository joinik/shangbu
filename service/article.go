package service

import (
	"context"
	"go_ctry/dao"
	"go_ctry/model"
	"go_ctry/pkg/e"
	"go_ctry/serializer"
	"mime/multipart"
	"strconv"
)

type ArticleService struct {
	Title   string `form:"title" json:"title"`
	Cate    uint   `form:"cate" json:"cate"`
	Area    uint   `form:"area" json:"area"`
	Content string `form:"content" json:"content"`
}

func (service *ArticleService) UploadArtPhoto(ctx context.Context, artid uint, files []*multipart.FileHeader) serializer.Response {
	code := e.SUCCESS
	// var err error

	// var cover []map[string]string
	for _, file := range files {
		println("----------->", file)

		// path, err := UploadToQiNiu(*file, fileSize)
		// if err != nil {
		// 	code = e.ErrorUploadFile
		// 	return serializer.Response{
		// 		Status: code,
		// 		Error:  path,
		// 		Data:   e.GetMsg(code),
		// 	}
		// }
	}

	// artDao := dao.NewArticleDao(ctx)
	// art, err := artDao.GetArtContentByArtID(artid)
	// art.Cover =

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// GetArtByArtID 获取文章数据
func (service *ArticleService) GetArtByArtID(ctx context.Context, artID string) serializer.Response {
	code := e.SUCCESS
	artid, _ := strconv.Atoi(artID)

	artDao := dao.NewArticleDao(ctx)

	art, err := artDao.GetArtByArtID(uint(artid))
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildArt(art),
	}
}

// 文章创建Create
func (service *ArticleService) Create(ctx context.Context, authid uint, files []*multipart.FileHeader) serializer.Response {
	code := e.SUCCESS

	artDao := dao.NewArticleDao(ctx)
	// 1. 数据判断

	// 2. 数据保存

	var artPhoto []map[string]string
	for _, file := range files {
		opFile, _ := file.Open()
		path, err := UploadToQiNiu(opFile, file.Size)
		if err != nil {
			code = e.ErrorUploadFile
			return serializer.Response{
				Status: code,
				Error:  path,
				Data:   e.GetMsg(code),
			}
		}
		cover := map[string]string{path[len(path)-10:]: path}
		artPhoto = append(artPhoto, cover)
	}

	art := &model.Article{
		CateID:   service.Cate,
		Title:    service.Title,
		AuthorID: authid,
		AreaID:   service.Area,
		Content: model.ArtContent{
			Content: service.Content,
		},
		Cover: model.Cover{
			V: artPhoto,
		},
	}

	err := artDao.Create(art)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildArt(art),
	}
}


// GetContentByArtID  获取文章内容 
func (service *ArticleService) GetContentByArtID(ctx context.Context, artID string) serializer.Response {
	code := e.SUCCESS
	artid, _ := strconv.Atoi(artID)

	artDao := dao.NewArticleDao(ctx)

	art, err := artDao.GetArtContentByArtID(uint(artid))
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildArtContent(art),
	}
}
