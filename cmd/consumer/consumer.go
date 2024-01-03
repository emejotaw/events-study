package main

import (
	"encoding/json"
	"log"

	"github.com/emejotaw/events/pkg/dto"
	"github.com/emejotaw/events/pkg/events/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	eventch := make(chan amqp.Delivery)
	queueName := "event-queue"
	mqConfig := &dto.RabbitMQDTO{
		Username: "root",
		Password: "root",
		Host:     "localhost",
		Port:     5672,
		Exchange: "amq.direct",
	}
	rabbitMq, err := rabbitmq.NewRabbitMQ(mqConfig)

	if err != nil {
		panic(err)
	}

	go rabbitMq.Consume(queueName, eventch)

	if err != nil {
		panic(err)
	}

	for message := range eventch {
		event := &dto.Event{}
		err := json.Unmarshal(message.Body, event)

		if err != nil {
			panic(err)
		}

		log.Printf("Event %s with received with message %s and code %d at %v\n",
			event.Name,
			event.Message,
			event.Code,
			event.CreatedAt,
		)

		message.Ack(false)
	}
}
