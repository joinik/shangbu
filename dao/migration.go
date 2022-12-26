package dao

import (
	"fmt"
	"go_ctry/model"
	"os"
)

func Migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			// &model.User{},
			&model.Area{},
			&model.User{},
			&model.UserProfile{},
			&model.Address{},
			&model.Article{},
			&model.ArtContent{},
			&model.Spece{},
			&model.Category{},
			&model.ArtRecord{},
		)

	if err != nil {
		fmt.Println("register table fail")
		os.Exit(0)
	}
	fmt.Println("register table success")
}
