package utilities

import "github.com/gofiber/fiber/v2"

type ResponseOptions struct {
	Context *fiber.Ctx
	Data    interface{}
	Info    string
	Status  int
}
