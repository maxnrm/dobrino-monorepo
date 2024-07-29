package main

import (
	"dobrino/internal/bot/userbot"
	"sync"
)

var wg sync.WaitGroup
var userBot = userbot.Init()

// var adminBot = adminbot.Init()

func main() {
	wg.Add(2)

	// go ws.Start()

	go userBot.Start()
	defer userBot.Stop()

	// go adminBot.Start()
	// defer adminBot.Stop()

	wg.Wait()
}
