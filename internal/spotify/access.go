package spotify

import (
	spotiyou "github.com/schicho/spotiyou/pkg/playlist"
	"github.com/zmb3/spotify/v2"
)

func (sc *SpotifyClient) getBasicPlaylist(pl spotify.SimplePlaylist) (spotiyou.BasicPlaylist, error) {
	tracks, err := sc.apiGetPlaylistTracks(pl.ID)
	if err != nil {
		return spotiyou.BasicPlaylist{}, err
	}

	basicTracks := make([]spotiyou.BasicTrack, 0, len(tracks))
	for _, t := range tracks {
		basicTracks = append(basicTracks, toBasicTrack(t))
	}

	return spotiyou.BasicPlaylist{
		Name:        pl.Name,
		Description: pl.Description,
		OwnerName:   pl.Owner.DisplayName,
		Tracks:      basicTracks,
	}, nil
}

func (sc *SpotifyClient) GetUserPlaylists(userID string) ([]spotiyou.BasicPlaylist, error) {
	apiPlaylists, err := sc.apiGetUserPlaylists(userID)
	if err != nil {
		return nil, err
	}

	basicPlaylists := make([]spotiyou.BasicPlaylist, 0, len(apiPlaylists))
	for _, pl := range apiPlaylists {
		basicPlaylist, err := sc.getBasicPlaylist(pl)
		if err != nil {
			return nil, err
		}

		basicPlaylists = append(basicPlaylists, basicPlaylist)
	}
	return basicPlaylists, nil
}
