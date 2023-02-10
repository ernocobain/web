package migration

import (
	"fmt"
	"log"

	"dhikrama.com/web/src/database"
)

func RunMigration(model interface{}) {
	err := database.DB.Debug().AutoMigrate(model)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database migrated")
}
