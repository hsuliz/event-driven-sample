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

func (p Producer) SendMessage(msg kafka.EngineMsg) error {
	log.Printf("sending message: %v", msg)
	marshaledMsg, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	if err := p.KafkaProducer.SendMessage(marshaledMsg); err != nil {
		return err
	}
	return nil
}

func (p Producer) Close() {
	p.KafkaProducer.Close()
}
