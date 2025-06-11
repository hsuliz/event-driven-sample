package order

import (
	"encoding/json"
	"event-driven-sample/pkg/hash"
	"event-driven-sample/pkg/kafka"
	"log"
	"math/rand"
	"slices"
	"time"
)

type Producer struct {
	Service       *Service
	KafkaProducer *kafka.Producer
}

func NewProducer(service *Service, brokers []string, topic string) (*Producer, error) {

	client, err := kafka.NewProducer(brokers, topic)
	if err != nil {
		return nil, err
	}
	return &Producer{service, client}, nil
}

func (p Producer) Produce(matrix int) error {
	for {
		matrix := p.Service.GenerateMatrix(matrix)
		marshaledMatrix, err := json.Marshal(matrix)
		if err != nil {
			log.Fatal(err)
		}
		// just for hash i dont care
		flattenMatrixInt := slices.Concat(matrix...)
		flattenMatrixIntMarshalled, err := json.Marshal(flattenMatrixInt)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("sending message:", hash.Encode(flattenMatrixIntMarshalled))

		if err := p.KafkaProducer.SendMessage(marshaledMatrix); err != nil {
			return err
		}

		duration := rand.Intn(10)
		time.Sleep(time.Second * time.Duration(duration))
	}
	return nil
}

func (p Producer) Close() {
	p.KafkaProducer.Close()
}
