package model

import (
	"gorm.io/gorm"
)

// Comment model
type Comment struct {
	gorm.Model
	Ccoment   string `gorm:"type:varchar(1024)"`
	CparentID uint   `gorm:"ForeignKey:CommentID"`
	ArtID     uint   `gorm:"ForeignKey:ArticleID"`
	UserID    uint   `gorm:"ForeignKey:UserID"`
	LikeCount int64
	Art       Article `gorm:"ForeignKey:ArtID"`
	User      User    `gorm:"ForeignKey:UserID"`
}

func (Comment) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "tb_comment"
}

type CommentRecord struct {
	gorm.Model
	CommentID uint    `gorm:"ForeignKey:CommentID"`
	UserID    uint    `gorm:"ForeignKey:UserID"`
	Commnet   Comment `gorm:"ForeignKey:CommentID"`
	User      User    `gorm:"ForeignKey:UserID"`
	Opera     int     `gorm:"type:tinyint(1); comment:'1 点赞 0 取消点赞'"`
}

func (CommentRecord) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "tb_com_record"
}
