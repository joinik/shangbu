package service

import (
	"context"
	"fmt"
	"go_ctry/dao"
	"go_ctry/model"
	"go_ctry/pkg/e"
	"go_ctry/serializer"
	"sync"
)

// 评论的服务
type CommentService struct {
	model.BasePage
	ParentID uint   `form:"parentID" json:"parentID"`
	Art      uint   `form:"artID" json:"artID" binding:"required"`
	Content  string `form:"content" json:"content"`
}

type ComRecordService struct {
	Opera int  `form:"opera" json:"opera" binding:"oneof=1 0"`
	ComID uint `form:"comID" josn:"comID" binding:"required"`
}

func (service *CommentService) Post(ctx context.Context, userID uint) serializer.Response {
	code := e.SUCCESS
	comDao := dao.NewCommentDao(ctx)

	// 判断数据 是否为空
	if service.Content == "" {
		code = e.InvalidParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 生成 Comment 模型
	com := &model.Comment{
		Ccoment: service.Content,
		UserID:  userID,
		ArtID:   service.Art,
	}
	if service.ParentID != 0 {
		com.CparentID = service.ParentID
	}

	err := comDao.Create(com)
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
	}

}

func (service *CommentService) Get(ctx context.Context) serializer.Response {
	code := e.SUCCESS
	var total int64
	var err error
	var comments []*model.Comment
	commentDao := dao.NewCommentDao(ctx)

	// 判断每页数据几条
	if service.BasePage.PageSize == 0 || service.Total == 0 {
		total, err = commentDao.CountCommentByCondition(map[string]any{"art_id": service.Art})
		service.BasePage.PageSize = 15
		if err != nil {
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}

	// 判断 是查询 子评论 还是文章评论
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		commentDao = dao.NewCommentDaoByDB(commentDao.DB)

		if service.ParentID == 0 {
			// 查询总条数
			comments, err = commentDao.ListCommentByArtID(service.Art, &service.BasePage, fmt.Sprintf("%s DESC", "created_at"))
		} else {
			comments, err = commentDao.ListCommentByID(service.ParentID, &service.BasePage)
		}
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

	return serializer.BuildListResponse(serializer.BuildListComment(comments), int(total))

}

func (service *ComRecordService) PostComRecord(ctx context.Context, userID uint) serializer.Response {
	code := e.SUCCESS
	dao := dao.NewCommentDao(ctx)
	var err error

	// 判断是否有此评论
	com, err := dao.GetCommentByID(service.ComID)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 创建评论 点赞记录
	comRecord := &model.CommentRecord{
		CommentID: service.ComID,
		Opera:     service.Opera,
		UserID:    userID,
	}
	err = dao.CreateComRecord(comRecord)

	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 如果是 点赞操作
	if service.Opera == 1 {
		// 评论点赞数 +1
		com.LikeCount += 1
	} else {
		// 评论点赞数 -1
		com.LikeCount -= 1
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		err = dao.UpdateCommentByID(&com)

	}()

	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	wg.Wait()

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
