package main

import (
	"log"

	"github.com/cecepsprd/discord-bot/bot"
	"github.com/cecepsprd/discord-bot/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	bot.Start(cfg)
}
