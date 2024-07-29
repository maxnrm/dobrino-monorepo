package adminbot

import (
	"context"
	"dobrino/config"
	"fmt"
	"log"
	"time"

	"dobrino/internal/helpers"
	"dobrino/internal/pg"
	"dobrino/internal/sendlimiter"

	tele "gopkg.in/telebot.v3"
)

var ctx = context.Background()
var sl = sendlimiter.Init(ctx, config.RATE_LIMIT_GLOBAL, config.RATE_LIMIT_BURST_GLOBAL)
var db = pg.DB

func Init() *tele.Bot {

	token := config.ADMIN_BOT_TOKEN

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
	bot.Use(CheckAuthorize())

	bot.Handle("/id", idHandler)

	return bot
}

func idHandler(c tele.Context) error {
	return c.Send(fmt.Sprintf("%d", c.Chat().ID))
}

func CheckAuthorize() tele.MiddlewareFunc {
	l := log.Default()

	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if c.Message().Text == config.ADMIN_AUTH_CODE {
				return next(c)
			}

			chatId := fmt.Sprint(c.Chat().ID)
			_, err := db.GetAdmin(chatId)

			if err != nil {
				c.Send("Unauthorized")
			}

			l.Println("Админ", chatId, "авторизован")
			return next(c)
		}
	}
}
