package main

import (
	"event-driven-sample/internal/config"
	"event-driven-sample/internal/order"
	"log"
)

func main() {
	service := order.NewService()

	kafkaConfig := config.GetKafkaConfig()
	matrixN := 3
	producer := order.NewProducer(
		service,
		kafkaConfig.Brokers,
		kafkaConfig.Topic,
	)
	log.Fatalln(producer.Produce(matrixN))

}
