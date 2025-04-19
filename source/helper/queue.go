package helper

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func DeclareAndPublish(channel *amqp.Channel, queueName string, requestMessageEncoded []byte) {
	_, err := channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // args
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("## QUEUE PUBLISH ##")
	fmt.Println("Send to queue : " + queueName)
	fmt.Printf("Encoded length : %d\n", len(requestMessageEncoded))
	err = channel.Publish(
		"", // exchange
		queueName,
		false, // mandatoryw
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        requestMessageEncoded,
		},
	)
	if err != nil {
		panic(err)
	}
}

func DeclareAndConsume(channel *amqp.Channel, queueName string) (msgs <-chan amqp.Delivery, err error) {
	_, err = channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // args
	)

	if err != nil {
		panic(err)
	}

	channel.Qos(1, 0, false)
	return channel.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       //args
	)
}
