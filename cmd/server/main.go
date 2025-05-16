package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sumityadav29/taskalley/config"
	"github.com/sumityadav29/taskalley/internal/applicationevents"
	"github.com/sumityadav29/taskalley/internal/middlewares"
	"github.com/sumityadav29/taskalley/internal/project"
	"github.com/sumityadav29/taskalley/internal/task"
)

func main() {
	cfg := config.Load()

	dbPool, err := pgxpool.New(context.Background(), cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer dbPool.Close()

	server, err := wireApplication(dbPool)
	if err != nil {
		log.Fatalf("failed to wire application: %v", err)
	}

	log.Printf("starting server on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, server); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func wireApplication(dbPool *pgxpool.Pool) (*chi.Mux, error) {
	eventBus := applicationevents.NewEventBus()
	applicationevents.NewKafkaApplicationEventHandler(eventBus)

	projectRepo := project.NewRepository(dbPool)
	projectService := project.NewService(projectRepo, eventBus)
	projectHandler := project.NewHandler(projectService)

	taskRepo := task.NewRepository(dbPool)
	taskService := task.NewService(taskRepo, eventBus)
	taskHandler := task.NewHandler(taskService)

	server := chi.NewRouter()
	server.Use(middlewares.AuthMiddleware)
	projectHandler.RegisterRoutes(server)
	taskHandler.RegisterRoutes(server)

	ServeDocs(server)

	return server, nil
}

func ServeDocs(r chi.Router) {
	fs := http.FileServer(http.Dir("./static/docs"))
	r.Handle("/docs/*", http.StripPrefix("/docs/", fs))
}
