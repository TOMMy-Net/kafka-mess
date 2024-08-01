package kafka

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

type Broker struct {
	Conn      *kafka.Conn
	Hosts     []string
	Topic     string
	Partition int
	Reconect  chan struct{}
}

var (
	ErrWithWrite = errors.New("the message was delivered but there was a problem on the intermediate node")
	ErrWithRead  = errors.New("error reading a message from kafka")
)

func ConnectBroker(topic string, partition int) (*Broker, error) {
	var hosts []string
	var ch = make(chan struct{})
	kafkaHost := fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))
	// Настройка подключения к брокеру
	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaHost, topic, partition)
	if err != nil {
		return &Broker{}, err
	}
	//conn.SetDeadline(time.Time{}.Add(time.Second * 3))
	hosts = append(hosts, kafkaHost)

	return &Broker{conn, hosts, topic, partition, ch}, nil
}

func (b *Broker) ConectKeeper(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-b.Reconect:
			b.Conn.ReadOffsets()
			conn, err := ConnectBroker(b.Topic, b.Partition)
			if err != nil {
				log.Println(err)
			}
			b.Conn = conn.Conn
		}
	}
}

func (b *Broker) ConnectKeeperV2() {
	for {
		<-time.After(2 * time.Second)
		_, _, err := b.Conn.ReadOffsets()
		if err != nil {
			log.Println(err)
			conn, err := ConnectBroker(b.Topic, b.Partition)
			if err != nil {
				log.Println(err)
				continue
			}
			(*b).Conn = conn.Conn
		}
	}
}

func (b *Broker) Connect() {

}

func NewBroker(hosts []string, topic string, partition int) *Broker {
	return &Broker{
		Hosts:     hosts,
		Topic:     topic,
		Partition: partition,
	}
}
