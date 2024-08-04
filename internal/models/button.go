package models

import (
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

	bs.updateButtons(dbButtons)

	return bs, nil
}

type Button struct {
	B            tele.ReplyButton
	FloodMessage *fm.FloodMessage
}

func (b Button) FromDB(dbButton *dbmodels.Button) (Button, error) {
	btnMessage, err := floodMessageFromDBButton(dbButton)
	if err != nil {
		fmt.Println("can't extract data from button", dbButton.Name)
		return Button{}, err
	}

	button := Button{
		B: tele.ReplyButton{
			Text: dbButton.Name,
		},
		FloodMessage: btnMessage,
	}

	return button, nil
}

type Buttons struct {
	sync.RWMutex
	hash          string
	replyKeyboard [][]tele.ReplyButton
	buttons       map[string]Button
}

func (bs *Buttons) updateButtons(dbButtons []*dbmodels.Button) {
	newButtons := make(map[string]Button)
	replyKeyboard := [][]tele.ReplyButton{}

	for _, b := range dbButtons {
		if b.Name != "/start" {
			replyKeyboard = append(replyKeyboard, []tele.ReplyButton{{Text: b.Name}})
		}
	}

	for _, b := range dbButtons {
		newButton, err := Button{}.FromDB(b)
		if err != nil {
			continue
		}

		newButtons[newButton.B.Text] = newButton
	}

	bs.buttons = newButtons
	bs.replyKeyboard = replyKeyboard
}

func (bs *Buttons) UpdateButtons(db *pg.PG, bot *tele.Bot) error {
	bs.Lock()
	defer bs.Unlock()

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
	case len(dbButtons) > 0:
		bs.updateButtons(dbButtons)
	}

	bs.hash = hashStr

	return nil
}

func (bs *Buttons) ReplyKeyboard() [][]tele.ReplyButton {
	return bs.replyKeyboard
}

func (bs *Buttons) Button(name string) (*Button, error) {
	button, ok := bs.buttons[name]
	if !ok {
		return nil, errors.New("button with that name does not exist")
	}

	return &button, nil
}
