package store

// PlaylistStore is a map of playlists for each user.
// It maps a user ID to a map of playlist names.
//
// The map can be used to check if a playlist exists for a user
// or if the user has created a new playlist.
type PlaylistStore map[string]map[string]struct{}

// NewPlaylistStore creates a new PlaylistStore.
func NewPlaylistStore() PlaylistStore {
	return make(PlaylistStore)
}

// AddPlaylist adds a playlist to the store.
func (ps PlaylistStore) AddPlaylist(userID, playlistName string) {
	if _, ok := ps[userID]; !ok {
		ps[userID] = make(map[string]struct{})
	}
	ps[userID][playlistName] = struct{}{}
}

// HasPlaylist checks if a playlist exists for a user.
func (ps PlaylistStore) HasPlaylist(userID, playlistName string) bool {
	_, ok := ps[userID][playlistName]
	return ok
}
