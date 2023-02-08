package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error // define error here to prevent overshadowing the global DB

	env := "admin:Dat4base@tcp(mariadb-1.cawqtwhwwb0c.ap-southeast-1.rds.amazonaws.com)/?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(env), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

}
