package utilities

import "github.com/gofiber/fiber/v2"

func DefaultErrorHandler(context *fiber.Ctx, err error) error {
	info := "INTERNAL_SERVER_ERROR"
	status := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		info = e.Message
		status = e.Code
	}

	return Response(ResponseOptions{
		Context: context,
		Info:    info,
		Status:  status,
	})
}
