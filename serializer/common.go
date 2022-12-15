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
