package spotify

import (
	"context"

	"github.com/zmb3/spotify/v2"
)

func (sc *SpotifyClient) apiGetUserPlaylists(userID string) ([]spotify.SimplePlaylist, error) {
	playlistPage, err := sc.client.GetPlaylistsForUser(context.Background(), userID)
	if err != nil {
		sc.logger.Printf("failed to get playlists for user %s: %v", userID, err)
		return nil, ErrRetrievalFailed
	}

	return playlistPage.Playlists, nil
}

func (sc *SpotifyClient) apiGetPlaylistTracks(playlistID spotify.ID) ([]spotify.SimpleTrack, error) {
	playlistItemPage, err := sc.client.GetPlaylistItems(context.Background(), playlistID)
	if err != nil {
		sc.logger.Printf("failed to get playlist tracks for playlist %s: %v", playlistID, err)
		return nil, ErrRetrievalFailed
	}

	return playlistItems2SimpleTracks(playlistItemPage.Items), nil
}

// playlistItems2SimpleTracks converts a slice of spotify.PlaylistItem to a slice of spotify.SimpleTrack.
//
// Spotify's API distinguishes between podcasts and music tracks, so we need to find our way to the track data.
// Additionally, this library takes type embedding to the extreme, so we have to do this to get some basic information
// like the track name and artist.
func playlistItems2SimpleTracks(items []spotify.PlaylistItem) []spotify.SimpleTrack {
	tracks := make([]spotify.SimpleTrack, 0, len(items))
	for _, item := range items {
		// In case we pass a podcast, we skip it.
		if item.Track.Track == nil {
			continue
		}

		tracks = append(tracks, item.Track.Track.SimpleTrack)
	}

	return tracks
}
