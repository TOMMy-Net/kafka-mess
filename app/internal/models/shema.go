package models

type Message struct {
	UID    string `json:"uid" db:"uid"`
	Text   string `json:"text" db:"message" validate:"required,min=1"`
	Status int    `json:"status" db:"status"`
}

type Text struct {
	Text string `json:"text" validate:"required,min=1"`
}

type MessageStats struct {
	TotalMessages  int `json:"total_messages" db:"total_messages"`
	UnSendMessages int `json:"delayed_messages" db:"delayed_messages"`
	SendMessages   int `json:"total_send_messages" db:"total_sent_messages"`
}
