package database

import (
	"database/sql"
	"dib/internal/models"

	_ "modernc.org/sqlite"
)

type fakeDatabase struct {
	playlists map[string]models.Playlist
}

func NewFakeDatabaseRepo() (Database, error) {
	return &fakeDatabase{
		playlists: make(map[string]models.Playlist),
	}, nil
}

func (db *fakeDatabase) Close() error {
	return nil
}

func (db *fakeDatabase) CreatePlaylist(playlist models.Playlist) error {
	db.playlists[playlist.Id] = playlist
	return nil
}

func (db *fakeDatabase) ListPlaylists() ([]models.Playlist, error) {
	playlists := make([]models.Playlist, 0)

	for _, playlist := range db.playlists {
		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

func (db *fakeDatabase) RetrievePlaylistById(playlistId string) (models.Playlist, error) {
	playlist, ok := db.playlists[playlistId]
	if !ok {
		return models.EmptyPlaylist, sql.ErrNoRows
	}

	return playlist, nil
}

func (db *fakeDatabase) DeletePlaylist(playlistId string) error {
	delete(db.playlists, playlistId)
	return nil
}

func (db *fakeDatabase) UpdatePlaylist(playlist models.Playlist) error {
	db.playlists[playlist.Id] = playlist
	return nil
}
