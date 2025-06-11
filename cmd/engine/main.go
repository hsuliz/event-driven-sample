package main

import (
	"event-driven-sample/internal/engine"
	"event-driven-sample/pkg/config"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	kafkaBrokers := []string{cfg.KafkaBroker}
	producer, err := engine.NewProducer(kafkaBrokers, cfg.KafkaEngineTopic)
	if err != nil {
		log.Println(err)
		return
	}
	defer producer.Close()

	service := engine.NewService(producer)
	consumer, err := engine.NewConsumer(service, kafkaBrokers, cfg.KafkaOrderTopic)
	if err != nil {
		log.Println(err)
		return
	}
	defer consumer.Close()

	if err := consumer.Listen(); err != nil {
		log.Println(err)
	}
}
