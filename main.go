package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"go-amqp-publisher/constants"
	"go-amqp-publisher/fmp"
)

func main() {
	envError := godotenv.Load()
	if envError != nil {
		log.Fatal("Error loading .env file!")
	}

	app := fiber.New()

	app.Use(favicon.New(favicon.Config{
		File: "./assets/favicon.ico",
		URL:  "/favicon.ico",
	}))
	app.Use(logger.New())

	fmpApiKey := os.Getenv("FMP_API_KEY")
	fmpEndpoint := os.Getenv("FMP_ENDPOINT")
	if fmpApiKey == "" || fmpEndpoint == "" {
		log.Fatal("Could not load FMP configuration")
	}

	// Initialize FMP service
	fmp.FMP.New(fmpApiKey, fmpEndpoint)

	app.Get("/api/quotes/list", func(context *fiber.Ctx) error {
		data, dataError := fmp.FMP.GetStocks()
		if dataError != nil {
			return context.Status(400).JSON(fiber.Map{
				"info":   "COULD_NOT_LOAD_STOCK_LIST",
				"status": 400,
			})
		}
		return context.Status(200).JSON(data)
	})

	app.Get("/api/quote/:quote", func(context *fiber.Ctx) error {
		quote := context.Params("quote", "")
		if quote == "" {
			return context.Status(400).JSON(fiber.Map{
				"info":   "MISSING_DATA",
				"status": 400,
			})
		}
		data, dataError := fmp.FMP.GetQuote(quote)
		if dataError != nil {
			return context.Status(400).JSON(fiber.Map{
				"info":   "COULD_NOT_LOAD_STOCK_DATA",
				"status": 400,
			})
		}
		return context.Status(200).JSON(data)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = constants.PORT
	}
	if launchError := app.Listen(fmt.Sprintf(":%s", port)); launchError != nil {
		log.Fatal(launchError)
	}
}
