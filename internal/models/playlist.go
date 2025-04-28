package models

type Playlist struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var EmptyPlaylist = Playlist{}
