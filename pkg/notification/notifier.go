package spotiyou

import "github.com/schicho/spotiyou/pkg/spotify"

type Notification struct {
	Playlists []spotify.BasicPlaylist
}

type Notifier interface {
	// Notify sends the notification to the user.
	Notify(notification Notification) error
}
