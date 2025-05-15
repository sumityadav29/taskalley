package project

import "time"

type Project struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ProjectCreate struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	CreatedBy   string `json:"createdBy"`
}
