package bot

import "github.com/bwmarrin/discordgo"

const URL string = "https://api.openweathermap.org/data/2.5/weather?"

type WeatherData struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`

	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`

	Wind struct {
		Speed float32 `json:"speed"`
	} `json:"wind"`
}

func getCurrentWeather(message string) *discordgo.MessageSend {

}
