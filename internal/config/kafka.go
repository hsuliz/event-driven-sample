package config

type KafkaConfig struct {
	Brokers []string
	Topic   string
}

func GetKafkaConfig() KafkaConfig {
	return KafkaConfig{
		Brokers: []string{"localhost:29092"},
		Topic:   "test-topic",
	}
}
