package router

import (
	"github.com/aiden-yoo/getthemall/handlers"
	"github.com/aiden-yoo/getthemall/middleware"
	"github.com/gofiber/fiber/v2"
)

func Initalize(router *fiber.App) {

	router.Use(middleware.Security)

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World!")
	})

	router.Use(middleware.Json)

	devices := router.Group("/devices")
	devices.Post("/", handlers.ScrapDevice)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})

}
