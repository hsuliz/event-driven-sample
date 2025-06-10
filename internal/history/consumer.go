package history

import (
	"encoding/json"
	"event-driven-sample/pkg/entity"
	"event-driven-sample/pkg/kafka"
	"log"
)

type Consumer struct {
	Service       *Service
	KafkaConsumer *kafka.Consumer
}

func NewConsumer(service *Service, brokers []string, topic string) (*Consumer, error) {
	kafkaConsumer, err := kafka.NewConsumer(brokers, topic)
	if err != nil {
		return nil, err
	}
	return &Consumer{service, kafkaConsumer}, nil
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

		if err := c.Service.SaveCalculation(entity.Calculation{
			Hash:  engineMsg.Hash,
			Done:  engineMsg.Done,
			Value: engineMsg.Value,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (c Consumer) Close() {
	c.KafkaConsumer.Close()
}
