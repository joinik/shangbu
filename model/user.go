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

// 定义角色 类型 和常量
type Roles int

const (
	// 买家
	Buyer Roles = iota + 1
	// 经理
	Manager
)

// 给Roles 添加 toString 方法
func (r Roles) String() string {
	switch r {
	case Buyer:
		return "买家"
	case Manager:
		return "经理"
	default:
		return "未知"
	}
}

// User 用户基本模型
type User struct {
	gorm.Model
	UserName   string    `gorm:"unique; type:varchar(20)"`                                   // 用户名
	Mobile     string    `gorm:"unique; not null"`                                           // 手机号
	Avatar     string    `gorm:"default:'https://m.sbwl.com/_nuxt/img/default.c462f6d.png'"` // 头像
	Last_login time.Time // 最后登陆时间
	Introduce  string    // 简介
	Role       int       `gorm:"default:1"` // 角色 默认是买家

	// has many
	UserCollects []UserCollect `gorm:"ForeignKey:UserID"`
}

// BeforeCreate 是一个 GORM hook，在用户对象插入数据库之前被调用。
// 它用于设置用户的最后登录时间为空值。
//
// 参数:
//
//	tx *gorm.DB: 数据库事务实例，可用于执行数据库操作。
//
// 返回值:
//
//	error: 如果操作中出现错误，返回该错误；否则返回 nil。
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// 在用户创建之前，设置最后登录时间为当前时间。
	// 在保存之前设置Last_login为当前时间
	u.Last_login = time.Now()
	return
}

// BeforeUpdate 是一个在用户信息更新之前自动执行的方法。
// 它通过GORM的钩子机制来设置用户的最后登录时间。
// 参数:
//
//	tx *gorm.DB: 数据库事务实例，用于执行数据库操作。
//
// 返回值:
//
//	error: 如果操作过程中出现错误，返回相应的错误信息；否则返回nil。
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	// 更新用户的最后登录时间为当前时间
	// 在用户创建之前，设置最后登录时间为当前时间。
	// 在保存之前设置Last_login为当前时间
	u.Last_login = time.Now()
	return
}

// 用户收藏模型
type UserCollect struct {
	gorm.Model
	UserID    uint // 用户ID
	CollectID uint // 收藏ID
}
