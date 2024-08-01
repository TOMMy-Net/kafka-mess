package services

import (
	"context"
	"log"
	"time"

	"github.com/TOMMy-Net/kafka-mess/internal/kafka"
	"github.com/TOMMy-Net/kafka-mess/internal/models"
)

type BrokerSend interface {
	SendMessage(message models.Message) error
}


type BaseUpdater interface {
	UpdateMessageStatus(ctx context.Context, uid string, status int) error
	UnSendMessages(ctx context.Context) ([]models.Message, error)
}



func MessageKeeper(ctx context.Context, b *kafka.Broker, db BaseUpdater) {
	for {
		<-time.After(10 * time.Second)
		messages, err := db.UnSendMessages(ctx)
		if err != nil {
			log.Println(err)
			continue
		}

		for i := 0; i < len(messages); i++ {
			if err := b.SendMessage(messages[i]); err != nil {
				continue
			}
			db.UpdateMessageStatus(ctx, messages[i].UID, 1)
		}
	}
}


