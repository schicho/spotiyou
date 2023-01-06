package api

import (
	spotiyou "github.com/schicho/spotiyou/internal/spotify"
	"github.com/schicho/spotiyou/pkg/playlist"
)

// API is the access point to the Spotify API.
// It wraps the internal/spotify package.
type API struct {
	sc *spotiyou.SpotifyClient
}

func New(clientID, clientSecret string) (*API, error) {
	spotifyClient, err := spotiyou.NewSpotifyClient(clientID, clientSecret)
	if err != nil {
		return nil, err
	}

	return &API{
		sc: spotifyClient,
	}, nil
}

func (a *API) GetUserPlaylists(userID string) ([]playlist.BasicPlaylist, error) {
	return a.sc.GetUserPlaylists(userID)
}
