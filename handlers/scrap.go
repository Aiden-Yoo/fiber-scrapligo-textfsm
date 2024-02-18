package handlers

import (
	"github.com/aiden-yoo/getthemall/model"
	"github.com/aiden-yoo/getthemall/util/scrap"
	"github.com/gofiber/fiber/v2"
)

type Device model.Device

func ScrapDevice(c *fiber.Ctx) error {
	// db := database.DB
	// json := new(Device)
	// if err := c.BodyParser(json); err != nil {
	// 	return c.JSON(fiber.Map{
	// 		"code":    400,
	// 		"message": "Invalid JSON",
	// 	})
	// }
	scrap.Scrap()

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
}
