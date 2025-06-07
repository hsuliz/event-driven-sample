package main

import (
	"event-driven-sample/internal/config"
	"event-driven-sample/internal/engine"
	"log"
)

func main() {
	kafkaConfig := config.GetKafkaConfig()
	producer, err := engine.NewProducer(kafkaConfig.Brokers, kafkaConfig.HistoryTopic)
	if err != nil {
		log.Fatalln(err)
	}

	service := engine.NewService(producer)
	consumer, err := engine.NewConsumer(service, kafkaConfig.Brokers, kafkaConfig.OrderTopic)
	if err != nil {
		log.Println(err)
	}

	if err := consumer.Listen(); err != nil {
		log.Println(err)
	}
	defer consumer.Close()
}
