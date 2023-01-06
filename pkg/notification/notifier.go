package spotiyou

import "github.com/schicho/spotiyou/pkg/playlist"

type Notification struct {
	Playlists []playlist.BasicPlaylist
}

type Notifier interface {
	// Notify sends the notification to the user.
	Notify(notification Notification) error
}
