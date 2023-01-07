package spotiyou

import (
	"github.com/schicho/spotiyou/internal/store"
	"github.com/schicho/spotiyou/pkg/api"
	notification "github.com/schicho/spotiyou/pkg/notification"
)

type Spotiyou struct {
	store store.PlaylistStore
	api   *api.API
	notification.Notifier
}

// NewSpotiyou creates a new Spotiyou instance.
//
// An instance of the API is required to access the Spotify API.
// A notifier is required to send notifications.
func NewSpotiyou(api *api.API, notifier notification.Notifier) *Spotiyou {
	return &Spotiyou{
		store:    store.NewPlaylistStore(),
		api:      api,
		Notifier: notifier,
	}
}

// AddUser adds a user to Spotiyou.
//
// It runs a initial sync of the user's playlists.
// No notifications are sent during this sync.
func (s *Spotiyou) AddUser(userID string) error {
	pl, err := s.api.GetUserPlaylists(userID)
	if err != nil {
		return err
	}

	for _, p := range pl {
		s.store.AddPlaylist(userID, p.Name)
	}
	return nil
}

// SpotUser checks if a user has created new playlists.
//
// If new playlists are found, a notification is sent.
// It takes the userID as the argument.
func (s *Spotiyou) SpotUser(userID string) error {
	pl, err := s.api.GetUserPlaylists(userID)
	if err != nil {
		return err
	}

	var notification notification.Notification

	for _, p := range pl {
		if !s.store.HasPlaylist(userID, p.Name) {
			s.store.AddPlaylist(userID, p.Name)
			notification.Playlists = append(notification.Playlists, p)
		}
	}

	if len(notification.Playlists) > 0 {
		return s.Notify(notification)
	}
	return nil
}
