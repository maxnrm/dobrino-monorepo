package models

import (
	"dobrino/internal/pg"
	"errors"
	"fmt"

	s "github.com/maxnrm/teleflood/pkg/sender"
)

type Broadcast struct {
	db     *pg.PG
	sender *s.Sender
}

func NewBroadcast(db *pg.PG, sender *s.Sender) *Broadcast {
	return &Broadcast{
		db:     db,
		sender: sender,
	}
}

func (b *Broadcast) Broadcast() error {
	dbMsg, err := b.db.GetBroadcastMessageForSend()
	if err != nil {
		// fmt.Println("broadcast: failed to get message to broadcast. error:", err)
		return err
	}

	msg, err := floodMessageFromDBBroadcastMessage(dbMsg)
	if err != nil {
		b.db.SetBroadcastMessageStatus(dbMsg.ID, false)
		fmt.Println("broadcast: failed to parse message. error:", err)
		return err
	}

	dbUsers, err := b.db.GetUsers()
	if err != nil {
		b.db.SetBroadcastMessageStatus(dbMsg.ID, false)
		fmt.Println("broadcast: failed to get users. error:", err)
		return err
	}

	if len(dbUsers) == 0 {
		b.db.SetBroadcastMessageStatus(dbMsg.ID, false)
		err := errors.New("no users to broadcase message to")
		fmt.Println("broadcast: erorr:", err)
		return err
	}

	for _, dbUser := range dbUsers {
		u, err := User{}.FromDB(dbUser)
		if err != nil {
			continue
		}
		b.sender.Send(u, msg, msg.SendOptions)
	}

	usersNumber := len(dbUsers)

	b.db.SetBroadcastMessageStatus(dbMsg.ID, true)
	fmt.Println("broadcast: message sent to ", usersNumber)
	return nil
}

// type BroadcastMessage struct {
// 	message *fm.FloodMessage
// }
