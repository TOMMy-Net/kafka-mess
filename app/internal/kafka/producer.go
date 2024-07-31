package kafka

import (
	"context"

	"github.com/TOMMy-Net/kafka-mess/internal/models"
	"github.com/segmentio/kafka-go"
)

func (b *Broker) SendMessage(message models.Message) error {
	msg := kafka.Message{
		Value: []byte(message.Text),
		Key:   []byte(message.UID),
	}

	// Отправка сообщения
	_, err := b.WriteMessages(msg)
	if err != nil {
		b.Reconect(b.Topic, b.Partition)
		return err
	}

	return nil
}

func (b *Broker) WriteWithConfig(ctx context.Context, m models.Message) error {
	// Создание производителя
	writer := &kafka.Writer{
		Addr:  kafka.TCP(b.Hosts...),
		Topic: b.Topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	// Отправка сообщения
	message := kafka.Message{
		Value: []byte(m.Text),
		Key:   []byte(m.UID),
	}

	err := writer.WriteMessages(ctx, message)
	if err != nil {
		return err
	}
	return nil
}
