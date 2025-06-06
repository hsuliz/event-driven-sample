package config

type KafkaConfig struct {
	Brokers []string
	Topic   string
}

func GetKafkaConfig() KafkaConfig {
	return KafkaConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "test-topic",
	}
}
