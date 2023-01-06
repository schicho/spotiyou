package spotify

import (
	spotiyou "github.com/schicho/spotiyou/pkg/spotify"
)

func (sc *SpotifyClient) getUserProtoPlaylists(userID string) ([]protoPlaylist, error) {
	playlists, err := sc.getUserPlaylists(userID)
	if err != nil {
		return nil, err
	}

	protoPlaylists := make([]protoPlaylist, 0, len(playlists))
	for _, pl := range playlists {
		protoPlaylists = append(protoPlaylists, toProtoPlaylist(pl))
	}
	return protoPlaylists, nil
}

func (sc *SpotifyClient) getBasicPlaylist(pl protoPlaylist) (spotiyou.BasicPlaylist, error) {
	tracks, err := sc.getPlaylistTracks(pl.SpotifyID)
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
		OwnerName:   pl.OwnerName,
		Tracks:      basicTracks,
	}, nil
}

func (sc *SpotifyClient) GetUserPlaylists(userID string) ([]spotiyou.BasicPlaylist, error) {
	protoPlaylists, err := sc.getUserProtoPlaylists(userID)
	if err != nil {
		return nil, err
	}

	basicPlaylists := make([]spotiyou.BasicPlaylist, 0, len(protoPlaylists))
	for _, pl := range protoPlaylists {
		basicPlaylist, err := sc.getBasicPlaylist(pl)
		if err != nil {
			return nil, err
		}

		basicPlaylists = append(basicPlaylists, basicPlaylist)
	}
	return basicPlaylists, nil
}
