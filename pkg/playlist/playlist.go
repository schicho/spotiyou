package playlist

import (
	"bytes"
)

// BasicArtist represents a simplified artist object.
type BasicArtist struct {
	Name string
}

// BasicTrack represents a simplified music track object.
type BasicTrack struct {
	Name    string
	Artists []BasicArtist
}

// BasicPlaylist represents a simplified playlist object.
type BasicPlaylist struct {
	Name        string
	Description string
	OwnerName   string
	Tracks      []BasicTrack
	// The image of the playlist is provided as a byte buffer.
	// May be empty if no image is available.
	Image *bytes.Buffer
	// URL which links to the public playlist on Spotify
	URL string
}
