package spotify

import (
	"context"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

type SpotifyClient struct {
	client *spotify.Client
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

	return &SpotifyClient{client: client}, nil
}
