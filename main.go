package main

import (
	config "go_ctry/conf"
	"go_ctry/routes"
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

func main() {
	config.Init()
	r := routes.NewRouter()
	_ = r.Run(config.HttpPort)

	// list := []*Node{
	//     {4, 3, "ABA", nil},
	//     {3, 1, "AB", nil},
	//     {1, 0, "A", nil},
	//     {2, 1, "AA", nil},
	// }
	// res := getTreeRecursive(list, 0)
	// bytes, _ := json.MarshalIndent(res, "", "    ")
	// fmt.Printf("%s\n", bytes)

}
