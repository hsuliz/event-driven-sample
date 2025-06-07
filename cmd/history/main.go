package main

import (
	"event-driven-sample/internal/config"
	"event-driven-sample/internal/history"
	"log"
)

func main() {
	kafkaConfig := config.GetKafkaConfig()
	consumer, err := history.NewConsumer(kafkaConfig.Brokers, kafkaConfig.HistoryTopic)
	if err != nil {
		log.Println(err)
		return
	}
	defer consumer.Close()

	if err := consumer.Listen(); err != nil {
		log.Println(err)
	}
}
