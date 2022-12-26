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
	Title      string `form:"title" json:"title"`             // 文章标题
	Cate       uint   `form:"cate" json:"cate"`               // 分类id
	Area       uint   `form:"area" json:"area"`               // 地区id
	Content    string `form:"content" json:"content"`         // 文章内容
	ArtID      uint   `form:"art_id" json:"art_id"`           // 文章id
	OptionFlag uint   `form:"option_flag" json:"option_flag"` // 1 点赞 点踩，0 取消点赞，点踩
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

	err = artDao.UpdateArt(art, "", nil, service.ArtID)
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
func (service *ArticleService) GetArtsByCateID(ctx context.Context, cateID string) serializer.Response {
	code := e.SUCCESS

	artDao := dao.NewArticleDao(ctx)
	cate_id, _ := strconv.Atoi(cateID)
	arts, err := artDao.GetArtByCateID(uint(cate_id))
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
		Data:   serializer.BuildArts(arts),
	}

}

// GetArtsByCateID 根据分类id 查询文章信息
func (service *ArticleService) GetArtsByAreaID(ctx context.Context, areaID string) serializer.Response {
	code := e.SUCCESS

	artDao := dao.NewArticleDao(ctx)
	area_id, _ := strconv.Atoi(areaID)
	arts, err := artDao.GetArtByAreaID(uint(area_id))
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
		Data:   serializer.BuildArts(arts),
	}

}

// 根据文章ID 文章点赞
func (service *ArticleService) ArtLiked(ctx context.Context, userid uint) serializer.Response {
	code := e.SUCCESS

	artDao := dao.NewArticleDao(ctx)

	if service.OptionFlag == 1 {
		// OptionFlag == 1 表示点赞
		artRecord := &model.ArtRecord{
			ArtID:  service.ArtID,
			UserID: userid,
			Option: 1, // 点赞记录
		}

		err := artDao.CreateArtRecord(service.ArtID, userid, artRecord)
		if err != nil {
			code := e.ErrorDatabase
			return serializer.Response{
				Msg:    e.GetMsg(code),
				Status: code,
			}
		}

		art, err := artDao.GetArtByArtID(service.ArtID)
		if err != nil {
			code := e.ErrorDatabase
			return serializer.Response{
				Msg:    e.GetMsg(code),
				Status: code,
			}
		}
		// 文章点赞数 +=1 
		art.LikeCount +=1 

		err = artDao.UpdateArt(art,"like_count",art.LikeCount ,service.ArtID)
		if err != nil {
			code := e.ErrorDatabase
			return serializer.Response{
				Msg:    e.GetMsg(code),
				Status: code,
			}
		}
		


	} else {
		// 取消点赞操作 

		artRecord := &model.ArtRecord{
			ArtID:  service.ArtID,
			UserID: userid,
			Option: 3, // 3 取消点赞
		}

		err := artDao.CreateArtRecord(service.ArtID, userid, artRecord)
		if err != nil {
			code := e.ErrorDatabase
			return serializer.Response{
				Msg:    e.GetMsg(code),
				Status: code,
			}
		}

		art, err := artDao.GetArtByArtID(service.ArtID)
		if err != nil {
			code := e.ErrorDatabase
			return serializer.Response{
				Msg:    e.GetMsg(code),
				Status: code,
			}
		}
		// 文章点赞数 -=1 
		art.LikeCount -=1 

		err = artDao.UpdateArt(nil, "like_count", art.LikeCount, service.ArtID)
		if err != nil {
			code := e.ErrorDatabase
			return serializer.Response{
				Msg:    e.GetMsg(code),
				Status: code,
			}
		}

	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}

}

// 根据文章ID 文章点踩
func (service *ArticleService) ArtDisliked(ctx context.Context, userid uint) serializer.Response {
	code := e.SUCCESS

	artDao := dao.NewArticleDao(ctx)

	if service.OptionFlag == 1 {
		// 点踩操作
		artRecord := &model.ArtRecord{
			ArtID:  service.ArtID,
			UserID: userid,
			Option: 2, // 点踩记录
		}

		err := artDao.CreateArtRecord(service.ArtID, userid, artRecord)
		if err != nil {
			code := e.ErrorDatabase
			return serializer.Response{
				Msg:    e.GetMsg(code),
				Status: code,
			}
		}

		art, err := artDao.GetArtByArtID(service.ArtID)

		if err != nil {
			code := e.ErrorDatabase
			return serializer.Response{
				Msg:    e.GetMsg(code),
				Status: code,
			}
		}
		// 文章点踩数+1
		art.DisLikeCount += 1
		err = artDao.UpdateArt(art, "dis_like_count",art.DisLikeCount ,service.ArtID)

		if err != nil {
			code := e.ErrorDatabase
			return serializer.Response{
				Msg:    e.GetMsg(code),
				Status: code,
			}
		}

	} else {
		// 取消点踩 操作 
		artRecord := &model.ArtRecord{
			ArtID:  service.ArtID,
			UserID: userid,
			Option: 4, // 取消点踩记录
		}

		err := artDao.CreateArtRecord(service.ArtID, userid, artRecord)
		if err != nil {
			code := e.ErrorDatabase
			return serializer.Response{
				Msg:    e.GetMsg(code),
				Status: code,
			}
		}

		art, err := artDao.GetArtByArtID(service.ArtID)

		if err != nil {
			code := e.ErrorDatabase
			return serializer.Response{
				Msg:    e.GetMsg(code),
				Status: code,
			}
		}
		// 文章点踩数 -1
		art.DisLikeCount -= 1
		err = artDao.UpdateArt(nil, "dis_like_count", art.DisLikeCount, service.ArtID)

		if err != nil {
			code := e.ErrorDatabase
			return serializer.Response{
				Msg:    e.GetMsg(code),
				Status: code,
			}
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}

}
