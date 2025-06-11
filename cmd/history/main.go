package main

import (
	"event-driven-sample/internal/history"
	"event-driven-sample/pkg/config"
	"event-driven-sample/pkg/mongodb"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	repository, err := mongodb.New(cfg.DBHost, cfg.DBDatabase, cfg.DBUsername, cfg.DBPassword)
	if err != nil {
		log.Println(err)
		return
	}
	defer repository.Close()

	service := history.NewService(repository)

	consumer, err := history.NewConsumer(service, []string{cfg.KafkaBroker}, cfg.KafkaEngineTopic)
	if err != nil {
		log.Println(err)
		return
	}
	defer consumer.Close()

	if err := consumer.Listen(); err != nil {
		log.Println(err)
	}
}
