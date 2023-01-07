package store

import (
	"testing"
)

func TestNewPlaylistStore(t *testing.T) {
	ps := NewPlaylistStore()
	if ps == nil {
		t.Error("NewPlaylistStore() returned nil")
	}
}

func TestPlaylistStore(t *testing.T) {
	ps := NewPlaylistStore()
	tests := []struct {
		name     string
		userID   string
		playlist string
		exists   bool
	}{
		{
			name:     "new playlist",
			userID:   "user1",
			playlist: "playlist1",
			exists:   false,
		},
		{
			name:     "existing playlist",
			userID:   "user1",
			playlist: "playlist1",
			exists:   true,
		},
		{
			name:     "new playlist for different user",
			userID:   "user2",
			playlist: "playlist1",
			exists:   false,
		},
		{
			name:     "existing playlist for different user",
			userID:   "user2",
			playlist: "playlist1",
			exists:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ps.HasPlaylist(tt.userID, tt.playlist) != tt.exists {
				t.Errorf("expected playlist %s for user %s to exist: %t", tt.playlist, tt.userID, tt.exists)
			}

			if !tt.exists {
				ps.AddPlaylist(tt.userID, tt.playlist)
			}

			if !ps.HasPlaylist(tt.userID, tt.playlist) {
				t.Errorf("expected playlist %s for user %s to exist", tt.playlist, tt.userID)
			}
		})
	}
}
