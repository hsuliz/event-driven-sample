package main

import (
	"event-driven-sample/internal/config"
	"event-driven-sample/internal/history"
	"log"
)

func main() {
	kafkaConfig := config.GetKafkaConfig()
	consumer, err := history.NewConsumer(kafkaConfig.Brokers, kafkaConfig.HistoryTopic)
	defer consumer.Close()

	if err != nil {
		log.Println(err)
		return
	}

	if err := consumer.Listen(); err != nil {
		log.Println(err)
		return
	}
}
