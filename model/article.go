package model

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

const (
	DRAFT      = 0 // 草稿
	UNREVIEWED = 1 // 待审核
	APPROVED   = 2 // 审核通过
	FAILED     = 3 // 审核失败
	DELETED    = 4 //  已删除
	BANNED     = 5 //  封禁
)

type Category struct {
	gorm.Model
	CateName string `gorm:"unique; type:varchar(64); not null"`
}

// 设置表名

func (Category) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "tb_category"
}

// Article 文章model
type Article struct {
	gorm.Model
	Title        string `gorm:"type:varchar(128); comment:'文章标题"`
	Cover        Cover  `gorm:"type:json; comment:'文章封面'"`
	Status       int32  `gorm:"default:2;comment:'文章状态'"`
	Reason       string `gorm:"type:varchar(256); comment:'反驳原因"`
	CommentCount int64  `gorm:"comment:'评论总数'"`
	LikeCount    int64  `gorm:"comment:'点赞数'"`
	DisLikeCount int64  `gorm:"comment:'点踩数'"`
	AuthorID     uint   `gorm:"ForeignKey:UserID; comment:'作者ID'"`
	CateID       uint   `gorm:"ForeignKey:CategoryID; comment:'分类ID'"`
	AreaID       uint   `gorm:"ForeignKey:AreaID; comment:'地区ID'"`

	// has one
	Content ArtContent
	Area Area
	Cate Category

	// Belongs To 
	Author User 

}

type Cover struct {
	// Name string `json:"name" form:"url"`
	// Url  string `json:"url" form:"url"`

	V []map[string]string
}

func (c Cover) Value() (driver.Value, error) {
	b, err := json.Marshal(c.V)
	return string(b), err
}

func (c *Cover) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), &c.V)
}

func (Article) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "article_basic"
}

// ArtContent 文章内容modle
type ArtContent struct {
	ArticleID uint   `gorm:"ForeignKey:ArticleID"`
	Content   string `gorm:"type:text"`
}

func (ArtContent) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "tb_article_content"
}

type Spece struct {
	gorm.Model
	SpeIntr  string `gorm:"type:varchar(256); comment:'当地介绍'"`
	SpeCult  string `gorm:"type:varchar(256); comment:'特色文化'"`
	SpeScene string `gorm:"type:varchar(256); comment:'特色美景'"`
	SpeSnack string `gorm:"type:varchar(256); comment:'特色小吃'"`
	SpePhoto Cover  `gorm:"type:json; comment:'特色照片'"`
	SpeTitle string `gorm:"type:varchar(128); comment:'特色标题'"` // 只有用户有
	Story    string `gorm:"type:text; comment:'我的故事'"`
	StoPhoto Cover  `gorm:"type:json; comment:'故事图片'"`
	AuthorID uint   `gorm:"ForeignKey:UserID; comment:'作者ID'"`
	AreaID   uint   `gorm:"ForeignKey:AreaID; comment:'地区ID'"`
}

func (Spece) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "tb_spece"
}


// 文章记录表 
type ArtRecord struct {
	gorm.Model
	ArtID uint `gorm:"ForeignKey:ArticleID"`
	Art Article
	UserID uint `gorm:"ForeignKey:UserID"`
	User User
	Option int64 `gorm:"type:tinyint(1); comment:'记录点赞(1)还是点踩(2)取消点赞(3)取消点踩(4)'"`
}

func (ArtRecord) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "tb_art_record"
}