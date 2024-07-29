package userbot

import (
	"context"
	"dobrino/config"
	"dobrino/internal/helpers"
	"dobrino/internal/pg"
	"dobrino/internal/sendlimiter"
	"fmt"
	"log"
	"time"

	"go.uber.org/ratelimit"
	tele "gopkg.in/telebot.v3"
)

type ButtonDataMap struct {
	Button tele.ReplyButton
	Data
}

var ctx = context.Background()
var db = pg.Init(config.POSTGRES_CONN_STRING)
var limit = config.RATE_LIMIT_GLOBAL
var sl = sendlimiter.Init(ctx, limit, limit)
var gButtons *[][]tele.ReplyButton
var buttonsRL = ratelimit.New(1, ratelimit.Per(2*time.Second), ratelimit.WithoutSlack)
var defaultOpts = tele.SendOptions{
	ParseMode: tele.ModeHTML,
}

func Init() *tele.Bot {

	token := config.USER_BOT_TOKEN

	log.Println("bot token:", token)

	bot, err := tele.NewBot(tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		buttonsRL.Take()
		gButtons = updateButtons()
	}()

	bot.Use(helpers.RateLimit(sl))
	bot.Use(helpers.BotMiniLogger())
	// bot.Use(CheckAuthorize())

	bot.Handle("/id", idHandler)
	bot.Handle("/start", startHandler)

	return bot
}

func startHandler(c tele.Context) error {
	c.Send("Start message")
	return nil
}

func idHandler(c tele.Context) error {

	fmt.Println(gButtons)

	opts := defaultOpts
	if gButtons != nil {
		opts.ReplyMarkup = &tele.ReplyMarkup{
			ReplyKeyboard: *gButtons,
		}
	}

	fmt.Println(opts)

	return c.Send(fmt.Sprintf("%d", c.Chat().ID), &opts)
}

func updateButtons() (*[][]tele.ReplyButton, error) {

	dbButtons, err := db.GetButtons()
	if err != nil {
		return nil, err
	}

	buttons := [][]tele.ReplyButton{}

	for _, b = range dbButtons {
		button := tele.ReplyButton{}
	}

	return &buttons, nil
}

func CheckAuthorize() tele.MiddlewareFunc {
	l := log.Default()
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if c.Message().Text == "Начать" {
				return next(c)
			}

			chatID := fmt.Sprint(c.Chat().ID)
			_, err := db.GetUser(chatID)
			if err != nil {
				c.Send("Not authorized")
				return nil
			}

			l.Println("Юзер", chatID, "авторизован")
			return next(c)
		}
	}
}
