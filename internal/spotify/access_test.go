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
			if i < 3 {
				t.Logf("playlist: %s", playlist.Name)
			}
		}
	})

	t.Run("playlist descriptions", func(t *testing.T) {
		for i, playlist := range playlists {
			if playlist.Description == "" {
				t.Errorf("playlist description is empty")
			}
			if i < 3 {
				t.Logf("playlist: %s", playlist.Description)
			}
		}
	})

	t.Run("playlist songs", func(t *testing.T) {
		for i, playlist := range playlists {
			if len(playlist.Tracks) == 0 {
				t.Errorf("playlist has no songs")
			}
			if i < 3 {
				t.Logf("playlist: %s", playlist.Tracks)
			}
		}
	})

	t.Run("playlist image", func(t *testing.T) {
		for i, playlist := range playlists {
			if i < 3 {
				if playlist.Image.Len() == 0 {
					t.Logf("playlist %s has no image", playlist.Name)
				} else {
					t.Logf("playlist %s has image of size %v bytes", playlist.Name, playlist.Image.Len())
				}
			}
		}
	})

	t.Run("playlist URL", func(t *testing.T) {
		for i, playlist := range playlists {
			if playlist.URL == "" {
				t.Errorf("playlist URL is empty")
			}
			if i < 3 {
				t.Logf("playlist: %s", playlist.URL)
			}
		}
	})

	if len(playlists) == 0 {
		t.Errorf("expected at least one playlist, got none")
	}
}
