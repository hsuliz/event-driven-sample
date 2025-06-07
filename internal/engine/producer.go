package engine

import (
	"encoding/json"
	"event-driven-sample/pkg/kafka"
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

func (p Producer) SendMessage(message kafka.EngineMsg) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	if err := p.KafkaProducer.SendMessage(data); err != nil {
		return err
	}
	return nil
}

func (p Producer) Close() {
	p.KafkaProducer.Close()
}
