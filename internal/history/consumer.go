package history

import (
	"encoding/json"
	"event-driven-sample/internal/kafka"
	"log"
)

type Consumer struct {
	KafkaConsumer *kafka.Consumer
}

func NewConsumer(brokers []string, topic string) (*Consumer, error) {
	kafkaConsumer, err := kafka.NewConsumer(brokers, topic)
	if err != nil {
		return nil, err
	}
	return &Consumer{kafkaConsumer}, nil
}

func (c Consumer) Listen() error {
	log.Println("listening for messages...")
	log.Println("listening for messages...")
	for message := range c.KafkaConsumer.Listen() {
		var determinant int
		if err := json.Unmarshal(message.Value, &determinant); err != nil {
			log.Printf("failed to unmarshal message: %v", err)
			continue
		}
		log.Println("got message:", determinant)
	}
	return nil
}

func (c Consumer) Close() {
	c.KafkaConsumer.Close()
}
