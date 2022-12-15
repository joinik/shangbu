package service

import (
	"context"
	"strconv"

	// "encoding/json"

	"go_ctry/model"
	"go_ctry/pkg/e"
	"go_ctry/serializer"

	logging "github.com/sirupsen/logrus"
)

type AreaService struct {
	AreaID    uint   `json:"area_id" form:"area_id"`
	AreaName  string `json:"area_name" form:"area_name"`
	CityCode  string `json:"city_code" form:"city_code"`
	CityLevel int    `json:"city_level" form:"city_level"`
	ParentID  int    `json:"parent_id" form:"parent_id"`
	// 	Children []*AreaService
}

// AreaGet服务
func (service *AreaService) Get(ctx context.Context, aid string) serializer.Response {
	code := e.SUCCESS

	// 转换 adi 为 int 类型
	areaID, _ := strconv.Atoi(aid)
	// 声明一个空的 Area 模型类
	area := &model.Area{}

	// 定义切片 area_li
	var area_li []*model.Area

	// 创建会话
	db := model.NewDBClient(ctx)

	if areaID != 0 {
		// 根据 aid 查询 下一级
		err := db.Model(area).Where("parent_id=?", uint(areaID)).Find(&area).Error
		if err != nil {
			logging.Info(err)
		}

		// 根据市辖区 查询 下一级(市)
		err = db.Model(area).Where("parent_id=?", uint(area.ID)).Find(&area_li).Error
		if err != nil {
			logging.Info(err)
		}

		// 查询 县
		for _, item := range area_li {
			var child []*model.Area
			err := db.Model(area).Where("parent_id=?", uint(item.ID)).Find(&child).Error
			if err != nil {
				logging.Info(err)
			}
			area_li = append(area_li, child...)

			// 查询 镇
			// for _, item := range area_li {
			// 	var child []*model.Area
			// 	err := db.Model(area).Where("parent_id=?", uint(item.ID)).Find(&child).Error
			// 	if err != nil {
			// 		logging.Info(err)
			// 	}
			// 	area_li = append(area_li, child...)

			// 	// 查询 村
			// 	for _, item := range area_li {
			// 		var child []*model.Area
			// 		err := db.Model(area).Where("parent_id=?", uint(item.ID)).Find(&child).Error
			// 		if err != nil {
			// 			logging.Info(err)
			// 		}
			// 		area_li = append(area_li, child...)

			// 	}

			// }

		}

		// 调用序列化数据
		res := serializer.BuildArea(area_li, uint(area.ID))
		return serializer.Response{
			Status: code,
			Data:   res,
			Msg:    e.GetMsg(code),
		}

	} else {
		// mysql 查询  所有 省份
		err := db.Model(area).Where("parent_id=?", 0).Find(&area_li).Error
		if err != nil {
			logging.Info(err)
		}

	}

	// 调用序列化数据
	res := serializer.BuildArea(area_li, uint(areaID))
	return serializer.Response{
		Status: code,
		Data:   res,
		Msg:    e.GetMsg(code),
	}

}
