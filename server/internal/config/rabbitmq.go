package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/streadway/amqp"
)

var Channel *amqp.Channel

func InitRabbitMQ() {
	amqpURL := os.Getenv("RABBITMQ_URL")
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}

	Channel = ch

	_, err = ch.QueueDeclare("notification_checkout", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}
}

func PublishCheckoutNotification(payload interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return Channel.Publish(
		"", // default exchange
		"notification_checkout",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
