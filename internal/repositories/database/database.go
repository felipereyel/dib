package database

import (
	"dib/internal/models"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

type database struct {
	app core.App
}

func NewDatabaseRepo(app core.App) database {
	return database{app}
}

func (db *database) CreatePlaylist(playlist models.Playlist) error {
	query := `INSERT INTO playlists (id, name) VALUES ({:Id}, {:Name})`
	q := db.app.DB().NewQuery(query).Bind(dbx.Params{
		"Id":   playlist.Id,
		"Name": playlist.Name,
	})

	_, err := q.Execute()
	return err
}

func (db *database) ListPlaylists() ([]models.Playlist, error) {
	query := `SELECT id, name FROM playlists`
	q := db.app.DB().NewQuery(query)

	playlists := make([]models.Playlist, 0)
	if err := q.All(&playlists); err != nil {
		return nil, err
	}

	return playlists, nil
}

func (db *database) RetrievePlaylistById(playlistId string) (models.Playlist, error) {
	query := `SELECT id, name FROM playlists WHERE id = {:Id}`
	q := db.app.DB().NewQuery(query).Bind(dbx.Params{
		"Id": playlistId,
	})

	var playlist models.Playlist
	err := q.One(&playlist)
	return playlist, err
}

func (db *database) DeletePlaylist(playlistId string) error {
	query := `DELETE FROM playlists WHERE id = {:Id}`
	q := db.app.DB().NewQuery(query).Bind(dbx.Params{
		"Id": playlistId,
	})

	_, err := q.Execute()
	return err
}

func (db *database) UpdatePlaylist(playlist models.Playlist) error {
	query := `UPDATE playlists SET name = {:Name} WHERE id = {:Id}`
	q := db.app.DB().NewQuery(query).Bind(dbx.Params{
		"Id":   playlist.Id,
		"Name": playlist.Name,
	})

	_, err := q.Execute()
	return err
}
