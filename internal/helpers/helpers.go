package helpers

import (
	"dobrino/internal/sendlimiter"
	"fmt"
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func BotMiniLogger() tele.MiddlewareFunc {
	l := log.Default()

	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			chatID := c.Chat().ID
			text := c.Message().Text
			l.Println(chatID, text, "ok")
			return next(c)
		}
	}
}

func RateLimit(sl *sendlimiter.SendLimiter) tele.MiddlewareFunc {
	l := log.Default()
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			chatID := fmt.Sprint(c.Chat().ID)
			userRateLimiter := sl.GetUserRateLimiter(chatID)
			if userRateLimiter == nil {
				sl.AddUserRateLimiter(chatID, 2, 2)
				userRateLimiter = sl.GetUserRateLimiter(chatID)
			}

			if !userRateLimiter.RateLimiter.Allow() {
				l.Println("Rate limit exceeded for", chatID, "returning...")
				return nil
			}

			return next(c)
		}
	}
}

func IsAfter(now, goal time.Time) bool {
	return now.After(goal)
}

func IsNowAfter(goal time.Time) bool {
	return time.Now().After(goal)
}
