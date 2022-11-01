package model

import (
	"time"

	"gorm.io/gorm"
)

// 嵌入 结构体
// gorm.Model 的定义
// type Model struct {
//   ID        uint           `gorm:"primaryKey"`
//   CreatedAt time.Time       默认为当前时间
//   UpdatedAt time.Time       默认为当前时间
//   DeletedAt gorm.DeletedAt `gorm:"index"`
// }

// User 用户模型

type User struct {
	gorm.Model
	UserName         string    `gorm:"unique; type:varchar(20)"` // 用户名
	Mobile           string    `gorm:"unique; not null"`         // 手机号
	Avatar           string    // 头像
	Last_login       time.Time // 最后登陆时间
	Intrude          string    // 简介
	Status           int       `gorm:"default:1"`         // 状态 0 不可用  1 可用
	Business         int       `gorm:"default:0"`         // 商家认证 0 不是 1 是
	Dianzan_num      int64     `gorm:"default:0"`         // 点赞数
	Travel_note_num  int64     `gorm:"default:0"`         // 游记数
	DianLiangAreaNum int64     `gorm:"default:0"`         // 点亮地区数
	Last_area_id     int64     `gorm:"ForeignKey:AreaID"` // 用户上次位置

}
