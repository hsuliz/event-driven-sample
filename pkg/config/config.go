package config

import (
	"log"
	"os"
)

type Config struct {
	// KAFKA
	KafkaBroker      string
	KafkaEngineTopic string
	KafkaOrderTopic  string
	// MONGODB
	DBHost     string
	DBDatabase string
	DBUsername string
	DBPassword string
}

func LoadConfig() *Config {
	return &Config{
		KafkaBroker:      getEnv("KAFKA_BROKER", "localhost:29092"),
		KafkaEngineTopic: getEnv("KAFKA_ENGINE_TOPIC", "engine-topic"),
		KafkaOrderTopic:  getEnv("KAFKA_ORDER_TOPIC", "order-topic"),

		DBHost:     getEnv("MONGODB_HOST", "localhost:27017"),
		DBDatabase: getEnv("MONGODB_DATABASE", "event-driven"),
		DBUsername: getEnv("MONGODB_USERNAME", "username"),
		DBPassword: getEnv("MONGODB_PASSWORD", "password"),
	}
}

func getEnv(key, fallback string) string {
	if val, exists := os.LookupEnv(key); exists {
		log.Println("can't find", key, "using fallback value", fallback)
		return val
	}
	return fallback
}
