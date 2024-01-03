package rabbitmq

import (
	"context"
	"fmt"
	"log"

	"github.com/emejotaw/events/pkg/dto"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	channel    *amqp.Channel
	exchange   string
	routingKey string
}

func NewRabbitMQ(mqConfig *dto.RabbitMQDTO) (*RabbitMQ, error) {

	rabbitMQ := &RabbitMQ{
		exchange:   mqConfig.Exchange,
		routingKey: mqConfig.RoutingKey,
	}
	err := rabbitMQ.connect(mqConfig)

	if err != nil {
		return nil, err
	}

	return rabbitMQ, nil
}

func (r *RabbitMQ) connect(mqConfig *dto.RabbitMQDTO) error {

	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		mqConfig.Username,
		mqConfig.Password,
		mqConfig.Host,
		mqConfig.Port,
	)
	log.Printf("%v", dsn)
	conn, err := amqp.Dial(dsn)

	if err != nil {
		log.Printf("could not establish rabbitmq connection, error: %v", err)
		return err
	}

	channel, err := conn.Channel()

	if err != nil {
		log.Printf("could not connect with channel, error: %v", err)
		return err
	}

	r.channel = channel
	return nil
}

func (r *RabbitMQ) Consume(queueName string, eventch chan amqp.Delivery) error {

	messagesch, err := r.channel.Consume(
		queueName,
		"event-consumer",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for message := range messagesch {
		eventch <- message
	}

	return nil
}

func (r *RabbitMQ) Publish(body []byte) error {

	ctx := context.Background()
	return r.channel.PublishWithContext(
		ctx,
		r.exchange,
		r.routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
