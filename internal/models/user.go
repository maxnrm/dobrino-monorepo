package models

import (
	"dobrino/internal/pg/dbmodels"

	"github.com/google/uuid"
)

type User struct {
	Id     uuid.UUID
	ChatId string
}

func (u User) FromDB(dbUser *dbmodels.User) (*User, error) {
	id, err := uuid.Parse(dbUser.ID)
	if err != nil {
		return nil, err
	}

	return &User{
		Id:     id,
		ChatId: dbUser.ChatID,
	}, nil
}

func (u *User) Recipient() string {
	return u.ChatId
}
