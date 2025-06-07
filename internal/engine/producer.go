package engine

import (
	"encoding/json"
	"event-driven-sample/pkg/kafka"
	"log"
)

type Producer struct {
	Brokers       []string
	Topic         string
	KafkaProducer *kafka.Producer
}

func NewProducer(brokers []string, topic string) (*Producer, error) {
	kafkaProducer, err := kafka.NewProducer(brokers, topic)
	if err != nil {
		return nil, err
	}
	return &Producer{
		brokers,
		topic,
		kafkaProducer,
	}, nil
}

func (p Producer) SendMessage(determinant int) error {
	data, err := json.Marshal(determinant)
	if err != nil {
		log.Fatal(err)
	}
	if err := p.KafkaProducer.SendMessage(string(data)); err != nil {
		return err
	}
	return nil
}

func (p Producer) Close() {
	p.KafkaProducer.Close()
}
