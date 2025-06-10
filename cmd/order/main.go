package main

import (
	"event-driven-sample/internal/order"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func main() {
	kafkaBroker, exists := os.LookupEnv("KAFKA_BROKER")
	if !exists {
		log.Fatal("can't find KAFKA_BROKER")
	}
	kafkaOrderTopic, exists := os.LookupEnv("KAFKA_ORDER_TOPIC")
	if !exists {
		log.Fatal("can't find KAFKA_ORDER_TOPIC")
	}

	service := order.NewService()

	producer, err := order.NewProducer(service, []string{kafkaBroker}, kafkaOrderTopic)
	if err != nil {
		log.Println(err)
		return
	}
	defer producer.Close()

	n := 3
	if err := producer.Produce(n); err != nil {
		log.Println(err)
	}
}
