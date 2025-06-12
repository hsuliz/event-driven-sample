package config

import (
	"log"
	"os"
)

type KafkaConfig struct {
	Broker      string
	EngineTopic string
	OrderTopic  string
}

type DBConfig struct {
	Host     string
	Database string
	Username string
	Password string
}

func LoadKafkaConfig() *KafkaConfig {
	return &KafkaConfig{
		Broker:      getEnv("KAFKA_BROKER", "localhost:29092"),
		EngineTopic: getEnv("KAFKA_ENGINE_TOPIC", "engine-topic"),
		OrderTopic:  getEnv("KAFKA_ORDER_TOPIC", "order-topic"),
	}
}

func LoadDBConfig() *DBConfig {
	return &DBConfig{
		Host:     getEnv("MONGODB_HOST", "localhost:27017"),
		Database: getEnv("MONGODB_DATABASE", "event-driven"),
		Username: getEnv("MONGODB_USERNAME", "username"),
		Password: getEnv("MONGODB_PASSWORD", "password"),
	}
}

func getEnv(key, fallback string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		log.Println("can't find", key, "using fallback value", fallback)
		return fallback
	}
	return val
}
