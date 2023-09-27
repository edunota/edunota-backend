package handlers

import "github.com/gofiber/fiber/v2"

func GetHello(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello world")
}
