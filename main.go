package main





import (
	config "go_ctry/conf"
	"go_ctry/routes"
	
	_ "go_ctry/docs" // 这里需要引入本地已生成文档

)

// import (
//     "encoding/json"
//     "fmt"
// )

// type Node struct {
//     Id       int     `json:"id"`
//     ParentId int     `json:"parent_id"`
//     Name     string  `json:"name"`
//     Children []*Node `json:"children"`
// }

// func getTreeRecursive(list []*Node, parentId int) []*Node {
//     res := make([]*Node, 0)
//     for _, v := range list {
//         if v.ParentId == parentId {
//             v.Children = getTreeRecursive(list, v.Id)
//             res = append(res, v)
//         }
//     }
//     return res
// }



// @title go_ctry API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in f
// @name FanOne
// @BasePath /main.go
func main() {
	config.Init()
	r := routes.NewRouter()
	_ = r.Run(config.HttpPort)
}
