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

func floodMessageFromDBBroadcastMessage(dbMessage *dbmodels.BroadcastMessage) (*fm.FloodMessage, error) {

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
	case dbMessage.Image != nil && dbMessage.Message != nil:
		photoURL := config.IMGPROXY_PUBLIC_URL + *dbMessage.Image
		err := md.Convert([]byte(*dbMessage.Message), &msgBuf)
		if err != nil {
			return nil, err
		}

		message.Photo = &tele.Photo{
			File:    tele.File{FileURL: photoURL},
			Caption: msgBuf.String(),
		}

		message.Type = fm.Photo

		return message, nil
	case dbMessage.Image != nil:
		photoURL := config.IMGPROXY_PUBLIC_URL + *dbMessage.Image

		message.Photo = &tele.Photo{
			File: tele.File{FileURL: photoURL},
		}

		message.Type = fm.Photo

		return message, nil
	case dbMessage.Message != nil:
		err := md.Convert([]byte(*dbMessage.Message), &msgBuf)
		if err != nil {
			return nil, err
		}

		text := msgBuf.String()
		message.Text = &text
		message.Type = fm.Text
		return message, nil
	}

	return nil, errors.New("invalid button data")
}

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
		photoURL := config.IMGPROXY_PUBLIC_URL + "/" + *dbButton.Image
		err := md.Convert([]byte(*dbButton.Message), &msgBuf) // just use it as usual
		if err != nil {
			return nil, err
		}

		message.Photo = &tele.Photo{
			File:    tele.File{FileURL: photoURL},
			Caption: msgBuf.String(),
		}

		message.Type = fm.Photo

		return message, nil
	case dbButton.Image != nil:
		photoURL := config.IMGPROXY_PUBLIC_URL + "/" + *dbButton.Image

		message.Photo = &tele.Photo{
			File: tele.File{FileURL: photoURL},
		}

		message.Type = fm.Photo

		return message, nil
	case dbButton.Message != nil:
		err := md.Convert([]byte(*dbButton.Message), &msgBuf) // just use it as usual
		if err != nil {
			return nil, err
		}

		text := msgBuf.String()
		message.Text = &text
		message.Type = fm.Text
		return message, nil
	}

	return nil, errors.New("invalid button data")
}
