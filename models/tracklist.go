package models

//TrackList : Our struct for capturing relevant data for spotify lookup
type TrackList []struct {
	// AffiliateLinkiPhone  string `json:"affiliateLinkiPhone"`
	// ProgramStart         string `json:"program_start"`
	// AffiliateLinkSpotify string `json:"affiliateLinkSpotify"`
	// PlayID               int    `json:"play_id"`
	// ItunesID             string `json:"itunes_id"`
	// ProgramID            string `json:"program_id"`
	// Datetime             string `json:"datetime"`
	// ProgramEnd           string `json:"program_end"`
	// AffiliateLinkRdio    string `json:"affiliateLinkRdio"`
	// SpotifyPreview       string `json:"spotify_preview"`
	// AlbumImage           string `json:"albumImage"`
	// Year                 int    `json:"year"`
	// Date                 string `json:"date"`
	// ProgramTitle         string `json:"program_title"`
	// Time                 string `json:"time"`
	// AlbumImageLarge      string `json:"albumImageLarge"`
	Album string `json:"album"`
	// Offset               int    `json:"offset"`
	SpotifyID string `json:"spotify_id"`
	Title     string `json:"title"`
	Artist    string `json:"artist"`
	// Host                 string `json:"host"`
	// Comments             string `json:"comments"`
	// Label                string `json:"label"`
	// AffiliateLinkiTunes  string `json:"affiliateLinkiTunes"`
	// ArtistURL            string `json:"artist_url"`
	// AffiliateLinkAmazon  string `json:"affiliateLinkAmazon"`
	// ItunesTime           int    `json:"itunes_time"`
	// ItunesURL            string `json:"itunes_url"`
	// Channel              string `json:"channel"`
}
