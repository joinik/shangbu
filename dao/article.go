package dao

import (
	"context"
	"go_ctry/model"

	"gorm.io/gorm"
)

type ArticleDao struct {
	*gorm.DB
}

func NewArticleDao(ctx context.Context) *ArticleDao {
	return &ArticleDao{NewDBClient(ctx)}
}

func NewArticleDaoByDB(db *gorm.DB) *ArticleDao {
	return &ArticleDao{db}
}



// Create 创建文章
func (dao *ArticleDao) Create(art *model.Article) (err error) {
	err = dao.DB.Model(&model.Article{}).Create(&art).Error
	return
}

// UpdateArtByAuthID 根据AuthID和ArtID更新文章
func (dao *ArticleDao) UpdateArt(art *model.Article, column string, value interface{}) (err error) {
	if value == nil {
		err = dao.DB.Model(&model.Article{}).Where("id=?", art.ID).Updates(&art).Error
	} else {
		err = dao.DB.Model(&model.Article{}).Where("id=?", art.ID).Update(column, value).Error
	}
	return
}

// UpdateArtContent 更新文章内容 
func (dao *ArticleDao) UpdateArtContent(art *model.Article) (err error) {
	err = dao.DB.Model(&model.ArtContent{}).Where("article_id=?", art.ID).Updates(art.Content).Error
	return
}

// CountArtByCondition 获取条件获取文章数量  
func (dao *ArticleDao) CountArtByCondition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Comment{}).Where(condition).Count(&total).Error
	return
}


// GetArtByAuthID 根据AuthID查询文章
func (dao *ArticleDao) ListArtByAuthID(authID uint, page *model.BasePage, order string) (art []*model.Article, err error) {
	err = dao.DB.Model(&model.Article{}).Where("author_id=?", authID).
		Preload("Cate").Preload("Area").Preload("Author").
		Offset((page.PageNum -1) * page.PageSize).Limit(page.PageSize).
		Order(order).Find(&art).Error
	return
}

// GetArtByAreaID 根据AreaID查询文章
func (dao *ArticleDao) ListArtByAreaID(areaID uint, page *model.BasePage, order string) (art []*model.Article, err error) {
	err = dao.DB.Model(&model.Article{}).Where("area_id=?", areaID).
		Preload("Cate").Preload("Area").Preload("Author").
		Offset((page.PageNum -1) * page.PageSize).Limit(page.PageSize).
		Order(order).Find(&art).Error
	return
}

// GetArtByCateID 根据CateID查询文章
func (dao *ArticleDao) ListArtByCateID(cateID uint, page *model.BasePage, order string) (art []*model.Article, err error) {
	err = dao.DB.Model(&model.Article{}).Where("cate_id=?", cateID).
		Preload("Cate").Preload("Area").Preload("Author").
		Offset((page.PageNum -1) * page.PageSize).Limit(page.PageSize).
		Order(order).Find(&art).Error
	return
}

// GetArtContentByArtID 根据文章id查询文章信息
func (dao *ArticleDao) GetArtByArtID(artid uint) (art *model.Article, err error) {
	err = dao.DB.Model(&model.Article{}).Where("id=?", artid).
		Preload("Cate").Preload("Area").Preload("Author").First(&art).Error
	return
}

// GetArtContentByArtID 根据文章id查询文章内容
func (dao *ArticleDao) GetArtContentByArtID(artID uint) (art *model.ArtContent, err error) {
	err = dao.DB.Model(&model.ArtContent{}).Where("article_id=?", artID).First(&art).Error
	return
}

// CreateArtRecord 创建文章点赞/点踩记录 
func (dao *ArticleDao) CreateArtRecord(artRecord *model.ArtRecord) (err error) {

	// 创建文章点赞/ 点踩记录
	err = dao.DB.Model(&model.ArtRecord{}).Create(&artRecord).Error
	return
}

// GetArtRecordByCondition 查询文章点赞 / 点踩 根据条件 
func (dao *ArticleDao) GetArtRecordByCondition(artID uint, userID uint, option int64) (artRecord *model.ArtRecord, err error) {

	// 创建文章点赞/ 点踩记录
	err = dao.DB.Model(&model.ArtRecord{}).Where("art_id=?,user_id=?,option=?", artID,userID,option).Limit(1).First(&artRecord).Error
	return
}