package service

import (
	"context"
	// "encoding/json"
	"fmt"
	"go_ctry/model"
	"go_ctry/pkg/e"
	"go_ctry/serializer"
	"strconv"

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

func (service *AreaService) Get(ctx context.Context, aid string) serializer.Response {
	code := e.SUCCESS

	// 转换 adi 为 int 类型
	area_id, _ := strconv.Atoi(aid)
	// 声明一个空的 Area 模型类
	area := &model.Area{}

	// 定义切片 area_li
	var area_li []*model.Area

	// 定义切片 area_li_2
	var area_li_2 []*model.Area

	// 创建会话
	db := model.NewDBClient(ctx)

	// mysql 查询  省份
	err := db.Model(area).Where("parent_id=? or id=?", area_id, area_id).Find(&area_li).Error
	if err != nil {
		logging.Info(err)
	}

	for _, item := range area_li {
		fmt.Println(item.Name)
	}

	// mysql 查询  市区 
	err1 := db.Model(area).Where("parent_id=?", area_li[1].ID).Find(&area_li_2).Error

	if err1 != nil {
		logging.Info(err1)
	}

	// 合共 两个 slice area_li_2... 可变参数
	area_li = append(area_li, area_li_2...)
	
	// 调用序列化数据
	res := serializer.BuildArea(area_li)

	return serializer.Response{
		Status: code,
		Data:   res,
		Msg:    e.GetMsg(code),
	}

}
