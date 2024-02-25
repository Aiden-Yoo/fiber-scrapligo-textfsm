package handlers

import (
	"github.com/aiden-yoo/getthemall/database"
	"github.com/aiden-yoo/getthemall/model"
	"github.com/aiden-yoo/getthemall/util/scrap"
	"github.com/gofiber/fiber/v2"
)

type Version model.Version

func GetVersion(c *fiber.Ctx) error {
	db := database.DB
	resutls, err := scrap.GetVersion()
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": err,
		})
	}

	if len(resutls) != 0 {
		for _, version := range resutls {
			newVersion := Version{
				Hostname: version["HOSTNAME"].(string),
				Platform: version["PLATFORM"].(string),
				Version:  version["OS"].(string),
				Uptime:   version["UPTIME"].(string),
			}
			err := db.Create(&newVersion).Error
			if err != nil {
				return c.SendStatus(fiber.StatusBadRequest)
			}
		}
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
}
