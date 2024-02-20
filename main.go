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
	"go-amqp-publisher/controllers"
	"go-amqp-publisher/fmp"
	"go-amqp-publisher/rabbitmq"
)

func main() {
	envError := godotenv.Load()
	if envError != nil {
		log.Fatal("Error loading .env file!")
	}

	fmp.FMP.New()
	rabbitmq.CreateConnection()

	app := fiber.New()
	app.Use(favicon.New(favicon.Config{
		File: "./assets/favicon.ico",
		URL:  "/favicon.ico",
	}))
	app.Use(logger.New())

	app.Get("/", controllers.IndexController)
	app.Get("/api", controllers.IndexController)
	app.Get("/api/quote/:quote", controllers.QuoteController)
	app.Get("/api/quotes/list", controllers.QuotesList)

	port := os.Getenv("PORT")
	if port == "" {
		port = constants.PORT
	}
	if launchError := app.Listen(fmt.Sprintf(":%s", port)); launchError != nil {
		log.Fatal(launchError)
	}
}
