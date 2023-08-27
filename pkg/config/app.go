package config

import (
	"gorm.io/gorm"
	_ "gorm.io/driver/mysql"
)

// This variable will help us interact with the db regarding any operations
var db *gorm.DB

// Use gorm to connect to mysql server
func Connect(){
	d, err := gorm.Open("mysql", "rootxrishabh:rootxrishabh/booksTable?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db = d
}

// return the object
func GetDB() *gorm.DB{
	return db
}

