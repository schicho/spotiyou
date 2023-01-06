package playlist

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
}
