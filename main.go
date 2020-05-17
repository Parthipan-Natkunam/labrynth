package main

import (
	"github.com/Parthipan-Natkunam/labrynth/config"
	"github.com/gofiber/fiber"
)

func testHandler(ctx *fiber.Ctx) {
	ctx.Send("Fiber Set-up Successfully")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", testHandler)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(config.PORT)
}
