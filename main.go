package main

import (
	"log"
	"os"

	"github.com/aiden-yoo/getthemall/database"
	"github.com/aiden-yoo/getthemall/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	godotenv.Load()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	database.ConnectDB()

	router.Initalize(app)
	log.Fatal(app.Listen(":" + getenv("PORT", "3000")))
}
