package task

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
	r.Route("/api/v1/projects/{projectId}/tasks", func(r chi.Router) {
		r.Post("/", h.Create)
	})
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var taskCreate TaskCreate

	if err := json.NewDecoder(r.Body).Decode(&taskCreate); err != nil {
		http.Error(w, "invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	// NOTE: ideally we should get the userId from the request context like auth token but for now I am using query param
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		http.Error(w, "userId is required", http.StatusBadRequest)
		return
	}

	task, err := h.service.Create(r.Context(), &taskCreate, userId)

	if err != nil {
		http.Error(w, "failed to create task: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}
