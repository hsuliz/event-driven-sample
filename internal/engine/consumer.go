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
	for msg := range c.KafkaConsumer.Listen() {
		// TODO add waiting when exit
		go func() {
			var matrix [][]int
			if err := json.Unmarshal(msg.Value, &matrix); err != nil {
				log.Printf("failed to unmarshal message: %v", err)
			}

			matrixHash, determinant, err := c.Service.Calculate(matrix)
			if err != nil {
				log.Printf("failed to calculate message: %v", err)
			}

			// mock intense work
			duration := rand.Intn(10)
			time.Sleep(time.Duration(duration) * time.Second)

			if err := c.Service.Process(matrixHash, determinant); err != nil {
				log.Printf("failed to process message: %v", err)
			}
		}()
	}
	return nil
}

func (c Consumer) Close() {
	c.KafkaConsumer.Close()
}
