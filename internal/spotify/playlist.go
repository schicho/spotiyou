package spotify

import (
	"context"

	"github.com/zmb3/spotify/v2"
)

func (sc *SpotifyClient) GetUserPlaylists(userID string) ([]spotify.SimplePlaylist, error) {
	playlistPage, err := sc.client.GetPlaylistsForUser(context.Background(), userID)
	if err != nil {
		sc.logger.Printf("failed to get playlists for user %s: %v", userID, err)
		return nil, ErrRetrievalFailed
	}

	return playlistPage.Playlists, nil
}

func (sc *SpotifyClient) GetPlaylistTracks(playlistID spotify.ID) ([]spotify.SimpleTrack, error) {
	playlistItemPage, err := sc.client.GetPlaylistItems(context.Background(), playlistID)
	if err != nil {
		sc.logger.Printf("failed to get playlist tracks for playlist %s: %v", playlistID, err)
		return nil, ErrRetrievalFailed
	}

	return playlistItems2SimpleTracks(playlistItemPage.Items), nil
}

func playlistItems2SimpleTracks(items []spotify.PlaylistItem) []spotify.SimpleTrack {
	tracks := make([]spotify.SimpleTrack, len(items))
	for i, item := range items {
		tracks[i] = item.Track.Track.SimpleTrack
	}

	return tracks
}
