package kafka

import (
	"context"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

var Broker *kafka.Conn



func ConnectBroker() error {

	kafkaHost := fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))
	// Настройка подключения к брокеру
	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaHost, os.Getenv("KAFKA_TOPIC"), 0)
	if err != nil {
		return err
	}
	
	Broker = conn


	return nil
}
