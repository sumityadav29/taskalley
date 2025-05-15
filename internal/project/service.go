package project

import (
	"context"
	"errors"
)

type Service interface {
	CreateProject(ctx context.Context, projectCreate *ProjectCreate) (*Project, error)
	GetAllProjectsByUser(ctx context.Context, userId string) ([]*Project, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateProject(ctx context.Context, projectCreate *ProjectCreate) (*Project, error) {
	if projectCreate.Name == "" {
		return nil, errors.New("project name is required")
	}

	return s.repo.Create(ctx, projectCreate)
}

func (s *service) GetAllProjectsByUser(ctx context.Context, userId string) ([]*Project, error) {
	if userId == "" {
		return nil, errors.New("userId is required")
	}
	return s.repo.GetAllByUser(ctx, userId)
}
