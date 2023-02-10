package downloader

import (
	d "github.com/Carbohz/gtfs-viewer/api/rest/downloader/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func setupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Compress(5))

	r.Get("/download", d.WebArchiveDownloadHandler())

	return r
}
