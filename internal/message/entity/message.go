package entity

import "github.com/google/uuid"

type Message struct {
	Id        uuid.UUID
	ReciverId int64
	Body      string
}

func NewMessage(reciverId int64, body string) *Message {
	return &Message{
		Id:        uuid.New(),
		ReciverId: reciverId,
		Body:      body,
	}
}
