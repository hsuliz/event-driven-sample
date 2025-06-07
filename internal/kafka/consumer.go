package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
)

type Consumer struct {
	Brokers                 []string
	Topic                   string
	SaramaConsumer          sarama.Consumer
	SaramaPartitionConsumer sarama.PartitionConsumer
}

func NewConsumer(brokers []string, topic string) (*Consumer, error) {
	config := sarama.NewConfig()
	saramaConsumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create kafka consumer client: %w", err)
	}

	partitionConsumer, err := saramaConsumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		return nil, fmt.Errorf("failed to start partition consumer: %w", err)
	}

	return &Consumer{
		brokers,
		topic,
		saramaConsumer,
		partitionConsumer,
	}, nil
}

func (c Consumer) Listen() <-chan *sarama.ConsumerMessage {
	return c.SaramaPartitionConsumer.Messages()
}

func (c Consumer) Close() {
	if c.SaramaPartitionConsumer != nil {
		c.SaramaPartitionConsumer.Close()
	}
	if c.SaramaConsumer != nil {
		c.SaramaConsumer.Close()
	}
}
