package main

import (
	"database/sql"
	"log"
	"log/slog"
	"net/http"
	"os"

	"calc_service/internal/config"
	"calc_service/internal/handlers"
	"calc_service/internal/middleware"
	"calc_service/internal/server"
	"calc_service/internal/storage"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, nil),
	)

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	repo := storage.NewPostgresRepository(db)
	h := handlers.New(repo)

	mux := http.NewServeMux()

	server.RegisterRoutes(mux, h)

	handler := middleware.RequestID(mux)
	handler = middleware.Logging(logger)(handler)
	handler = middleware.Auth(cfg.AuthSecret)(handler)
	handler = middleware.RateLimit(cfg.RateLimit, cfg.RateBurst)(handler)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: handler,
	}

	logger.Info("server starting", "port", cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
