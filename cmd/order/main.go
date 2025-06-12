package main

import (
	"event-driven-sample/internal/order"
	"event-driven-sample/pkg/config"
	"event-driven-sample/pkg/kafka"
	"event-driven-sample/pkg/mongodb"
	"log"
	"math/rand"
	"time"
)

func main() {
	dbCfg := config.LoadDBConfig()
	repository, err := mongodb.New(dbCfg.Host, dbCfg.Database, dbCfg.Username, dbCfg.Password)
	if err != nil {
		log.Println(err)
		return
	}
	defer repository.Close()

	kafkaCfg := config.LoadKafkaConfig()
	producer, err := kafka.NewProducer([]string{kafkaCfg.Broker}, kafkaCfg.OrderTopic)
	if err != nil {
		log.Println(err)
		return
	}
	defer producer.Close()

	service := order.NewService(repository, producer)

	go func() {
		n := 3
		for {
			matrix := service.GenerateMatrix(n)
			err := service.ProcessCalculation(matrix)
			if err != nil {
				log.Println(err)
			}

			duration := rand.Intn(10)
			time.Sleep(time.Second * time.Duration(duration))
		}
	}()

	consumer, err := order.NewConsumer(service, []string{kafkaCfg.Broker}, kafkaCfg.EngineTopic)
	if err != nil {
		log.Println(err)
	}
	defer consumer.Close()

	if err := consumer.Listen(); err != nil {
		log.Println(err)
	}
}
