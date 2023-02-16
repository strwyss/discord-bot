package main

import (
	"log"

	"github.com/strwysthvn/discord-bot/bot"
	"github.com/strwysthvn/discord-bot/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	bot.Start(cfg)
}
