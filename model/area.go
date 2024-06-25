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
	// 自引用 has many
	ParentID *uint
	Cicys    []Area `gorm:"ForeignKey:ParentID"`
}

// 房子类详情
type House struct {
	gorm.Model
	AffairsGardnName string `gorm:"type:varchar(20)"`  // 广场/花园  （天德广场）
	HouseType        int    `gorm:"type:int"`          // 房屋类型
	HouseArea        int    `gorm:"type:int"`          // 房屋面积 2000
	HousePrice       int    `gorm:"type:int"`          // 房屋价格 67600
	Decoration       string `gorm:"type:varchar(20)"`  // 装修情况	精装修
	Cover            string `gorm:"type:varchar(100)"` // 房屋封面
	City             Area   `gorm:"ForeignKey:CityID"` // 房屋所在城市
	Title            string `gorm:"type:varchar(88)"`  // 房屋标题	天德广场 精装修 高层 2000平米
	RegionName       string `gorm:"type:varchar(20)"`  // 房屋区域名称
	BusinessAreaName string `gorm:"type:varchar(20)"`  // 房屋商圈名称 珠江新城中
	FloorName        string `gorm:"type:varchar(20)"`  // 房屋楼层名称 高层
	PersonCountStart int    `gorm:"type:int"`          // 240-360人
	PersonCountEnd   int    `gorm:"type:int"`
}
