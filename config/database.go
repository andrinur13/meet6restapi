package config

import (
	"meet6restapi/structs"

	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	connection := "andri:root@tcp(127.0.0.1:3306)/glngkm?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connection)

	if err != nil {
		panic("failed connect to database")
	}

	db.AutoMigrate(&structs.Person{})

	return db
}
