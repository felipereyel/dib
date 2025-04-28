package database

import (
	"dib/internal/models"

	_ "modernc.org/sqlite"
)

type Database interface {
	Close() error

	CreatePlaylist(playlist models.Playlist) error
	RetrievePlaylistById(playlistId string) (models.Playlist, error)
	UpdatePlaylist(playlist models.Playlist) error
	DeletePlaylist(playlistId string) error
	ListPlaylists() ([]models.Playlist, error)
}
