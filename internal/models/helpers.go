package models

import (
	"bytes"
	"dobrino/config"
	"dobrino/internal/pg/dbmodels"
	"errors"

	"github.com/leonid-shevtsov/telegold"
	fm "github.com/maxnrm/teleflood/pkg/message"
	"github.com/yuin/goldmark"
	tele "gopkg.in/telebot.v3"
)

func floodMessageFromDBButton(dbButton *dbmodels.Button) (*fm.FloodMessage, error) {
	// set parseMode HTML as default
	// as we are using goldmark and telegold
	message := &fm.FloodMessage{
		SendOptions: &tele.SendOptions{
			ParseMode: tele.ModeHTML,
		},
	}

	var msgBuf bytes.Buffer
	md := goldmark.New(goldmark.WithRenderer(telegold.NewRenderer()))

	switch {
	case dbButton.Image != nil && dbButton.Message != nil:
		photoURL := config.IMGPROXY_PUBLIC_URL + *dbButton.Image
		md.Convert([]byte(*dbButton.Message), &msgBuf) // just use it as usual

		message.Photo = &tele.Photo{
			File:    tele.File{FileURL: photoURL},
			Caption: msgBuf.String(),
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
		md.Convert([]byte(*dbButton.Message), &msgBuf) // just use it as usual
		text := msgBuf.String()
		message.Text = &text
		message.Type = fm.Text
		return message, nil
	}

	return nil, errors.New("invalid button data")
}
