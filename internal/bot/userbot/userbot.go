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

	"go.uber.org/ratelimit"
	tele "gopkg.in/telebot.v3"
)

var ctx = context.Background()
var limit = config.RATE_LIMIT_GLOBAL
var sl = sendlimiter.Init(ctx, limit, limit)
var db = pg.Init(config.POSTGRES_CONN_STRING)

var buttonsRL = ratelimit.New(1, ratelimit.Per(2*time.Second), ratelimit.WithoutSlack)

type WrappedTelebot struct {
	db      *pg.PG
	bot     *tele.Bot
	buttons *models.Buttons
}

func (wt *WrappedTelebot) Start() {
	go func() {
		for {
			buttonsRL.Take()
			wt.buttons.UpdateButtons(wt.db, wt.bot)
		}
	}()

	wt.bot.Start()
	defer wt.bot.Stop()
}

func Init() *WrappedTelebot {

	token := config.USER_BOT_TOKEN

	log.Println("bot token:", token)

	bot, err := tele.NewBot(tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	bot.Use(helpers.RateLimit(sl))
	bot.Use(helpers.BotMiniLogger())
	// bot.Use(CheckAuthorize())

	bot.Handle("/id", idHandler)

	buttons, err := models.InitButtons(&db, bot)
	if err != nil {
		panic(err)
	}

	wBot := &WrappedTelebot{
		db:      &db,
		bot:     bot,
		buttons: buttons,
	}

	return wBot
}

func idHandler(c tele.Context) error {
	return c.Send(fmt.Sprintf("%d", c.Chat().ID))
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
