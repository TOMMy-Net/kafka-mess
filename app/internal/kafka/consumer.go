package kafka

import (
	"context"
	"io"

	"github.com/segmentio/kafka-go"
)

func (b *Broker) ReadWithConfig(ch chan string) error {

	config := kafka.ReaderConfig{
		Brokers: b.Hosts,
		Topic:   b.Topic,
		GroupID: "my-group-1",
	}

	// Создание потребителя
	reader := kafka.NewReader(config)
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			if err == io.EOF {
				break
			}
			ch <- err.Error()
		}

		ch <- string(msg.Value)

	}
	return nil
}

func (b *Broker) Read() {

}
