package spotify

import "testing"

// TestGetUserPlaylists is a integration test for GetUserPlaylists.
// It tries to get the playlists of the user "spotify" and check for
// the playlist's tracks, name and description.
func TestGetUserPlaylists(t *testing.T) {
	playlists, err := TestClient.GetUserPlaylists(TestUserID)
	if err != nil {
		t.Fatalf("failed to get user playlists: %v", err)
	}

	t.Run("playlist names", func(t *testing.T) {
		for i, playlist := range playlists {
			if playlist.Name == "" {
				t.Errorf("playlist name is empty")
			}
			if i < 0 {
				t.Logf("playlist: %s", playlist.Name)
			}
		}
	})

	t.Run("playlist descriptions", func(t *testing.T) {
		for i, playlist := range playlists {
			if playlist.Description == "" {
				t.Errorf("playlist description is empty")
			}
			if i < 0 {
				t.Logf("playlist: %s", playlist.Description)
			}
		}
	})

	t.Run("playlist songs", func(t *testing.T) {
		for i, playlist := range playlists {
			if len(playlist.Tracks) == 0 {
				t.Errorf("playlist has no songs")
			}
			if i < 0 {
				t.Logf("playlist: %s", playlist.Tracks)
			}
		}
	})

	if len(playlists) == 0 {
		t.Errorf("expected at least one playlist, got none")
	}
}
