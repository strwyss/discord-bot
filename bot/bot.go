package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/strwyss/discord-bot/config"
)

func Start(cfg *config.AppConfig) {

	dgo, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dgo.AddHandler(messageHandler)

	dgo.Identify.Intents = discordgo.IntentsGuildMessages

	if err = dgo.Open(); err != nil {
		fmt.Println("error opening connection", err)
		return
	}
	defer dgo.Close()

	fmt.Println("Running...")

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
