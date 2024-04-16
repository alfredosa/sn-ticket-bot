package routers

import (
	"github.com/alfredosa/sn-ticket-bot/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func Routers(db *sqlx.DB) *chi.Mux {
	api_config := handlers.NewAPIConfig(db)

	r := chi.NewRouter()

	apiRouter := chi.NewRouter()
	apiRouter.Get("/healthz", api_config.HealthHandler)

	r.Mount("/api", apiRouter)

	adminRouter := chi.NewRouter()
	adminRouter.Get("/metrics", api_config.MetricsHandler)
	r.Mount("/admin", adminRouter)
	return r
}
