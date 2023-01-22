package spotify

import (
	"bytes"
	"sync"

	spotiyou "github.com/schicho/spotiyou/pkg/playlist"
	"github.com/zmb3/spotify/v2"
)

const openSpotifyPlaylistBaseURL = "https://open.spotify.com/playlist/"

func (sc *SpotifyClient) getBasicPlaylist(pl spotify.SimplePlaylist) (spotiyou.BasicPlaylist, error) {
	tracks, err := sc.apiGetPlaylistTracks(pl.ID)
	if err != nil {
		sc.logger.Printf("failed to get playlist tracks for playlist %s %s: %v", pl.ID, pl.Name, err)
		return spotiyou.BasicPlaylist{}, err
	}

	basicTracks := make([]spotiyou.BasicTrack, 0, len(tracks))
	for _, t := range tracks {
		basicTracks = append(basicTracks, toBasicTrack(t))
	}

	imgBuffer := &bytes.Buffer{}
	if pl.Images != nil && len(pl.Images) > 0 {
		pl.Images[0].Download(imgBuffer)
	}

	return spotiyou.BasicPlaylist{
		Name:        pl.Name,
		Description: pl.Description,
		OwnerName:   pl.Owner.DisplayName,
		Tracks:      basicTracks,
		Image:       imgBuffer,
		URL:         openSpotifyPlaylistBaseURL + pl.ID.String(),
	}, nil
}

// GetUserPlaylists returns all playlists of a user.
//
// The data is fetched from the Spotify API. It returns a slice of
// BasicPlaylist which contains only the most important information
// about the playlist and its tracks.
//
// The function may take a significant amount of time to complete.
// Multiple parallel API calls are made to fetch all playlists and their tracks.
func (sc *SpotifyClient) GetUserPlaylists(userID string) ([]spotiyou.BasicPlaylist, error) {
	apiPlaylists, err := sc.apiGetUserPlaylists(userID)
	if err != nil {
		return nil, err
	}

	basicPlaylists := make([]spotiyou.BasicPlaylist, 0, len(apiPlaylists))

	// download playlist tracks in parallel to speed up the otherwise
	// sequential process with blocking API calls.
	downloadChan := make(chan spotiyou.BasicPlaylist)
	wg := sync.WaitGroup{}

	for _, pl := range apiPlaylists {
		wg.Add(1)
		go func(pl spotify.SimplePlaylist) {
			defer wg.Done()
			basicPlaylist, err := sc.getBasicPlaylist(pl)
			if err != nil {
				return
			}
			downloadChan <- basicPlaylist
		}(pl)
	}

	go func() {
		wg.Wait()
		close(downloadChan)
	}()

	for pl := range downloadChan {
		basicPlaylists = append(basicPlaylists, pl)
	}

	return basicPlaylists, nil
}
