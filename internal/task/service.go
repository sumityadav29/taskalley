package task

import (
	"context"
	"errors"

	"github.com/sumityadav29/taskalley/internal/applicationevents"
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
	repo     Repository
	eventBus *applicationevents.EventBus
}

func NewService(repo Repository, eventBus *applicationevents.EventBus) Service {
	return &service{repo: repo, eventBus: eventBus}
}

func (s *service) Create(ctx context.Context, taskCreate *TaskCreate, userId string) (*Task, error) {
	if taskCreate.ProjectId == "" {
		return nil, errors.New("projectId is required")
	}

	if taskCreate.Title == "" {
		return nil, errors.New("title is required")
	}

	if taskCreate.DueDate.IsZero() {
		return nil, errors.New("dueDate is required")
	}

	if userId == "" {
		return nil, errors.New("userId is required")
	}

	createdTask, err := s.repo.Create(ctx, taskCreate)
	if err != nil {
		return nil, err
	}

	createdTask.CreatedBy = userId

	s.eventBus.Publish(applicationevents.TaskCreated, createdTask)

	return createdTask, nil
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

	err := s.repo.DeleteById(ctx, id)
	if err != nil {
		return err
	}

	s.eventBus.Publish(applicationevents.TaskDeleted, id)

	return nil
}

func (s *service) UpdateById(ctx context.Context, id string, task *TaskUpdate) (*Task, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	updatedTask, err := s.repo.UpdateById(ctx, id, task)
	if err != nil {
		return nil, err
	}

	s.eventBus.Publish(applicationevents.TaskUpdated, updatedTask)
	return updatedTask, nil
}

func (s *service) UpdateStatus(ctx context.Context, id string, status Status) (*Task, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	if status == "" {
		return nil, errors.New("status is required")
	}

	updatedTask, err := s.repo.UpdateStatus(ctx, id, status)
	if err != nil {
		return nil, err
	}

	s.eventBus.Publish(applicationevents.TaskUpdated, updatedTask)
	return updatedTask, nil
}
