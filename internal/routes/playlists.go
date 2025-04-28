package routes

import (
	"dib/internal/components"
	"dib/internal/models"
	"dib/internal/repositories/database"
	"errors"

	"github.com/pocketbase/pocketbase/core"
)

func playlistList(e *core.RequestEvent) error {
	db := database.NewDatabaseRepo(e.App)
	playlists, err := db.ListPlaylists()
	if err != nil {
		return err
	}

	return sendPage(e, components.PlaylistListPage(playlists))
}

func playlistNew(e *core.RequestEvent) error {
	db := database.NewDatabaseRepo(e.App)

	playlist := models.Playlist{
		Id:   models.GenerateId(),
		Name: "New Playlist",
	}

	if err := db.CreatePlaylist(playlist); err != nil {
		return err
	}

	return e.Redirect(302, "/edit/"+playlist.Id)
}

func playlistEdit(e *core.RequestEvent) error {
	playlistId := e.Request.PathValue("id")
	if playlistId == "" {
		return errors.New("playlist id is required")
	}

	db := database.NewDatabaseRepo(e.App)
	playlist, err := db.RetrievePlaylistById(playlistId)

	if err != nil {
		return err
	}

	return sendPage(e, components.PlaylistEditPage(playlist))
}

type PlaylistChange struct {
	Name string `json:"name" form:"name"`
}

func playlistSave(e *core.RequestEvent) error {
	playlistId := e.Request.PathValue("id")
	if playlistId == "" {
		return errors.New("playlist id is required")
	}

	var changes PlaylistChange
	if err := e.BindBody(&changes); err != nil {
		return err
	}

	db := database.NewDatabaseRepo(e.App)
	playlist, err := db.RetrievePlaylistById(playlistId)
	if err != nil {
		return err
	}

	if changes.Name != "" {
		playlist.Name = changes.Name
	}

	if err := db.UpdatePlaylist(playlist); err != nil {
		return err
	}

	return e.String(200, "ok")
}
