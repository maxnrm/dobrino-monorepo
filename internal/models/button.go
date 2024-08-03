package models

import (
	"dobrino/config"
	"dobrino/internal/pg"
	"dobrino/internal/pg/dbmodels"
	"encoding/hex"
	"errors"
	"fmt"
	"hash/fnv"
	"sync"

	fm "github.com/maxnrm/teleflood/pkg/message"
	tele "gopkg.in/telebot.v3"
)

var hashStartStr = "hashstart"

func InitButtons(db *pg.PG, bot *tele.Bot) (*Buttons, error) {
	dbButtons, err := db.GetButtons()
	if err != nil {
		return nil, err
	}

	bs := &Buttons{}

	hash := fnv.New32()
	hash.Write([]byte(hashStartStr))
	for _, b := range dbButtons {
		msg := ""
		img := ""
		if b.Message != nil {
			msg = *b.Message
		}
		if b.Image != nil {
			img = *b.Image
		}

		hash.Write([]byte(b.Name + msg + img + fmt.Sprint(b.Sort)))
	}

	bs.hash = hex.EncodeToString(hash.Sum(nil))

	bs.updateButtons(dbButtons, bot)

	return bs, nil
}

type Button struct {
	B            tele.ReplyButton
	FloodMessage *fm.FloodMessage
}

func (b *Button) HandlerFunc() tele.HandlerFunc {
	return func(c tele.Context) error {

		msg := b.FloodMessage

		msg.Recipient = fm.Recipient{ChatId: fmt.Sprint(c.Chat().ID)}

		return c.Send(msg, msg.SendOptions)
	}
}

func (b Button) FromDB(dbButton *dbmodels.Button, replyKeyboard [][]tele.ReplyButton) (Button, error) {
	btnMessage, err := FloodMessageFromDBButton(dbButton)
	if err != nil {
		fmt.Println("can't extract data from button", dbButton.Name)
		return Button{}, err
	}

	btnMessage.SendOptions = &tele.SendOptions{
		ReplyMarkup: &tele.ReplyMarkup{
			ReplyKeyboard: replyKeyboard,
		},
	}

	button := Button{
		B: tele.ReplyButton{
			Text: dbButton.Name,
		},
		FloodMessage: btnMessage,
	}

	return button, nil
}

func FloodMessageFromDBButton(dbButton *dbmodels.Button) (*fm.FloodMessage, error) {
	message := &fm.FloodMessage{}

	switch {
	case dbButton.Image != nil && dbButton.Message != nil:
		photoURL := config.IMGPROXY_PUBLIC_URL + *dbButton.Image

		message.Photo = &tele.Photo{
			File:    tele.File{FileURL: photoURL},
			Caption: *dbButton.Message,
		}

		message.Type = fm.Photo

		return message, nil
	case dbButton.Image != nil:
		photoURL := config.IMGPROXY_PUBLIC_URL + *dbButton.Image

		message.Photo = &tele.Photo{
			File: tele.File{FileURL: photoURL},
		}

		message.Type = fm.Photo

		return message, nil
	case dbButton.Message != nil:
		message.Text = dbButton.Message
		message.Type = fm.Text
		return message, nil
	}

	return nil, errors.New("invalid button data")
}

type Buttons struct {
	sync.RWMutex
	hash          string
	replyKeyboard [][]tele.ReplyButton
	buttons       []Button
}

func (bs *Buttons) updateButtons(dbButtons []*dbmodels.Button, bot *tele.Bot) {
	newButtons := []Button{}
	replyKeyboard := [][]tele.ReplyButton{}

	for _, b := range dbButtons {
		if b.Name != "/start" {
			replyKeyboard = append(replyKeyboard, []tele.ReplyButton{{Text: b.Name}})
		}
	}

	for _, b := range dbButtons {
		newButton, err := Button{}.FromDB(b, replyKeyboard)
		if err != nil {
			continue
		}

		bot.Handle(newButton.B.Text, newButton.HandlerFunc())
		newButtons = append(newButtons, newButton)
	}

	bs.buttons = newButtons
	bs.replyKeyboard = replyKeyboard
}

func (bs *Buttons) UpdateButtons(db *pg.PG, bot *tele.Bot) error {
	bs.Lock()
	defer bs.Unlock()

	// db buttons could be:
	// 1. nil, which should produce error and update should be aborted
	// 2. len(dbButtons) == 0, which should remove current buttons
	// 3. len(dbButtons) > 0, which should remove old buttons and set new buttons

	dbButtons, err := db.GetButtons()
	if err != nil {
		return err
	}

	// check if buttons are the same and in the same order
	// do not update if so
	hash := fnv.New32()
	hash.Write([]byte(hashStartStr))
	for _, b := range dbButtons {
		msg := ""
		img := ""
		if b.Message != nil {
			msg = *b.Message
		}
		if b.Image != nil {
			img = *b.Image
		}

		hash.Write([]byte(b.Name + msg + img + fmt.Sprint(b.Sort)))
	}

	hashStr := hex.EncodeToString(hash.Sum(nil))

	if hashStr == bs.hash {
		fmt.Println("hash is the same, returning...")
		return nil
	}

	fmt.Println("hash is NOT the same, operating...")

	switch {
	case len(dbButtons) == 0 && bs.buttons != nil:
		bs.buttons = nil

		for _, b := range bs.buttons {
			// remove handlers for removed buttons
			bot.Handle(b.B.Text, nil)
		}
	case len(dbButtons) > 0 && bs.buttons != nil:
		for _, b := range bs.buttons {
			// remove handlers for removed buttons
			bot.Handle(b.B.Text, nil)
		}

		bs.updateButtons(dbButtons, bot)
	case len(dbButtons) > 0 && bs.buttons == nil:
		bs.updateButtons(dbButtons, bot)
	}

	bs.hash = hashStr

	return nil
}

func (bs *Buttons) ReplyKeyboard() [][]tele.ReplyButton {
	bs.RLock()
	defer bs.RUnlock()

	return bs.replyKeyboard
}
