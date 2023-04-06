package bot

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	DiscordToken string
	WeatherToken string
)

func Run() {
	// Create new Discord Session
	discord, err := discordgo.New("Bot " + DiscordToken)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Token", discord.Token)
	discord.AddHandler(newMessage)
	discord.Open()
	defer discord.Close()

	fmt.Println("Bot running...")
	c := make(chan os.Signal, 1)
	<-c
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore bot messaage
	if message.Author.ID == discord.State.User.ID {
		return
	}
	fmt.Printf("Message received: %v\n", message.Content)
	// Respond to messages
	switch {
	case strings.Contains(message.Content, "weather"):
		discord.ChannelMessageSend(message.ChannelID, "There we go, I can you there!")
	case strings.Contains(message.Content, "bot"):
		discord.ChannelMessageSend(message.ChannelID, "Hi how can I help you!")
	case strings.Contains(message.Content, "!zip"):
		currentWeather := getCurrentWeather(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, currentWeather)
	}
}
