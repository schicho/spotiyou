package spotify

import (
	"errors"
	"os"
	"testing"
)

var TestClient *SpotifyClient

func TestMain(m *testing.M) {
	var err error
	TestClient, err = createTestSpotifyClient()
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func createTestSpotifyClient() (*SpotifyClient, error) {
	var clientID, clientSecret string
	clientID = os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		return nil, errors.New("missing client ID or secret")
	}
	return NewSpotifyClient(clientID, clientSecret)
}

func TestGetUserPlaylists(t *testing.T) {

	const userID = "spotify"

	playlists, err := TestClient.GetUserPlaylists(userID)
	if err != nil {
		t.Errorf("failed to get user playlists: %v", err)
	}

	//for _, playlist := range playlists {
	//	t.Logf("playlist: %s", playlist.Name)
	//}

	if len(playlists) == 0 {
		t.Errorf("expected at least one playlist, got none")
	}
}

func TestGetPlaylistTracks(t *testing.T) {

	const playlistID = "37i9dQZF1DXcBWIGoYBM5M"

	tracks, err := TestClient.GetPlaylistTracks(playlistID)
	if err != nil {
		t.Errorf("failed to get playlist tracks: %v", err)
	}

	//for _, track := range tracks {
	//	t.Logf("track: %s", track.Name)
	//}

	if len(tracks) == 0 {
		t.Errorf("expected at least one track, got none")
	}
}
