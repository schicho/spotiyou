package api

import (
	"log"
	"sync/atomic"
	"time"

	spotiyou "github.com/schicho/spotiyou/internal/spotify"
	"github.com/schicho/spotiyou/pkg/playlist"
)

// API is the access point to the Spotify API.
// It wraps the internal/spotify package.
type API struct {
	sc atomic.Pointer[spotiyou.SpotifyClient]
}

func New(clientID, clientSecret string) (*API, error) {
	spotifyClient, err := spotiyou.NewSpotifyClient(clientID, clientSecret)
	if err != nil {
		return nil, err
	}

	var api API = API{}
	api.sc.Store(spotifyClient)

	// renew token every 45 minutes
	// this is a limitation of the auth mechanism of
	// the library.
	go func() {
		t := time.NewTicker(45 * time.Minute)
		for {
			<-t.C
			spotifyClient, err := spotiyou.NewSpotifyClient(clientID, clientSecret)
			if err != nil {
				log.Println("error renewing token:", err)
				continue
			}
			api.sc.Store(spotifyClient)
		}
	}()

	return &api, nil
}

func (a *API) GetUserPlaylists(userID string) ([]playlist.BasicPlaylist, error) {
	spotifyClient := a.sc.Load()
	return spotifyClient.GetUserPlaylists(userID)
}
