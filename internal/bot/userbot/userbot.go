package userbot

import (
	"context"
	"dobrino/config"
	"dobrino/internal/helpers"
	"dobrino/internal/models"
	"dobrino/internal/pg"
	"dobrino/internal/sendlimiter"
	"fmt"
	"log"
	"time"

	"github.com/maxnrm/teleflood/pkg/sender"
	"go.uber.org/ratelimit"
	tele "gopkg.in/telebot.v3"
)

var ctx = context.Background()
var sl = sendlimiter.Init(ctx, limit, limit)
var db = pg.Init(config.POSTGRES_CONN_STRING)

var limit = config.RATE_LIMIT_GLOBAL
var cronRL = ratelimit.New(1, ratelimit.Per(2*time.Second), ratelimit.WithoutSlack)

var captchaButtonText = "Я не робот"
var captchaButton = tele.ReplyButton{Text: captchaButtonText}
var captchaReplyMarkup = &tele.ReplyMarkup{ResizeKeyboard: true, ReplyKeyboard: [][]tele.ReplyButton{{captchaButton}}}

type WrappedTelebot struct {
	db        *pg.PG
	bot       *tele.Bot
	buttons   *models.Buttons
	broadcast *models.Broadcast
	sender    *sender.Sender
}

func (wt *WrappedTelebot) Start() {
	go func() {
		for {
			cronRL.Take()
			wt.buttons.UpdateButtons()
			wt.broadcast.Broadcast()
		}
	}()

	wt.bot.Start()
	defer wt.bot.Stop()
}

func Init() *WrappedTelebot {

	token := config.USER_BOT_TOKEN

	log.Println("bot token:", token)

	bot, err := tele.NewBot(tele.Settings{
		Token:     token,
		ParseMode: tele.ModeHTML,
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	buttons, err := models.InitButtons(db)
	if err != nil {
		panic(err)
	}

	broadcastRL := ratelimit.New(config.RATE_LIMIT_BROADCAST, ratelimit.Per(1*time.Second), ratelimit.WithoutSlack)
	broadcastSender, _ := sender.New(bot, broadcastRL)
	broadcast := models.NewBroadcast(db, broadcastSender)

	globalRL := ratelimit.New(config.RATE_LIMIT_GLOBAL, ratelimit.Per(1*time.Second), ratelimit.WithSlack(2))
	globalSender, _ := sender.New(bot, globalRL)
	wBot := &WrappedTelebot{
		db:        db,
		bot:       bot,
		buttons:   buttons,
		broadcast: broadcast,
		sender:    globalSender,
	}

	bot.Use(helpers.RateLimit(sl))
	bot.Use(helpers.BotMiniLogger())
	bot.Use(CheckAuthorize())

	bot.Handle("/id", idHandler)
	bot.Handle(tele.OnText, wBot.OnTextHandler())

	return wBot
}

func (wt *WrappedTelebot) OnTextHandler() tele.HandlerFunc {
	return func(c tele.Context) error {
		wt.buttons.RLock()
		defer wt.buttons.RUnlock()

		text := c.Text()

		replyKeyboard := wt.buttons.ReplyKeyboard()
		markup := c.Bot().NewMarkup()
		markup.ResizeKeyboard = true

		if len(replyKeyboard) > 0 {
			markup.ReplyKeyboard = replyKeyboard
		} else {
			markup.RemoveKeyboard = true
		}

		button, err := wt.buttons.Button(text)
		if err != nil {
			c.Bot().Send(c.Recipient(), "Выберите одну из команд на клавиатуре ⬇️", markup)
			return nil
		}

		opts := &tele.SendOptions{
			ReplyMarkup: markup,
			ParseMode:   tele.ModeHTML,
		}

		return wt.sender.Send(c.Chat(), button.FloodMessage, opts)
	}
}

func idHandler(c tele.Context) error {
	return c.Send(fmt.Sprintf("%d", c.Chat().ID))
}

func CheckAuthorize() tele.MiddlewareFunc {
	l := log.Default()
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if c.Message().Text == "/start" {
				next(c)
			}

			chatId := fmt.Sprint(c.Chat().ID)
			err := db.IncrementUserInteractions(chatId)
			if err == nil {
				l.Println("Юзер", chatId, "авторизован")
				return next(c)
			}

			if c.Message().Text == captchaButtonText {
				err := db.CreateUser(chatId)
				if err == nil {
					l.Println("Юзер", chatId, "зарегистрирован")
					return next(c)
				}
				return c.Send("Подтвердите, что вы не робот", captchaReplyMarkup)
			}

			l.Println("Юзер", chatId, "НЕ авторизован")
			return c.Send("Подтвердите, что вы не робот", captchaReplyMarkup)
		}
	}
}
