package handlers

import (
	"fmt"

	"dhikrama.com/web/src/database"
	"dhikrama.com/web/src/database/migration"
	"dhikrama.com/web/src/entity"
	"github.com/gofiber/fiber/v2"
)

func UserGetAll(ctx *fiber.Ctx) error {
	var users []entity.UsersModel

	migration.RunMigration(&users)
	result := database.DB.Find(&users)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return ctx.JSON(users)
}
