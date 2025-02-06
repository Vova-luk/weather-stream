package kafka

import (
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

func NewProduces(brokers []string, topic string, log *logrus.Logger) (sarama.SyncProducer, error) {
	if err := ensureTopicExists(brokers, topic, log); err != nil {
		return nil, err
	}
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}

func NewConsumer(brokers []string, groupId string) (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupId, config)
	if err != nil {
		return nil, err
	}

	return consumerGroup, nil
}

func ensureTopicExists(brokers []string, topic string, log *logrus.Logger) error {
	config := sarama.NewConfig()

	admin, err := sarama.NewClusterAdmin(brokers, config)
	if err != nil {
		return err
	}
	defer admin.Close()

	topiks, err := admin.ListTopics()
	if err != nil {
		return nil
	}

	if _, ok := topiks[topic]; !ok {
		err := admin.CreateTopic(topic, &sarama.TopicDetail{
			NumPartitions:     3,
			ReplicationFactor: 1,
		}, false)
		if err != nil {
			return err
		}
		log.Printf("Created topic: %s", topic)
	} else {
		log.Printf("Topic %s already exists", topic)
	}
	return nil

}
