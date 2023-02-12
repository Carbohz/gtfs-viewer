package rest

import (
	d "github.com/Carbohz/gtfs-viewer/api/rest/handlers"
	"github.com/go-chi/chi"
)

func setupRouter() *chi.Mux {
	r := chi.NewRouter()
	//r.Use(middleware.Compress(5))

	r.Get("/download", d.WebArchiveDownloadHandler())

	return r
}
