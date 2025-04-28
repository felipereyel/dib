package routes

import (
	"dib/internal/embeded"

	"github.com/pocketbase/pocketbase/core"

	"github.com/pocketbase/pocketbase/apis"
)

var assetsHandler = apis.Static(embeded.Assets, true)

func healthzHandler(e *core.RequestEvent) error {
	return e.String(200, "ok")
}

func SetupRoutes(se *core.ServeEvent) error {
	se.Router.GET("/", playlistList)
	se.Router.GET("/new", playlistNew)
	se.Router.GET("/edit/{id}", playlistEdit)
	se.Router.POST("/edit/{id}", playlistSave)

	se.Router.GET("/healthz", healthzHandler)
	se.Router.GET("/statics/{path...}", assetsHandler)

	return se.Next()
}
