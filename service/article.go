package service

import (
	"context"
	"go_ctry/dao"
	"go_ctry/model"
	"go_ctry/pkg/e"
	"go_ctry/serializer"
	"mime/multipart"
	"strconv"
	"sync"
)

type ArticleService struct {
	Title   string `form:"title" json:"title"`     // 文章标题
	Cate    uint   `form:"cate" json:"cate"`       // 分类id
	Area    uint   `form:"area" json:"area"`       // 地区id
	Content string `form:"content" json:"content"` // 文章内容
	ArtID   uint   `form:"art_id" json:"art_id"`   // 文章id
	model.BasePage
}

type ArtRecordService struct {
	ArtID      uint `form:"art_id" json:"art_id" binding:"ne=0"`                    // 文章id
	OptionFlag uint `form:"option" json:"option" binding:"oneof=1 2 3 4"` // 1 点赞 2 点踩，3 取消点赞，4 取消点踩

}

// UploadArtPhoto 根据文章id 更新文章
func (service *ArticleService) UploadArt(ctx context.Context, authid uint, files []*multipart.FileHeader) serializer.Response {
	code := e.SUCCESS
	artDao := dao.NewArticleDao(ctx)

	// 上传图片到 七牛云
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

	art, err := artDao.GetArtByArtID(service.ArtID)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 更新文章信息数据
	art.Title = service.Title
	art.Content.Content = service.Content
	art.Cover = model.Cover{V: cover}
	err = artDao.UpdateArt(art, "", nil)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 更新文章内容数据
	err = artDao.UpdateArtContent(art)
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
	var total int64
	var err error
	var arts []*model.Article

	artDao := dao.NewArticleDao(ctx)
	cate_id, _ := strconv.Atoi(cateID)

	if service.Total == 0 || service.PageSize == 0 && cate_id != 0 {
		// 获取文章 总数量
		total, err = artDao.CountArtByCondition(map[string]interface{}{"cate_id": service.Cate})
		service.PageSize = 15
		if err != nil {
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	} else {
		code = e.InvalidParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		artDao = dao.NewArticleDaoByDB(artDao.DB)
		arts, err = artDao.ListArtByCateID(uint(cate_id), &service.BasePage, "created_at DESC")
		wg.Done()
	}()
	wg.Wait()

	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.BuildListResponse(serializer.BuildArts(arts), int(total))
}

// GetArtsByCateID 根据分类id 查询文章信息
func (service *ArticleService) GetArtsByAreaID(ctx context.Context, areaID string) serializer.Response {
	code := e.SUCCESS
	var total int64
	var err error
	var arts []*model.Article
	artDao := dao.NewArticleDao(ctx)
	area_id, _ := strconv.Atoi(areaID)

	if service.Total == 0 || service.PageSize == 0 && area_id != 0 {
		// 获取文章 总数量
		total, err = artDao.CountArtByCondition(map[string]interface{}{"cate_id": service.Cate})
		service.PageSize = 15
		if err != nil {
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	} else {
		code = e.InvalidParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		artDao = dao.NewArticleDaoByDB(artDao.DB)
		arts, err = artDao.ListArtByAreaID(uint(area_id), &service.BasePage, "created_at DESC")
		wg.Done()
	}()
	wg.Wait()

	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.BuildListResponse(serializer.BuildArts(arts), int(total))

}

// 根据文章ID 文章点赞
func (service *ArtRecordService) ArtLiked(ctx context.Context, userID uint) serializer.Response {
	code := e.SUCCESS
	artDao := dao.NewArticleDao(ctx)
	var err error
	var artRecord *model.ArtRecord
	var art *model.Article

	if service.ArtID == 0 {
		code = e.InvalidParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		artDao = dao.NewArticleDaoByDB(artDao.DB)
		// 查询是否有此文章 根据id
		art, err = artDao.GetArtByArtID(service.ArtID)

	}()
	if err != nil {
		code := e.ErrorDatabase
		return serializer.Response{
			Msg:    e.GetMsg(code),
			Status: code,
		}
	}
	if art == nil {
		return serializer.Response{
			Status: 400,
			Msg: "错误请求",
		}
	}

	go func() {
		// 查询是否已经点赞过
		artDao = dao.NewArticleDaoByDB(artDao.DB)
		artRecord, _ = artDao.GetArtRecordByCondition(service.ArtID, userID, int64(service.OptionFlag))
		wg.Done()
	}()
	if artRecord != nil {
		code = e.ErrorCallApi
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 进行业务判断 

	if service.OptionFlag == 1 {
		// OptionFlag == 1 表示点赞
		wg.Add(2)
		go func() {
			artRecord := &model.ArtRecord{
				ArtID:  service.ArtID,
				UserID: userID,
				Option: 1, // 点赞记录
			}
			artDao = dao.NewArticleDaoByDB(artDao.DB)
			err = artDao.CreateArtRecord(artRecord)
			wg.Done()
		}()

		if err != nil {
			code := e.ErrorDatabase
			return serializer.Response{
				Msg:    e.GetMsg(code),
				Status: code,
			}
		}
		go func() {
			// 文章点赞数 +=1
			art.LikeCount += 1
			artDao = dao.NewArticleDaoByDB(artDao.DB)
			err = artDao.UpdateArt(art, "like_count", art.LikeCount)
			wg.Done()
		}()

		if err != nil {
			code := e.ErrorDatabase
			return serializer.Response{
				Msg:    e.GetMsg(code),
				Status: code,
			}
		}

	} else if service.OptionFlag == 3 {
		// 取消点赞操作
		wg.Add(2)
		go func() {
			artRecord := &model.ArtRecord{
				ArtID:  service.ArtID,
				UserID: userID,
				Option: 3, // 3 取消点赞
			}
			artDao = dao.NewArticleDaoByDB(artDao.DB)
			err = artDao.CreateArtRecord(artRecord)
			wg.Done()
		}()

		if err != nil {
			code := e.ErrorDatabase
			return serializer.Response{
				Msg:    e.GetMsg(code),
				Status: code,
			}
		}

		go func() {
			// 文章点赞数 -=1
			art.LikeCount -= 1
			artDao = dao.NewArticleDaoByDB(artDao.DB)
			err = artDao.UpdateArt(art, "like_count", art.LikeCount)
		}()
		if err != nil {
			code := e.ErrorDatabase
			return serializer.Response{
				Msg:    e.GetMsg(code),
				Status: code,
			}
		}

	}else if service.OptionFlag == 2 {
		// 点踩操作
		wg.Add(2)
		go func() {
			artRecord := &model.ArtRecord{
				ArtID:  service.ArtID,
				UserID: userID,
				Option: 2, // 2 点踩操作
			}
			artDao = dao.NewArticleDaoByDB(artDao.DB)
			err = artDao.CreateArtRecord(artRecord)
			wg.Done()
		}()

		if err != nil {
			code := e.ErrorDatabase
			return serializer.Response{
				Msg:    e.GetMsg(code),
				Status: code,
			}
		}

		go func() {
			// 文章点赞数 -=1
			art.DisLikeCount += 1
			artDao = dao.NewArticleDaoByDB(artDao.DB)
			err = artDao.UpdateArt(art, "like_count", art.LikeCount)
		}()
		if err != nil {
			code := e.ErrorDatabase
			return serializer.Response{
				Msg:    e.GetMsg(code),
				Status: code,
			}
		}

	}else if service.OptionFlag == 4 {
		// 取消点踩操作
		wg.Add(2)
		go func() {
			artRecord := &model.ArtRecord{
				ArtID:  service.ArtID,
				UserID: userID,
				Option: 4, // 4 取消点踩
			}
			artDao = dao.NewArticleDaoByDB(artDao.DB)
			err = artDao.CreateArtRecord(artRecord)
			wg.Done()
		}()

		if err != nil {
			code := e.ErrorDatabase
			return serializer.Response{
				Msg:    e.GetMsg(code),
				Status: code,
			}
		}

		go func() {
			// 文章点踩数 -=1
			art.DisLikeCount -= 1
			artDao = dao.NewArticleDaoByDB(artDao.DB)
			err = artDao.UpdateArt(art, "like_count", art.LikeCount)
		}()
		if err != nil {
			code := e.ErrorDatabase
			return serializer.Response{
				Msg:    e.GetMsg(code),
				Status: code,
			}
		}

	}
	wg.Wait()

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}

}

func (service *ArticleService) GetCates(ctx context.Context) serializer.Response {
	code := e.SUCCESS
	artDao := dao.NewArticleDao(ctx)
	cates, err := artDao.ListCategory()
	if err !=nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Data: serializer.BuildListCate(cates),
		Msg: e.GetMsg(code),
	}


}


