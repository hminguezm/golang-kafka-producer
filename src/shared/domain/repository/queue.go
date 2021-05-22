package repository

type QueueProducer interface {
	Send(topic string, message interface{}) error
	Close(topic string) error
}
