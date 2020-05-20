package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// CreateDBConnection returns db connection
func CreateDBConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(mysql-container:3306)/go_app")
	if err != nil {
		fmt.Println("error in db connection")
		return db
	}

	db.LogMode(true)

	return db
}
