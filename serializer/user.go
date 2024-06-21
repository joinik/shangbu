package serializer

import (
	"go_ctry/model"
)

type User struct {
	ID         uint   `json:"id"`
	UserName   string `json:"user_name"`
	Mobile     string `json:"mobile"`
	Avatar     string `json:"avatar"`
	Last_login int64  `json:"last_login"`
	Introduce  string `json:"introduce"`
	Status     uint   `json:"status"`
}

type UserProfile struct {
	ID               uint   `json:"user_id"`
	Gender           string `json:"user_gender"`
	Age              uint   `json:"age"`
	Email            string `json:"email"`
	DefaultAddressID uint   `json:"default_address_id"`
}

// BuildUser 序列化用户
func BuildUser(user *model.User) User {

	return User{
		ID: user.ID,
		UserName: user.UserName,
		Mobile: user.Mobile,
		Avatar: user.Avatar,
		Last_login: user.Last_login.Unix(),
		Introduce: user.Introduce,
	}
}

