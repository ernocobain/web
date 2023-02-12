package main

import (
	"dhikrama.com/web/src"
	"dhikrama.com/web/src/database"
)

func main() {
	src.MainSrc()
	database.ConnectDB()
}
