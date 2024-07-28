package kafka

import (
	"github.com/TOMMy-Net/kafka-mess/internal/models"
	"github.com/segmentio/kafka-go"
)

func SendMessage(message models.Message) error {
	msg := kafka.Message{
		Value: []byte(message.Text),
	}

	// Отправка сообщения
	_, err := Broker.WriteMessages(msg)
	if err != nil {
		return err
	}
	return nil
}
