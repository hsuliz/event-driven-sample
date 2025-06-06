package engine

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
)

type Consumer struct {
	Service *Service
	Brokers []string
	Topic   string
}

func NewConsumer(
	service *Service,
	brokers []string,
	topic string,
) *Consumer {
	return &Consumer{
		service,
		brokers,
		topic,
	}
}

func (c Consumer) Consume() {
	config := sarama.NewConfig()

	consumer, err := sarama.NewConsumer(c.Brokers, config)
	if err != nil {
		log.Fatalf("failed to start Kafka consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(
		c.Topic,
		0,
		sarama.OffsetNewest,
	)
	if err != nil {
		log.Fatalf("failed to start partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	log.Println("Listening for messages...")
	for message := range partitionConsumer.Messages() {
		var matrix [][]int
		if err := json.Unmarshal(message.Value, &matrix); err != nil {
			log.Fatalf("Failed to start Kafka consumer: %v", err)
		}
		log.Println(c.Service.CalculateDeterminant(matrix))
	}
}
