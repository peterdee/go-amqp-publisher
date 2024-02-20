package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"

	"go-amqp-publisher/fmp"
	"go-amqp-publisher/rabbitmq"
	"go-amqp-publisher/utilities"
)

func QuoteController(context *fiber.Ctx) error {
	quote := context.Params("quote", "")
	if quote == "" {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"MISSING_DATA",
		)
	}

	data, dataError := fmp.FMP.GetQuote(quote)
	if dataError != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"COULD_NOT_LOAD_STOCK_DATA",
		)
	}

	preparedData, parsingError := json.Marshal(data)
	if parsingError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	publishError := rabbitmq.Publish(preparedData, "application/json")
	if publishError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return utilities.Response(utilities.ResponseOptions{
		Context: context,
		Data:    data,
	})
}
