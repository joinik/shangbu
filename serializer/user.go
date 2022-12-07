package serializer

import (
	"go_ctry/model"
)

type UserProfile struct {
	ID               uint `json:"user_id"`
	Gender           string `json:"user_gender"`
	Age              uint   `json:"age"`
	Email            string `json:"email"`
	DefaultAddressID uint   `json:"default_address_id"`
}

// BuildUser 序列化用户
func BuildUserProfile(userProfile *model.UserProfile) UserProfile {

	return UserProfile{
		ID:               userProfile.UserID,
		Gender:           userProfile.UserGender.Gender,
		Age:              userProfile.Age,
		Email:            userProfile.Email,
		DefaultAddressID: userProfile.DefaultAddressID,
	}
}
