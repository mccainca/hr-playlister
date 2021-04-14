package main

import (
	"fmt"

	"github.com/mccainca/hr-playlister/playlist"
	"github.com/mccainca/hr-playlister/scraper"
)

const playlistId = "3rrzkFFuWGIU8ATzkUlChp"

func main() {

	var trackList = scraper.Scrape("2021/02/20")
	for _, v := range trackList {
		fmt.Printf("%+v\n", v)
		if v.SpotifyID != "" {
			err := playlist.AddTrackToSpotifyPlaylist(v.SpotifyID, playlistId)
			if err != nil {
				fmt.Printf("Error: %s", err)
			}
		}
	}
}
