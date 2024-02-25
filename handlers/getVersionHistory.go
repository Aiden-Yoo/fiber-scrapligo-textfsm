package handlers

import (
	"github.com/aiden-yoo/getthemall/database"
	"github.com/gofiber/fiber/v2"
)

func GetVersionHistory(c *fiber.Ctx) error {
	db := database.DB
	Versions := []Version{}
	db.Model(&Version{}).Find(&Versions)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
		"data":    Versions,
	})
}
