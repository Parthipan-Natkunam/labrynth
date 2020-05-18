package main

import (
	"github.com/Parthipan-Natkunam/labrynth/config"
	"github.com/Parthipan-Natkunam/labrynth/controllers"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get(config.APIROOT+"/list", controllers.GetDefault)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(config.PORT)
}
