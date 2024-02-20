package controllers

import (
	"github.com/gofiber/fiber/v2"

	"go-amqp-publisher/utilities"
)

func IndexController(context *fiber.Ctx) error {
	return utilities.Response(utilities.ResponseOptions{Context: context})
}
