package my

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func Migrate() {
	db, er := gorm.Open("sqlite3", "data.sqlite3")
	if er != nil {
		fmt.Println(er)
		return
	}
	defer db.Close()

	db.AutoMigrate(&User{}, &Group{}, &Post{}, &Comment{})
}

func Reset() {
	db, er := gorm.Open("sqlite3", "data.sqlite3")
	if er != nil {
		fmt.Println(er)
		return
	}
	defer db.Close()

	db.DropTable(&User{}, &Group{}, &Post{}, &Comment{})
}
