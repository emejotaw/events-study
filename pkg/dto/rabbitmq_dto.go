package dto

type RabbitMQDTO struct {
	Username   string
	Password   string
	Exchange   string
	RoutingKey string
	Host       string
	Port       int
}
