package controllers

import (
	"github.com/gofiber/fiber/v2"

	"go-amqp-publisher/fmp"
	"go-amqp-publisher/utilities"
)

func QuotesList(context *fiber.Ctx) error {
	data, dataError := fmp.FMP.GetStocks()
	if dataError != nil {
		return utilities.Response(utilities.ResponseOptions{
			Context: context,
			Info:    "COULD_NOT_LOAD_STOCK_LIST",
			Status:  fiber.StatusBadRequest,
		})
	}

	return utilities.Response(utilities.ResponseOptions{
		Context: context,
		Data:    data,
	})
}
