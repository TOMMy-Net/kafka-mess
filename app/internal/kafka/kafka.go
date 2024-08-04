package kafka

import (

	"errors"
	"fmt"


	"github.com/IBM/sarama"
)

type Broker struct {
	Producer  sarama.SyncProducer
	Consumer  sarama.ConsumerGroup
	Hosts     []string
	Topic     string
	Partition int
	Reconect  chan struct{}
}

var (
	ErrWithWrite = errors.New("the message was delivered but there was a problem on the intermediate node")
	ErrWithRead  = errors.New("error reading a message from kafka")
)




func Connect(addr []string, topic string) (*Broker, error) {
	var broker = new(Broker)

	producer, err := setupProducer(addr...)
	if err != nil {
		return &Broker{}, err
	}
	broker.Hosts = addr
	broker.Producer = producer
	broker.Topic = topic

	return broker, nil
}

func setupProducer(addr ...string) (sarama.SyncProducer, error) {
	config := newProducerConfig()

	producer, err := sarama.NewSyncProducer(addr,
		config)
	if err != nil {
		return nil, fmt.Errorf("failed to setup producer: %w", err)
	}
	return producer, nil
}

func initializeConsumerGroup(groupId string, addr ...string) (sarama.ConsumerGroup, error) {
    config := sarama.NewConfig()

    consumerGroup, err := sarama.NewConsumerGroup(
        addr, groupId, config)
    if err != nil {
        return nil, fmt.Errorf("failed to initialize consumer group: %w", err)
    }

    return consumerGroup, nil
}


func newProducerConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewHashPartitioner
	return config
}

func NewBroker(hosts []string, topic string, partition int) *Broker {
	return &Broker{
		Hosts:     hosts,
		Topic:     topic,
		Partition: partition,
	}
}
