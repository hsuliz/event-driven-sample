package engine

import (
	"encoding/json"
	"event-driven-sample/pkg/kafka"
	"log"
	"math/rand"
	"time"
)

type Consumer struct {
	Service       *Service
	KafkaConsumer *kafka.Consumer
}

func NewConsumer(service *Service, brokers []string, topic string) (*Consumer, error) {
	client, err := kafka.NewConsumer(brokers, topic)
	if err != nil {
		return nil, err
	}
	return &Consumer{service, client}, nil
}

func (c Consumer) Listen() error {
	log.Println("listening for messages...")
	for message := range c.KafkaConsumer.Listen() {
		var matrix [][]int
		if err := json.Unmarshal(message.Value, &matrix); err != nil {
			log.Printf("failed to unmarshal message: %v", err)
			continue
		}
		determinant, err := c.Service.Calculate(matrix)
		if err != nil {
			log.Printf("failed to calculate message: %v", err)
			continue
		}

		// mock intense work
		duration := rand.Intn(10)
		time.Sleep(time.Duration(duration) * time.Second)

		if err := c.Service.Process(determinant); err != nil {
			return err
		}
	}
	return nil
}

func (c Consumer) Close() {
	c.KafkaConsumer.Close()
}
