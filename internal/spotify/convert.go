package spotify

import (
	spotiyou "github.com/schicho/spotiyou/pkg/spotify"
	"github.com/zmb3/spotify/v2"
)

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
