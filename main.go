package main

import (
	// "encoding/json"
	"flag"
	"log"

	// "fmt"
	// "io/ioutil"
	// "net/http"
	// "os"
	// "os/signal"
	// "strings"
	// "syscall"
	// "github.com/bwmarrin/discordgo"
	"github.com/saurabh3460/weather-discord-bot/bot"
)

// const KuteGoAPIURL = "https://kutego-api-xxxxx-ew.a.run.app"

func init() {
	flag.StringVar(&bot.DiscordToken, "dt", "", "Discord Token")
	flag.StringVar(&bot.WeatherToken, "wt", "", "OpenWeather Token")
	flag.Parse()
}

func main() {
	if bot.DiscordToken == "" {
		log.Fatal("No Discord Token provided")
	}
	if bot.WeatherToken == "" {
		log.Fatal("No OpenWeather Token provided")
	}
	// fmt.Printf("%s \n", bot.DiscordToken)
	// fmt.Printf("%s", bot.WeatherToken)

	bot.Run()
}
