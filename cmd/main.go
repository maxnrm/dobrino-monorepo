package main

import (
	"dobrino/internal/bot/userbot"
)

var userBot = userbot.Init()

// var adminBot = adminbot.Init()

func main() {
	userBot.Start()
}
