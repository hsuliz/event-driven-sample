package main

import (
	"event-driven-sample/internal/config"
	"event-driven-sample/internal/engine"
)

func main() {
	service := engine.NewService()

	kafkaConfig := config.GetKafkaConfig()
	handler := engine.NewConsumer(
		service,
		kafkaConfig.Brokers,
		kafkaConfig.Topic,
	)
	handler.Consume()
}
