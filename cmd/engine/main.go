package main

import (
	"event-driven-sample/internal/engine"
	"event-driven-sample/pkg/config"
	"event-driven-sample/pkg/kafka"
	"log"
)

func main() {
	kafkaCfg := config.LoadKafkaConfig()

	kafkaBrokers := []string{kafkaCfg.Broker}
	producer, err := kafka.NewProducer(kafkaBrokers, kafkaCfg.EngineTopic)
	if err != nil {
		log.Println(err)
		return
	}
	defer producer.Close()

	service := engine.NewService(producer)
	consumer, err := engine.NewConsumer(service, kafkaBrokers, kafkaCfg.OrderTopic)
	if err != nil {
		log.Println(err)
		return
	}
	defer consumer.Close()

	if err := consumer.Listen(); err != nil {
		log.Println(err)
	}
}
