package spotify

import (
	spotiyou "github.com/schicho/spotiyou/pkg/spotify"
	"github.com/zmb3/spotify/v2"
)

// protoPlaylist is a simplified version of spotify.SimplePlaylist.
// However, it does not contain the tracks of the playlist, as these need to be retrieved separately.
// Therefore it contains the Spotify ID of the playlist, which can be used to retrieve the tracks.
type protoPlaylist struct {
	Name        string
	Description string
	OwnerName   string
	SpotifyID   spotify.ID
}

func toProtoPlaylist(pl spotify.SimplePlaylist) protoPlaylist {
	return protoPlaylist{
		Name:        pl.Name,
		Description: pl.Description,
		OwnerName:   pl.Owner.DisplayName,
		SpotifyID:   pl.ID,
	}
}

func toBasicArtist(a spotify.SimpleArtist) spotiyou.BasicArtist {
	return spotiyou.BasicArtist{
		Name: a.Name,
	}
}

func toBasicTrack(t spotify.SimpleTrack) spotiyou.BasicTrack {
	artists := make([]spotiyou.BasicArtist, 0, len(t.Artists))
	for _, a := range t.Artists {
		artists = append(artists, toBasicArtist(a))
	}

	return spotiyou.BasicTrack{
		Name:    t.Name,
		Artists: artists,
	}
}
