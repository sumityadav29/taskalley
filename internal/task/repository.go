package task

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, task *TaskCreate) (*Task, error)
	// GetAllByProject(ctx context.Context, projectId string) ([]*Task, error)
	// GetById(ctx context.Context, id string) (*Task, error)
	// UpdateById(ctx context.Context, id string, task *Task) (*Task, error)
	// DeleteById(ctx context.Context, id string) error
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, task *TaskCreate) (*Task, error) {
	projectId := task.ProjectId
	title := task.Title
	description := task.Description
	dueDate := task.DueDate
	createdBy := task.CreatedBy

	var id string
	var createdAt time.Time
	var updatedAt time.Time
	var status Status
	err := r.db.QueryRow(ctx, `
		INSERT INTO tasks (project_id, title, description, due_date, created_by)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, title, description, due_date, created_by, created_at, updated_at, status
	`, projectId, title, description, dueDate, createdBy).Scan(&id, &title, &description, &dueDate, &createdBy, &createdAt, &updatedAt, &status)

	if err != nil {
		return nil, err
	}

	return &Task{
		Id:          id,
		ProjectId:   projectId,
		Title:       title,
		Description: description,
		Status:      Status(status),
		DueDate:     dueDate,
		CreatedBy:   createdBy,
		CreatedAt:   createdAt,
	}, nil
}
