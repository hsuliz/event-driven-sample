package main

import (
	"event-driven-sample/internal/history"
	"event-driven-sample/pkg/mongodb"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func main() {
	dbHost, exists := os.LookupEnv("MONGODB_HOST")
	if !exists {
		log.Fatal("can't find MONGODB_HOST")
	}
	dbDatabase, exists := os.LookupEnv("MONGODB_DATABASE")
	if !exists {
		log.Fatal("can't find MONGODB_DATABASE")
	}
	dbUsername, exists := os.LookupEnv("MONGODB_USER")
	if !exists {
		log.Fatal("can't find MONGODB_USER")
	}
	dbPassword, exists := os.LookupEnv("MONGODB_PASSWORD")
	if !exists {
		log.Fatal("can't find MONGODB_PASSWORD")
	}

	kafkaBroker, exists := os.LookupEnv("KAFKA_BROKER")
	if !exists {
		log.Fatal("can't find KAFKA_BROKER")
	}
	kafkaEngineTopic, exists := os.LookupEnv("KAFKA_ENGINE_TOPIC")
	if !exists {
		log.Fatal("can't find KAFKA_ENGINE_TOPIC")
	}

	repository, err := mongodb.New(dbHost, dbDatabase, dbUsername, dbPassword)
	if err != nil {
		log.Println(err)
		return
	}
	defer repository.Close()

	service := history.NewService(repository)

	consumer, err := history.NewConsumer(service, []string{kafkaBroker}, kafkaEngineTopic)
	if err != nil {
		log.Println(err)
		return
	}
	defer consumer.Close()

	if err := consumer.Listen(); err != nil {
		log.Println(err)
	}
}
