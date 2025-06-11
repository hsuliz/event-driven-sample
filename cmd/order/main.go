package main

import (
	"event-driven-sample/internal/order"
	"event-driven-sample/pkg/config"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	service := order.NewService()
	producer, err := order.NewProducer(service, []string{cfg.KafkaBroker}, cfg.KafkaOrderTopic)
	if err != nil {
		log.Println(err)
		return
	}
	defer producer.Close()

	n := 3
	if err := producer.Produce(n); err != nil {
		log.Println(err)
	}
}
