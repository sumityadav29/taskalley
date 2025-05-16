package task

import (
	"context"
	"errors"

	"github.com/sumityadav29/taskalley/internal/task/taskfilters"
)

type Service interface {
	Create(ctx context.Context, task *TaskCreate, userId string) (*Task, error)
	GetAllByProject(ctx context.Context, projectId string, status Status, start int, limit int) ([]*Task, error)
	GetById(ctx context.Context, id string) (*Task, error)
	DeleteById(ctx context.Context, id string) error
	UpdateById(ctx context.Context, id string, task *TaskUpdate) (*Task, error)
	UpdateStatus(ctx context.Context, id string, status Status) (*Task, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, task *TaskCreate, userId string) (*Task, error) {
	if task.ProjectId == "" {
		return nil, errors.New("projectId is required")
	}

	if task.Title == "" {
		return nil, errors.New("title is required")
	}

	if task.DueDate.IsZero() {
		return nil, errors.New("dueDate is required")
	}

	if userId == "" {
		return nil, errors.New("userId is required")
	} else {
		task.CreatedBy = userId
	}

	return s.repo.Create(ctx, task)
}

func (s *service) GetAllByProject(ctx context.Context, projectId string, status Status, start int, limit int) ([]*Task, error) {
	if projectId == "" {
		return nil, errors.New("projectId is required")
	}

	var filters []taskfilters.TaskFilter

	projectIdFilter := taskfilters.StringMatchTaskFilter{
		ColumnName:  "project_id",
		ColumnValue: projectId,
	}

	filters = append(filters, &projectIdFilter)

	if status != "" {
		statusFilter := taskfilters.StringMatchTaskFilter{
			ColumnName:  "status",
			ColumnValue: string(status),
		}
		filters = append(filters, &statusFilter)
	}

	return s.repo.GetAllByFilters(ctx, filters, start, limit)
}

func (s *service) GetById(ctx context.Context, id string) (*Task, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.repo.GetById(ctx, id)
}

func (s *service) DeleteById(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return s.repo.DeleteById(ctx, id)
}

func (s *service) UpdateById(ctx context.Context, id string, task *TaskUpdate) (*Task, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.repo.UpdateById(ctx, id, task)
}

func (s *service) UpdateStatus(ctx context.Context, id string, status Status) (*Task, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	if status == "" {
		return nil, errors.New("status is required")
	}

	return s.repo.UpdateStatus(ctx, id, status)
}
