package models

import (
	"dobrino/internal/pg"
	"errors"

	tele "gopkg.in/telebot.v3"
)

type Broadcast struct {
	// broadcastMessages []BroadcastMessage
}

func (b *Broadcast) Broadcast(bot *tele.Bot, db *pg.PG) error {
	dbMsg, err := db.GetBroadcastMessageForSend()
	if err != nil {
		return errors.New("no message for broadcast provided")
	}

	msg, err := floodMessageFromDBBroadcastMessage(dbMsg)
	if err != nil {
		db.SetBroadcastMessageStatus(dbMsg.ID, false)
		return err
	}

	dbUsers, err := db.GetUsers()
	if err != nil {
		db.SetBroadcastMessageStatus(dbMsg.ID, false)
		return err
	}

	if len(dbUsers) == 0 {
		return errors.New("no users to broadcast message to")
	}

	for _, dbUser := range dbUsers {

		u, err := User{}.FromDB(dbUser)
		if err != nil {
			continue
		}

		bot.Send(u, msg)
	}

	db.SetBroadcastMessageStatus(dbMsg.ID, true)
	return nil
}

// type BroadcastMessage struct {
// 	message *fm.FloodMessage
// }
