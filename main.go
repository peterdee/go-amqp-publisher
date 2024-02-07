package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
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

	app.Get("/", func(context *fiber.Ctx) error {
		s := fmp.FMP.New(os.Getenv("FMP_API_KEY"), os.Getenv("FMP_ENDPOINT"))
		fmt.Println(s)
		return context.Status(200).Send([]byte("OK"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = constants.PORT
	}
	app.Listen(fmt.Sprintf(":%s", port))
}
