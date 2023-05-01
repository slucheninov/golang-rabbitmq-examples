package main

import (
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	start := time.Now()
	connection, err := amqp.Dial(os.Getenv("RABBITMQ_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	duration := time.Since(start)
	log.Println("Successfully connected to RabbitMQ", " - ", duration)

	// opening a channel over the connection established to interact with RabbitMQ
	channel, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()

	start = time.Now()
	// declaring consumer with its properties over channel opened
	msgs, err := channel.Consume(
		"testing", // queue
		"",        // consumer
		true,      // auto ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       //args
	)
	if err != nil {
		log.Fatal(err)
	}
	duration = time.Since(start)
	// print consumed messages from queue
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			log.Printf("Received Message: %s - %v\n", msg.Body, duration)
		}
	}()

	log.Println("Waiting for messages...")
	<-forever
}
