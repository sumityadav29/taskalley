package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/sumityadav29/taskalley/config"
	"github.com/sumityadav29/taskalley/internal/project"
)

func main() {
	cfg := config.Load()

	dbPool, err := pgxpool.New(context.Background(), cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer dbPool.Close()

	projectRepo := project.NewRepository(dbPool)
	projectService := project.NewService(projectRepo)
	projectHandler := project.NewHandler(projectService)

	r := chi.NewRouter()
	projectHandler.RegisterRoutes(r)

	log.Printf("starting server on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
