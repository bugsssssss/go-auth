package database

import (
	"auth/models"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	login := "root"
	password := "rootroot"
	db := "auth"
	
	connection, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@/%s", login, password, db)), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}	

	DB = connection

	connection.AutoMigrate(&models.User{})
}