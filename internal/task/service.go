package task

import (
	"context"
	"errors"
)

type Service interface {
	Create(ctx context.Context, task *TaskCreate, userId string) (*Task, error)
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
