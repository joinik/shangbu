package model

type BasePage struct {
	PageNum  int `form:"pageNum" json:"pageNum"`
	PageSize int `form:"pageSize" json:"pageSize"`
	Total    int `form:"total" json:"total"`
}
