package order

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

type Producer struct {
	Service *Service
	Brokers []string
	Topic   string
}

func NewProducer(
	service *Service,
	brokers []string,
	topic string,
) *Producer {
	return &Producer{
		service,
		brokers,
		topic,
	}
}

func (p Producer) Produce(matrixN int) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(
		p.Brokers,
		config,
	)
	if err != nil {
		return fmt.Errorf("failed to setup producer: %w", err)
	}

	for {
		matrix := p.Service.GenerateMatrix(matrixN)
		data, err := json.Marshal(matrix)
		if err != nil {
			log.Fatal(err)
		}

		msg := &sarama.ProducerMessage{
			Topic: p.Topic,
			Value: sarama.StringEncoder(data),
		}

		if _, _, err := producer.SendMessage(msg); err != nil {
			return fmt.Errorf("failed to setup producer: %w", err)
		}
	}

	return nil
}
