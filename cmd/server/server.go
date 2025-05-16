package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/sumityadav29/taskalley/config"
	"github.com/sumityadav29/taskalley/internal/middlewares"
	"github.com/sumityadav29/taskalley/internal/project"
	"github.com/sumityadav29/taskalley/internal/task"
)

func NewServer() http.Handler {
	dbConfig := config.Load()

	dbPool, err := pgxpool.New(context.Background(), dbConfig.DatabaseUrl)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer dbPool.Close()

	projectRepo := project.NewRepository(dbPool)
	projectService := project.NewService(projectRepo)
	projectHandler := project.NewHandler(projectService)

	taskRepo := task.NewRepository(dbPool)
	taskService := task.NewService(taskRepo)
	taskHandler := task.NewHandler(taskService)

	r := chi.NewRouter()

	r.Use(middlewares.AuthMiddleware)

	projectHandler.RegisterRoutes(r)
	taskHandler.RegisterRoutes(r)

	return r
}
