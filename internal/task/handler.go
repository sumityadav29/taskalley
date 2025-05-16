package task

import (
	"encoding/json"
	"net/http"
	"strconv"

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
		r.Get("/", h.GetAllByProject)
	})
	r.Route("/api/v1/projects/{projectId}/tasks/{taskId}", func(r chi.Router) {
		r.Get("/", h.GetById)
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

func (h *Handler) GetAllByProject(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	status := Status(r.URL.Query().Get("status"))

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}

	start := 0
	if val := r.URL.Query().Get("start"); val != "" {
		if parsed, err := strconv.Atoi(val); err == nil {
			start = parsed
		}
	}

	limit := 10
	if val := r.URL.Query().Get("limit"); val != "" {
		if parsed, err := strconv.Atoi(val); err == nil {
			limit = parsed
		}
	}

	tasks, err := h.service.GetAllByProject(r.Context(), projectId, status, start, limit)

	if err != nil {
		http.Error(w, "failed to get tasks: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	taskId := chi.URLParam(r, "taskId")

	if taskId == "" {
		http.Error(w, "taskId is required", http.StatusBadRequest)
		return
	}

	task, err := h.service.GetById(r.Context(), taskId)

	if err != nil {
		http.Error(w, "failed to get task: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}
