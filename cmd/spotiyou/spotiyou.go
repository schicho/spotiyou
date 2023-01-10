package main

import (
	"log"
	"os"
	"time"

	"github.com/schicho/spotiyou/internal/telegram"
	"github.com/schicho/spotiyou/pkg/api"
	"github.com/schicho/spotiyou/pkg/spotiyou"
)

func main() {
	var telegramToken string
	var telegramChatID string
	var spotifyClientID string
	var spotifyClientSecret string

	telegramToken = os.Getenv("TELEGRAM_TOKEN")
	telegramChatID = os.Getenv("TELEGRAM_CHAT_ID")
	spotifyClientID = os.Getenv("SPOTIFY_CLIENT_ID")
	spotifyClientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")

	if telegramToken == "" || telegramChatID == "" || spotifyClientID == "" || spotifyClientSecret == "" {
		log.Fatal("missing environment variables")
	}

	teleNotifier := telegram.NewTelegramNotifier(telegramToken, telegramChatID)
	spotifyApi, err := api.New(spotifyClientID, spotifyClientSecret)
	if err != nil {
		log.Fatal(err)
	}

	spot := spotiyou.NewSpotiyou(spotifyApi, teleNotifier)

	spot.AddUser("spotify") // add user to spotiyou by userID (can be found in the URL of the user's profile)

	ticker := time.NewTicker(3 * time.Minute)
	for range ticker.C {
		log.Println("spotting users")
		err := spot.SpotAllUsers()
		if err != nil {
			log.Println(err)
		}
	}
}
