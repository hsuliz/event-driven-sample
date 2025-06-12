package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
)

type Producer struct {
	Brokers            []string
	Topic              string
	SaramaSyncProducer sarama.SyncProducer
}

func NewProducer(brokers []string, topic string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, fmt.Errorf("failed to setup producer: %w", err)
	}

	return &Producer{
		brokers,
		topic,
		producer,
	}, nil
}

func (p Producer) SendMessage(message []byte) error {
	producerMessage := &sarama.ProducerMessage{
		Topic: p.Topic,
		Value: sarama.ByteEncoder(message),
	}
	if _, _, err := p.SaramaSyncProducer.SendMessage(producerMessage); err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}

func (p Producer) Close() {
	p.SaramaSyncProducer.Close()
}
