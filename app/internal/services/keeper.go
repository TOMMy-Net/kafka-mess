package services

import (
	"context"
	"log"
	"time"

	"github.com/TOMMy-Net/kafka-mess/internal/models"
)

type BrokerSend interface {
	SendMessage(message models.Message) error
}

type BaseUpdater interface {
	UpdateMeesageStatus(ctx context.Context, uid string, status int) error
	UnSendMessages(ctx context.Context) ([]models.Message, error)
}

func MessageKeeper(ctx context.Context, b BrokerSend, db BaseUpdater) {
	for {
		<-time.After(10 * time.Second)
		messages, err := db.UnSendMessages(ctx)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, v := range messages {
			if err := b.SendMessage(v); err != nil {
				continue
			}
			db.UpdateMeesageStatus(ctx, v.UID, 1)
		}

	}
}
