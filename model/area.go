package model

import "gorm.io/gorm"

type Area struct {
	gorm.Model
	Name       string  `gorm:"type:varchar(20)"`                //  地区名称
	City_code  string  `gorm:"type:varchar(12); not null"`      // 地区编码
	City_level int     `gorm:"not null"`                        // 地区级别
	ParentID   uint     `gorm:"ForeignKey:AreaID; index"` // 父级id


	// Parent     []*Area `gorm:"ForeignKey:ParentID"`
}
