package main

import (
	"encoding/json"
	"time"

	"github.com/emejotaw/events/pkg/dto"
	"github.com/emejotaw/events/pkg/events/rabbitmq"
)

func main() {

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

	events := []dto.Event{
		{Name: "evento-solar", Code: 0, Message: "Event 1", CreatedAt: time.Now()},
		{Name: "evento-maritimo", Code: 0, Message: "Event 2", CreatedAt: time.Now()},
		{Name: "evento-atomica", Code: 0, Message: "Event 3", CreatedAt: time.Now()},
		{Name: "evento-plasma", Code: 0, Message: "Event 4", CreatedAt: time.Now()},
		{Name: "evento-niobio", Code: 0, Message: "Event 5", CreatedAt: time.Now()},
		{Name: "evento-calcario", Code: 0, Message: "Event 6", CreatedAt: time.Now()},
		{Name: "evento-diamante", Code: 0, Message: "Event 7", CreatedAt: time.Now()},
		{Name: "evento-ruby", Code: 0, Message: "Event 8", CreatedAt: time.Now()},
		{Name: "evento-obsidian", Code: 0, Message: "Event 9", CreatedAt: time.Now()},
		{Name: "evento-madeira", Code: 0, Message: "Event 10", CreatedAt: time.Now()},
	}

	for _, event := range events {

		body, err := json.Marshal(event)

		if err != nil {
			panic(err)
		}
		err = rabbitMq.Publish(body)

		if err != nil {
			panic(err)
		}

		time.Sleep(time.Second * 2)
	}

}
