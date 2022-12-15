package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
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

// User 用户基本模型

type User struct {
	gorm.Model
	UserName   string    `gorm:"unique; type:varchar(20)"` // 用户名
	Mobile     string    `gorm:"unique; not null"`         // 手机号
	Avatar     string    // 头像
	Last_login time.Time // 最后登陆时间

	Introduce        string // 简介
	Status           uint   `gorm:"default:1"`         // 状态 0 不可用  1 可用
	Business         uint   `gorm:"default:0"`         // 商家认证 0 不是 1 是
	Dianzan_num      uint   `gorm:"default:0"`         // 点赞数
	Travel_note_num  uint   `gorm:"default:0"`         // 游记数
	DianLiangAreaNum uint   `gorm:"default:0"`         // 点亮地区数
	Last_area_id     uint   `gorm:"ForeignKey:AreaID"` // 用户上次位置

	// has one
	UserProfile UserProfile `gorm:"ForeignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// 设置表名

// func (User) TableName() string {
// 	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
// 	return "user_basic"
// }

// 自定义性别
type MyGender struct {
	Gender []byte
}

func NewGender(v string) (MyGender, error) {
	var g MyGender
	if v != "MAN" && v != "WOMAN" {
		return g, errors.New("只支持 “MAN” 或者 “WOMAN”")
	}
	g.Gender = []byte(v)
	return g, nil
}

// 实现这两个 Value Scan 方法 才能满足 数据库读写
func (g MyGender) Value() (driver.Value, error) {
	return g.Gender, nil
}

func (g *MyGender) Scan(v interface{}) error {
	// g.Gender = v.(string)
	// g.Gender = v.([]uint8)

	switch src := v.(type) {
	case nil:
		return nil

	case string:
		// if an empty gender comes from a table, we return a null gender
		if src == "" {
			return nil
		}
		g.Gender = []byte(src)

	case []byte:
		// if an empty gender comes from a table, we return a null gender
		if len(src) == 0 {
			return nil
		}
		g.Gender = src

	default:
		return fmt.Errorf("Scan: unable to scan type %T into Gender", src)
	}

	return nil
}

// User 用户详情模型
type UserProfile struct {
	UserProfileID    uint `gorm:"primaryKey"`
	UserID           uint
	UserGender       MyGender `gorm:"type: varchar(5); default: 'MAN'"`
	Age              uint
	Email            string `gorm:"type: varchar(20)"`
	DefaultAddressID uint   `gorm:"default: null"`
	// bloog to
	DefaultAddress Address
}

// 用户地址表
type Address struct {
	gorm.Model
	Title      string `gorm:"type: varchar(20)"`  // 地址名称
	ProvinceID uint   `gorm:"ForeignKey: AreaID"` // 省
	CityID     uint   `gorm:"ForeignKey: AreaID"` // 市
	DistrictID uint   `gorm:"ForeignKey: AreaID"` // 区
	Place      string `gorm:"type: varchar(30)"`

	// belong to
	Province Area `gorm:"ForeignKey: ProvinceID"`
	City     Area `gorm:"ForeignKey: CityID"`
	District Area `gorm:"ForeignKey: DistrictID"`
}

// 设置 表名
func (Address) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "tb_address"
}
