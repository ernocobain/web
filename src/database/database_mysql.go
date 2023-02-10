package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error // define error here to prevent overshadowing the global DB

	const MYSQL = "admin:Dat4base@tcp(mariadb-1.cawqtwhwwb0c.ap-southeast-1.rds.amazonaws.com)/dhikrama?charset=utf8mb4&parseTime=True&loc=Local"
	env := MYSQL
	DB, err = gorm.Open(mysql.Open(env), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Database")

}
