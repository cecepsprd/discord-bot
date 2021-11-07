package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/cecepsprd/discord-bot/config"
)

func Start(cfg *config.AppConfig) {

	dg, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageHandler)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	if err = dg.Open(); err != nil {
		fmt.Println("error opening connection", err)
		return
	}
	defer dg.Close()

	fmt.Println("Bot is running!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

// This function will be called whenever bot received a message
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}

}
