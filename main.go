package main

import (
	"log"

	"github.com/strwyss/discord-bot/bot"
	"github.com/strwyss/discord-bot/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	bot.Start(cfg)
}
