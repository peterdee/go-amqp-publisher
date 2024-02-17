package main

import (
	ctx "context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"

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

	// Initialize FMP service
	fmpApiKey := os.Getenv("FMP_API_KEY")
	fmpEndpoint := os.Getenv("FMP_ENDPOINT")
	if fmpApiKey == "" || fmpEndpoint == "" {
		log.Fatal("Could not load FMP configuration")
	}
	fmp.FMP.New(fmpApiKey, fmpEndpoint)

	// Connect to RabbitMQ
	rabbitMQHost := os.Getenv("RABBITMQ_HOST")
	rabbitMQPassword := os.Getenv("RABBITMQ_PASSWORD")
	rabbitMQPort := os.Getenv("RABBITMQ_PORT")
	rabbitMQUser := os.Getenv("RABBITMQ_USER")
	if rabbitMQPassword == "" || rabbitMQUser == "" {
		log.Fatal("Could not load RabbitMQ configuration")
	}
	rabbitMQConnection, connectionError := amqp.Dial(
		fmt.Sprintf(
			"amqp://%s:%s@%s:%s/",
			rabbitMQUser,
			rabbitMQPassword,
			rabbitMQHost,
			rabbitMQPort,
		),
	)
	if connectionError != nil {
		log.Fatal("Could not connect to RabbitMQ:", connectionError)
	}
	channel, channelError := rabbitMQConnection.Channel()
	if channelError != nil {
		log.Fatal(channelError)
	}
	rabbitMQQueue, queueError := channel.QueueDeclare(
		"quotes",
		false,
		false,
		false,
		false,
		nil,
	)
	if queueError != nil {
		log.Fatal(queueError)
	}

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
		publishContext, cancel := ctx.WithTimeout(ctx.Background(), 5*time.Second)
		defer cancel()
		preparedData, parsingError := json.Marshal(data)
		if parsingError != nil {
			return context.Status(500).JSON(fiber.Map{
				"info":   "INTERNAL_SERVER_ERROR",
				"status": 500,
			})
		}
		publishError := channel.PublishWithContext(
			publishContext,
			"",
			rabbitMQQueue.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        preparedData,
			},
		)
		if publishError != nil {
			return context.Status(500).JSON(fiber.Map{
				"info":   "INTERNAL_SERVER_ERROR",
				"status": 500,
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
