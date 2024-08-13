package kafka

import (
	"context"

	"github.com/IBM/sarama"
)

type Consumer struct {
	ch chan *sarama.ConsumerMessage
}

func (c *Consumer) ConsumeClaim(
	sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	msg:= <-claim.Messages()
	c.ch <- msg
	sess.MarkMessage(msg, "")
	
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (b *Broker) Read(ctx context.Context) (*sarama.ConsumerMessage, error) {
	var con = NewConsumer()
	err := b.Consumer.Consume(ctx, []string{b.Topic}, con)
	if err != nil {
		return nil, err
	}
	return <-con.ch, nil
}

func NewConsumer() *Consumer {
	return &Consumer{
		ch: make(chan *sarama.ConsumerMessage),
	}
}