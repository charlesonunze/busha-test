package main

import (
	"log"

	"github.com/charlesonunze/busha-test/database"
	"github.com/charlesonunze/busha-test/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB()
	defer database.CloseDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
