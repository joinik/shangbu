package model

import "gorm.io/gorm"

// 自定义类型 房屋类型
type AddrKing int

const (
	//共享办公
	SharedOffice AddrKing = iota + 1
	//租写字楼
	RentOfficeSpace
	//购买写字楼
	BuyOfficeSpace
	//找楼盘
	FindHouse
	//买商业新房
	BuyNewHouse
	//租商业新房
	RentNewHouse
	//租商铺
	RentShop
	//购买商铺
	BuyShop
	//租厂房
	RentDeliveryRoom
	//买厂房
	BuyDeliveryRoom
)

// 地区类
type Area struct {
	gorm.Model
	Name      string `gorm:"uniqured; type:varchar(20)"` //  地区名称
	City_code string `gorm:"not null"`                   // 地区英文
}

// 房子类详情
// type House struct {

// }
