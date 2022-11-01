package model

import (
	"fmt"
	"os"
)

func Migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			// &model.User{},
			&Area{},
		)

	if err != nil {
		fmt.Println("register table fail")
		os.Exit(0)
	}
	fmt.Println("register table success")
}
