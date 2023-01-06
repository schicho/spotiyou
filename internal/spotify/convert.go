package spotify

import (
	spotiyou "github.com/schicho/spotiyou/pkg/playlist"
	"github.com/zmb3/spotify/v2"
)

// toBasicArtist converts a spotify.SimpleArtist to a spotiyou.BasicArtist.
func toBasicArtist(a spotify.SimpleArtist) spotiyou.BasicArtist {
	return spotiyou.BasicArtist{
		Name: a.Name,
	}
}

// toBasicTrack converts a spotify.SimpleTrack to a spotiyou.BasicTrack.
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
