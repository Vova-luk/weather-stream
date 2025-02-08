package kafka

import (
	"github.com/IBM/sarama"
)

func NewConsumerGroup(brokers []string, group_id string) (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()

	consumerGroup, err := sarama.NewConsumerGroup(brokers, group_id, config)
	if err != nil {
		return nil, err
	}

	return consumerGroup, nil
}
