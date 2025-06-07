package main

import (
	"event-driven-sample/internal/config"
	"event-driven-sample/internal/order"
	"log"
)

func main() {
	service := order.NewService()
	kafkaConfig := config.GetKafkaConfig()
	producer, err := order.NewProducer(service, kafkaConfig.Brokers, kafkaConfig.OrderTopic)
	defer producer.Close()
	if err != nil {
		log.Println(err)
		return
	}
	n := 3
	if err := producer.Produce(n); err != nil {
		log.Println(err)
		return
	}
}
