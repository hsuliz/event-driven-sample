package main

import (
	"event-driven-sample/internal/config"
	"event-driven-sample/internal/history"
	"event-driven-sample/pkg/mongodb"
	"log"
)

func main() {
	repository, err := mongodb.NewMongoDB("localhost:27017", "event-driven", "username", "password")
	if err != nil {
		panic(err)
	}
	defer repository.Close()

	service := history.NewService(repository)

	kafkaConfig := config.GetKafkaConfig()
	consumer, err := history.NewConsumer(service, kafkaConfig.Brokers, kafkaConfig.HistoryTopic)
	if err != nil {
		log.Println(err)
		return
	}
	defer consumer.Close()

	if err := consumer.Listen(); err != nil {
		log.Println(err)
	}
}
