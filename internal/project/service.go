package project

import (
	"context"
	"errors"

	"github.com/sumityadav29/taskalley/internal/applicationevents"
)

type Service interface {
	CreateProject(ctx context.Context, projectCreate *ProjectCreate) (*Project, error)
	GetAllProjectsByUser(ctx context.Context, userId string) ([]*Project, error)
}

type service struct {
	repo     Repository
	eventBus *applicationevents.EventBus
}

func NewService(repo Repository, eventBus *applicationevents.EventBus) Service {
	return &service{repo: repo, eventBus: eventBus}
}

func (s *service) CreateProject(ctx context.Context, projectCreate *ProjectCreate) (*Project, error) {
	if projectCreate.Name == "" {
		return nil, errors.New("project name is required")
	}

	project, err := s.repo.Create(ctx, projectCreate)
	if err != nil {
		return nil, err
	}

	s.eventBus.Publish(applicationevents.ProjectCreated, project)
	return project, nil
}

func (s *service) GetAllProjectsByUser(ctx context.Context, userId string) ([]*Project, error) {
	if userId == "" {
		return nil, errors.New("userId is required")
	}
	return s.repo.GetAllByUser(ctx, userId)
}
