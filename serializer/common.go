package serializer

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// 带有token的data
type TokenData struct {
	User         interface{} `json:"user"`
	Token        string      `json:"token"`
	RefreshToken string      `json:"refreshToken"`
}

type DataList struct {
	Item  any `json:"item"`
	Total int `json:"total"`
}

// BulidListResponse 带有总数的列表构建器
func BuildListResponse(items any, total int) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: "ok",
	}
}
