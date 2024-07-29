package models

import "github.com/google/uuid"

type User struct {
	ID     uuid.UUID `json:"id"`
	ChatID string    `json:"chat_id"`
}
