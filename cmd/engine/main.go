package main

import (
	"event-driven-sample/internal/engine"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func main() {
	kafkaBroker, exists := os.LookupEnv("KAFKA_BROKER")
	if !exists {
		log.Fatal("can't find KAFKA_BROKER")
	}
	kafkaEngineTopic, exists := os.LookupEnv("KAFKA_ENGINE_TOPIC")
	if !exists {
		log.Fatal("can't find KAFKA_ENGINE_TOPIC")
	}
	kafkaOrderTopic, exists := os.LookupEnv("KAFKA_ORDER_TOPIC")
	if !exists {
		log.Fatal("can't find KAFKA_ORDER_TOPIC")
	}

	kafkaBrokers := []string{kafkaBroker}
	producer, err := engine.NewProducer(kafkaBrokers, kafkaEngineTopic)
	if err != nil {
		log.Println(err)
		return
	}
	defer producer.Close()

	service := engine.NewService(producer)
	consumer, err := engine.NewConsumer(service, kafkaBrokers, kafkaOrderTopic)
	if err != nil {
		log.Println(err)
		return
	}
	defer consumer.Close()

	if err := consumer.Listen(); err != nil {
		log.Println(err)
	}
}
