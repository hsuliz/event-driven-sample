package history

import (
	"encoding/json"
	"event-driven-sample/pkg/kafka"
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
	for msg := range c.KafkaConsumer.Listen() {
		var engineMsg kafka.EngineMsg
		if err := json.Unmarshal(msg.Value, &engineMsg); err != nil {
			log.Printf("failed to unmarshal message: %v", err)
			continue
		}
		log.Println("got message:", engineMsg)
	}
	return nil
}

func (c Consumer) Close() {
	c.KafkaConsumer.Close()
}
