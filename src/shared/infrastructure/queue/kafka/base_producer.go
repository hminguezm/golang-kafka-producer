package kafka

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
	"reflect"
	"time"
	"wrk-connector/src/shared/domain/service"
	log "wrk-connector/src/shared/infrastructure/config"
)

type producer struct {
	brokers      []string
	dialer       *kafka.Dialer
	kafkaWriters map[string]*kafka.Writer
}

func NewKafkaProducer(kafkaDialer *kafka.Dialer, brokers ...string) *producer {
	return &producer{
		brokers:      brokers,
		dialer:       kafkaDialer,
		kafkaWriters: make(map[string]*kafka.Writer, 0),
	}
}

func (p *producer) getKafkaWriter(topic string) *kafka.Writer {
	if p.kafkaWriters[topic] == nil {
		p.kafkaWriters[topic] = kafka.NewWriter(kafka.WriterConfig{
			Balancer:         &kafka.LeastBytes{},
			BatchSize:        1,
			BatchTimeout:     10 * time.Millisecond,
			Brokers:          p.brokers,
			CompressionCodec: snappy.NewCompressionCodec(),
			Dialer:           p.dialer,
			ReadTimeout:      10 * time.Second,
			Topic:            topic,
			WriteTimeout:     10 * time.Second,
		})
	}
	return p.kafkaWriters[topic]
}

func (p *producer) Send(topic string, message interface{}) error {
	kafkaMessages, err := createKafkaMessages(message)
	if err != nil {
		return errors.New("error on publishing kafka message")
	}
	kafkaWriter := p.getKafkaWriter(topic)

	err = kafkaWriter.WriteMessages(context.Background(), kafkaMessages...)
	if err != nil {
		log.WithError(err).Error("error writing kafka message")
		return err
	} else {
		log.WithField("topic", topic).Info("kafka message published successfully")
	}

	return nil
}

func createKafkaMessages(message interface{}) ([]kafka.Message, error) {
	var kafkaMessages []kafka.Message

	if service.IsNilFixed(message) {
		return nil, errors.New("kafka error the message is empty")
	}

	switch reflect.TypeOf(message).Kind() {
	case reflect.Array, reflect.Slice:
		value := reflect.ValueOf(message)
		for index := 0; index < value.Len(); index++ {
			kafkaMessages = append(kafkaMessages, createMessageKafka(value.Index(index)))
		}
	default:
		kafkaMessages = append(kafkaMessages, createMessageKafka(message))
	}

	return kafkaMessages, nil
}

func createMessageKafka(message interface{}) kafka.Message {
	payload := service.EntityToJson(message)
	key := uuid.New().String()
	kafkaMessage := kafka.Message{
		Key:   []byte(key),
		Value: []byte(payload),
		Time:  time.Now(),
	}

	return kafkaMessage
}

func (p *producer) Close(topic string) error {
	if p.kafkaWriters[topic] == nil {
		return errors.New("error trying to close kafka connection, connection does not exist")
	}
	_ = p.kafkaWriters[topic].Close()
	delete(p.kafkaWriters, topic)

	return nil
}
