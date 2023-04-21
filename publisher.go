package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	connection, err := amqp.Dial(os.Getenv("RABBITMQ_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	log.Println("Successfully connected to RabbitMQ")

	// opening a channel over the connection established to interact with RabbitMQ
	channel, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()

	// declaring queue with its properties over the the channel opened
	queue, err := channel.QueueDeclare(
		"testing", // name
		false,     // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // args
	)
	if err != nil {
		log.Fatal(err)
	}

	// publishing a message
	err = channel.Publish(
		"",        // exchange
		"testing", // key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Test Message"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Queue status:", queue)
	log.Println("Successfully published message")
}
