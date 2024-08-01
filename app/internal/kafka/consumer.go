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
	close(ch)
	return nil
}


func (b *Broker) Read(ctx context.Context) (kafka.Message, error) {
	var ch = make(chan kafka.Message)

	go func() {
		mess, err := b.Conn.ReadMessage(1024 * 10)
		if err != nil {
			ch <- kafka.Message{}
			return
		}
		ch <- mess
	}()

	select {
		case <-ctx.Done():
			return kafka.Message{}, ctx.Err()
		case res := <- ch:
			return res, nil
	}
}
