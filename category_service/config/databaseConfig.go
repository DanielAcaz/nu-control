package config

import (
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql" //MySQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func GetConnection() *gorm.DB {
	db, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/category_service_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err.Error())
	}

	return db

}
