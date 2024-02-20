package utilities

import (
	"github.com/gofiber/fiber/v2"
	"github.com/julyskies/gohelpers"
)

func Response(options ResponseOptions) error {
	info := options.Info
	if info == "" {
		info = "OK"
	}
	status := options.Status
	if status == 0 {
		status = fiber.StatusOK
	}
	response := fiber.Map{
		"datetime": gohelpers.MakeTimestamp(),
		"info":     info,
		"request":  options.Context.OriginalURL() + " [" + options.Context.Method() + "]",
		"status":   status,
	}
	if options.Data != nil {
		response["data"] = options.Data
	}
	return options.Context.Status(status).JSON(response)
}
