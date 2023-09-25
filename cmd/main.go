package main

import (
	"github.com/gofiber/fiber"
)

func main() {

	 app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
    c.SendString("Hello, World!")
})

    app.Listen(":3000")
}
