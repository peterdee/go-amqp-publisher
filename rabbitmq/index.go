package rabbitmq

import (
	ctx "context"
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Channel *amqp.Channel

var Queue amqp.Queue

func CreateConnection() {
	rabbitMQHost := os.Getenv("RABBITMQ_HOST")
	rabbitMQPassword := os.Getenv("RABBITMQ_PASSWORD")
	rabbitMQPort := os.Getenv("RABBITMQ_PORT")
	rabbitMQUser := os.Getenv("RABBITMQ_USER")
	if rabbitMQHost == "" || rabbitMQPassword == "" ||
		rabbitMQPort == "" || rabbitMQUser == "" {
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
	queue, queueError := channel.QueueDeclare(
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
	Channel = channel
	Queue = queue
}

func Publish(data []byte, contentType string) error {
	publishContext, cancel := ctx.WithTimeout(ctx.Background(), 5*time.Second)
	defer cancel()

	return Channel.PublishWithContext(
		publishContext,
		"",
		Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: contentType,
			Body:        data,
		},
	)
}
