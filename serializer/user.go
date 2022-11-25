package serializer

import (
	"fmt"
	"go_ctry/model"
)

type UserProfile struct {
	ID               uint   `json:"user_id"`
	Gender           string `json:"user_gender"`
	Age              uint    
	Email            string
	DefaultAddressID uint
}

// BuildUser 序列化用户
func BuildUserProfile(userProfile *model.UserProfile) UserProfile {
	fmt.Println("-----------<>____________")
	fmt.Printf("%v", userProfile.UserGender)

	return UserProfile{
		ID:               userProfile.UserID,
		Gender:           "gender",
		Age:              userProfile.Age,
		Email:            userProfile.Email,
		DefaultAddressID: userProfile.DefaultAddressID,
	}
}
