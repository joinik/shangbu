package dao

import (
	"context"
	"go_ctry/model"

	"gorm.io/gorm"
)

type AreaDao struct {
	*gorm.DB
}

func NewAreaDao(ctx context.Context) *AreaDao {
	return &AreaDao{NewDBClient(ctx)}
}

// GetAreaByParentID 根据parent_id 查询area数据
func (dao *AreaDao) GetAreaByParentID(parent uint) (areaLi []*model.Area, err error) {
    err = dao.DB.Model(&model.Area{}).Where("parent_id=?", parent).Find(&areaLi).Error
    return
}
