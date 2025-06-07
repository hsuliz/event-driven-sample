package order

import (
	"encoding/json"
	"event-driven-sample/internal/kafka"
	"log"
	"time"
)

type Producer struct {
	Service       *Service
	KafkaProducer *kafka.Producer
}

func NewProducer(service *Service, brokers []string, topic string) (*Producer, error) {
	client, err := kafka.NewProducer(brokers, topic)
	if err != nil {
		return nil, err
	}
	return &Producer{service, client}, nil
}

func (p Producer) Produce(matrix int) error {
	for {
		matrix := p.Service.GenerateMatrix(matrix)
		data, err := json.Marshal(matrix)
		if err != nil {
			log.Fatal(err)
		}
		if err := p.KafkaProducer.SendMessage(string(data)); err != nil {
			return err
		}
		log.Println("produced")
		// based on kafka resources sleep can be removed
		time.Sleep(time.Second)
	}
	return nil
}
