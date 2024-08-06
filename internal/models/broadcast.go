package models

import (
	"dobrino/internal/pg"
	"errors"
	"fmt"

	tele "gopkg.in/telebot.v3"
)

type Broadcast struct {
	// broadcastMessages []BroadcastMessage
}

func (b *Broadcast) Broadcast(bot *tele.Bot, db *pg.PG) error {
	dbMsg, err := db.GetBroadcastMessageForSend()
	if err != nil {
		fmt.Println("broadcast: failed to get message to broadcast. error:", err)
		return err
	}

	msg, err := floodMessageFromDBBroadcastMessage(dbMsg)
	if err != nil {
		db.SetBroadcastMessageStatus(dbMsg.ID, false)
		fmt.Println("broadcast: failed to parse message. error:", err)
		return err
	}

	dbUsers, err := db.GetUsers()
	if err != nil {
		db.SetBroadcastMessageStatus(dbMsg.ID, false)
		fmt.Println("broadcast: failed to get users. error:", err)
		return err
	}

	if len(dbUsers) == 0 {
		db.SetBroadcastMessageStatus(dbMsg.ID, false)
		err := errors.New("no users to broadcase message to")
		fmt.Println("broadcast: erorr:", err)
		return err
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
