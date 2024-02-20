package controllers

import (
	ctx "context"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"

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

	publishContext, cancel := ctx.WithTimeout(ctx.Background(), 5*time.Second)
	defer cancel()

	preparedData, parsingError := json.Marshal(data)
	if parsingError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	publishError := rabbitmq.Channel.PublishWithContext(
		publishContext,
		"",
		rabbitmq.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        preparedData,
		},
	)

	if publishError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return utilities.Response(utilities.ResponseOptions{
		Context: context,
		Data:    data,
	})
}
