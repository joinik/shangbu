package dao

import (
	"context"
	"go_ctry/model"

	"gorm.io/gorm"
)

type CommentDao struct {
	*gorm.DB
}

func NewCommentDao(ctx context.Context) *CommentDao {
	return &CommentDao{NewDBClient(ctx)}
}

func NewCommentDaoByDB(db *gorm.DB) *CommentDao {
	return &CommentDao{db}
}

// Create 创建评论
func (commentDao *CommentDao) Create(comm *model.Comment) (err error) {
	err = commentDao.DB.Model(&model.Comment{}).Create(&comm).Error
	return
}

// CountCommentByCondition 根据条件获取评论数量
func (commentDao *CommentDao) CountCommentByCondition(condition map[string]interface{}) (total int64, err error) {
	err = commentDao.DB.Model(&model.Comment{}).Where(condition).Count(&total).Error
	return
}

// GetCommentsByArtID 根据 ArtID 查询 评论 返回切片
func (commentDao *CommentDao) ListCommentByArtID(artid uint, page *model.BasePage, order string) (comments []*model.Comment, err error) {
	err = commentDao.DB.Model(&model.Comment{}).Where("art_id=? and cparent_id=0", artid).Preload("User").
		Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).
		Order(order).
		Find(&comments).Error
	return
}

// GetCommentsByID 根据 commentID 查询 子评论 返回切片
func (commentDao *CommentDao) ListCommentByID(comid uint, page *model.BasePage) (comments []*model.Comment, err error) {
	err = commentDao.DB.Model(&model.Comment{}).Where("cparent_id=?", comid).Preload("User").
		Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&comments).Error
	return
}

// GetCommentByID 根据评论id 查询评论
func (commentDao *CommentDao) GetCommentByID(comid uint) (com model.Comment, err error) {
	err = commentDao.DB.Model(&model.Comment{}).Preload("User").Where("id=?", comid).First(&com).Error
	return
}

// UpdateCommentByID 根据评论id 更新评论
func (commentDao *CommentDao) UpdateCommentByID(com *model.Comment) (err error) {
	err = commentDao.DB.Model(&model.Comment{}).Where("id=?", com.ID).Updates(&com).Error
	return
}

// CreateComRecord 创建评论点赞
func (commentDao *CommentDao) CreateComRecord(comRecord *model.CommentRecord) (err error) {
	err = commentDao.DB.Model(&model.CommentRecord{}).Create(&comRecord).Error
	return
}

// GetComRecordByUserID 根据用户id 查询评论点赞
func (commentDao *CommentDao) GetComRecordsByUserID(userID uint) (comRecords []model.CommentRecord, err error) {
	err = commentDao.DB.Model(&model.CommentRecord{}).Where("user_id=?", userID).
		Order("created_at DESC").Find(&comRecords).Error

	return
}
