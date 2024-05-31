package entity

import "github.com/google/uuid"

type Message struct {
	Id     uuid.UUID
	UserId int64
	Body   string
}

func NewMessage(userId int64, body string) *Message {
	return &Message{
		Id:     uuid.New(),
		UserId: userId,
		Body:   body,
	}
}
