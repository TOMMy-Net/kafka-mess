package kafka

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

type Broker struct {
	*kafka.Conn
	Hosts []string
	Topic string
	Partition int
}

var (
	ErrWithWrite = errors.New("the message was delivered but there was a problem on the intermediate node")
)

func ConnectBroker(topic string, partition int) (*Broker, error) {
	var hosts []string

	kafkaHost := fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))
	// Настройка подключения к брокеру
	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaHost, topic, partition)
	if err != nil {
		return &Broker{}, err
	}
	hosts = append(hosts, kafkaHost)
	
	return &Broker{conn, hosts, topic, partition}, nil
}


func (b *Broker) Reconect(topic string, partition int) {
}

func NewBroker(hosts []string, topic string, partition int) *Broker {
	return &Broker{
		Hosts: hosts,
		Topic: topic,
		Partition: partition,
	}
}