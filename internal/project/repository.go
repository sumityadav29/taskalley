package project

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, p *ProjectCreate) (*Project, error)
	GetAllByUser(ctx context.Context, userId string) ([]*Project, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (repo *repository) Create(ctx context.Context, p *ProjectCreate) (*Project, error) {
	now := time.Now()

	var id string
	var description string

	if p.Description != "" {
		description = p.Description
	}

	err := repo.db.QueryRow(ctx, `
		INSERT INTO projects (name, description, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, description
	`, p.Name, description, p.CreatedBy, now, now).Scan(&id, &description)

	if err != nil {
		return nil, fmt.Errorf("create project: %w", err)
	}

	return &Project{
		Id:          id,
		Name:        p.Name,
		Description: description,
		CreatedBy:   p.CreatedBy,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (repo *repository) GetAllByUser(ctx context.Context, userId string) ([]*Project, error) {
	rows, err := repo.db.Query(ctx, `
		SELECT id, name, description, created_by, created_at, updated_at
		FROM projects
		WHERE created_by = $1
		ORDER BY created_at DESC
	`, userId)

	if err != nil {
		return nil, fmt.Errorf("get all projects by user: %w", err)
	}
	defer rows.Close()

	var projects []*Project

	for rows.Next() {
		var project Project
		err := rows.Scan(&project.Id, &project.Name, &project.Description, &project.CreatedBy, &project.CreatedAt, &project.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("scan project: %w", err)
		}
		projects = append(projects, &project)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error: %w", rows.Err())
	}

	return projects, nil
}
