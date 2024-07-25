package models

type Message struct {
	Text string `json:"text" db:"message" validate:"required,min=1"`
}
