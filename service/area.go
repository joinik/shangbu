package service

// import (
// 	"context"
// 	"strconv"

// 	// "encoding/json"

// 	"go_ctry/dao"
// 	"go_ctry/pkg/e"
// 	"go_ctry/serializer"

// 	logging "github.com/sirupsen/logrus"
// )

// type AreaService struct {
// 	AreaID    uint   `json:"area_id" form:"area_id"`
// 	AreaName  string `json:"area_name" form:"area_name"`
// 	CityCode  string `json:"city_code" form:"city_code"`
// 	CityLevel int    `json:"city_level" form:"city_level"`
// 	ParentID  int    `json:"parent_id" form:"parent_id"`
// 	// 	Children []*AreaService
// }

// // AreaGet服务
// func (service *AreaService) Get(ctx context.Context, aid string) serializer.Response {
// 	code := e.SUCCESS

// 	// 转换 adi 为 int 类型
// 	areaID, _ := strconv.Atoi(aid)

// 	areaDao := dao.NewAreaDao(ctx)

// 	if areaID != 0 {
// 		// 根据 aid 查询 下一级
// 		area, err := areaDao.GetAreaByParentID(uint(areaID))
// 		if err != nil {
// 			logging.Info(err)
// 		}

// 		if len(area) == 1 {
// 			// 查询 下一级
// 			areas, err := areaDao.GetAreaByParentID(uint(area[0].ID))
// 			if err != nil {
// 				logging.Info(err)
// 				code = e.ErrorDatabase
// 				return serializer.Response{
// 					Status: code,
// 					Msg:    e.GetMsg(code),
// 				}
// 			}

// 			// 调用序列化数据
// 			res := serializer.BuildArea(areas, uint(area[0].ID))
// 			return serializer.Response{
// 				Status: code,
// 				Data:   res,
// 				Msg:    e.GetMsg(code),
// 			}

// 		}

// 		// 调用序列化数据
// 		res := serializer.BuildArea(area, uint(areaID))
// 		return serializer.Response{
// 			Status: code,
// 			Data:   res,
// 			Msg:    e.GetMsg(code),
// 		}

// 	} else {
// 		// mysql 查询  所有 省份
// 		child, err := areaDao.GetAreaByParentID(uint(0))
// 		if err != nil {
// 			logging.Info(err)
// 			code = e.ErrorDatabase
// 			return serializer.Response{
// 				Status: code,
// 				Msg:    e.GetMsg(code),
// 			}
// 		}

// 		// 调用序列化数据
// 		res := serializer.BuildArea(child, uint(areaID))
// 		return serializer.Response{
// 			Status: code,
// 			Data:   res,
// 			Msg:    e.GetMsg(code),
// 		}

// 	}

// }