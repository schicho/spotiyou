package spotify

import (
	"context"
	"errors"
	"log"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

var ErrRetrievalFailed = errors.New("spotify: retrieval failed")

type SpotifyClient struct {
	client *spotify.Client
	logger *log.Logger
}

func NewSpotifyClient(ClientID, ClientSecret string) (*SpotifyClient, error) {
	config := clientcredentials.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		TokenURL:     spotifyauth.TokenURL,
	}

	token, err := config.Token(context.Background())
	if err != nil {
		return nil, err
	}
	httpClient := spotifyauth.New().Client(context.Background(), token)
	client := spotify.New(httpClient)

	return &SpotifyClient{client: client, logger: newSpotifyLogger()}, nil
}

func newSpotifyLogger() *log.Logger {
	return log.New(log.Writer(), "spotify", log.Flags())
}
