package kafka

import (
	"github.com/IBM/sarama"
	"github.com/TOMMy-Net/kafka-mess/internal/models"
)


func (b *Broker) SendMessageSarama(message models.Message) error {
	msg := &sarama.ProducerMessage{
		Topic: b.Topic,
		Key:   sarama.StringEncoder(message.UID),
		Value: sarama.StringEncoder(message.Text),
	}

	_, _, err := b.Producer.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}

