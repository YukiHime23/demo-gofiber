package main

import (
	"demo-gofiber/controllers"
	"demo-gofiber/database"
	"demo-gofiber/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"gorm.io/gorm"
	"log"
)

func main() {

	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")
	// If you want other engine, just replace with following
	// Create a new engine with django
	// engine := django.New("./views", ".django")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	db, err := database.OpenPostgresDB()
	if err != nil {
		log.Fatal("[main] DB connect error: ", err)
	}

	setupRoutesV1(app, db)

	app.Use(cors.New())

	// Initialize default config
	app.Use(logger.New())

	// Or extend your config for customization
	// Logging remote IP and Port
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	app.Listen(":8080")
}

func setupRoutesV1(app *fiber.App, db *gorm.DB) {
	ser := services.NewAppService(db)
	sCon := controllers.NewSampleController(ser)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/v1/sample", sCon.SampleFunc)

	app.Get("/render-html", sCon.RenderHtml)
}
