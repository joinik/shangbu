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
	ArtID   uint   `form:"art_id" json:"art_id"`
}

// UploadArtPhoto 根据文章id 更新文章
func (service *ArticleService) UploadArt(ctx context.Context, authid uint, files []*multipart.FileHeader) serializer.Response {
	code := e.SUCCESS
	// var err error
	artDao := dao.NewArticleDao(ctx)
	var cover []map[string]string
	for _, file := range files {
		content, _ := file.Open()
		// 七牛云上传
		path, err := UploadToQiNiu(content, file.Size)
		if err != nil {
			code = e.ErrorUploadFile
			return serializer.Response{
				Status: code,
				Error:  path,
				Data:   e.GetMsg(code),
			}
		}
		cover = append(cover, map[string]string{"url": path})
	}

	// artDao := dao.NewArticleDao(ctx)
	art, err := artDao.GetArtByArtID(service.ArtID)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	art.Title = service.Title
	art.Content.Content = service.Content
	art.Cover = model.Cover{V: cover}

	err = artDao.UpdateArt(art, authid, service.ArtID)
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
		Data: serializer.BuildArt(art),
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

// Create 文章创建
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
		cover := map[string]string{"url": path}
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


// getArtsByCateID 根据分类id 查询文章信息
func (service *ArticleService) GetArtsByCateID(ctx context.Context, cateID string) serializer.Response  {
	code := e.SUCCESS

	artDao := dao.NewArticleDao(ctx)
	cate_id, _ := strconv.Atoi(cateID)
	arts, err := artDao.GetArtByCateID(uint(cate_id))
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
		Data: serializer.BuildArts(arts),
	}


}


// getArtsByCateID 根据分类id 查询文章信息
func (service *ArticleService) GetArtsByAreaID(ctx context.Context, areaID string) serializer.Response  {
	code := e.SUCCESS

	artDao := dao.NewArticleDao(ctx)
	area_id, _ := strconv.Atoi(areaID)
	arts, err := artDao.GetArtByAreaID(uint(area_id))
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
		Data: serializer.BuildArts(arts),
	}


}



