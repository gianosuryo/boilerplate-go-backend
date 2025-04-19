package config

import (
	"os"
	"sarana-dafa-ai-service/storage/env"

	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenQueueConnection() (*amqp.Connection, error) {
	rabbitMQHost := os.Getenv(env.RABBITMQ_HOST)
	rabbitMQPort := os.Getenv(env.RABBITMQ_PORT)
	rabbitMQCredential := os.Getenv(env.RABBITMQ_USER) + ":" + os.Getenv(env.RABBITMQ_PASSWORD)

	return amqp.Dial("amqp://" + rabbitMQCredential + "@" + rabbitMQHost + ":" + rabbitMQPort + "/")
}

func InitChannelConnection() *amqp.Channel {
	conn, err := OpenQueueConnection()
	if err != nil {
		panic(err) // TODO : Set to logrus
	}

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return channel
}
