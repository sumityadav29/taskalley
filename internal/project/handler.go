package project

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/api/v1/projects", func(r chi.Router) {
		r.Get("/", h.getAllProjects)
		r.Post("/", h.createProject)
	})
}

func (h *Handler) getAllProjects(w http.ResponseWriter, r *http.Request) {
	// NOTE: ideally we should get the userId from the request context like auth token but for now I am using query param
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		http.Error(w, "missing userId query parameter", http.StatusBadRequest)
		return
	}

	projects, err := h.service.GetAllProjectsByUser(r.Context(), userId)

	if err != nil {
		http.Error(w, "failed to get projects: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func (h *Handler) createProject(w http.ResponseWriter, r *http.Request) {
	var projectCreate ProjectCreate

	if err := json.NewDecoder(r.Body).Decode(&projectCreate); err != nil {
		http.Error(w, "invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	project, err := h.service.CreateProject(r.Context(), &projectCreate)
	if err != nil {
		http.Error(w, "failed to create project: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}
