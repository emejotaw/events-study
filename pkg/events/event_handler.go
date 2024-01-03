package events

type EventHandler interface {
	Consume() error
	Publish(body []byte) error
}
