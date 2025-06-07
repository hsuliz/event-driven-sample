package config

type KafkaConfig struct {
	Brokers      []string
	OrderTopic   string
	HistoryTopic string
}

func GetKafkaConfig() KafkaConfig {
	return KafkaConfig{
		Brokers:      []string{"localhost:29092"},
		OrderTopic:   "order-topic",
		HistoryTopic: "history-topic",
	}
}
