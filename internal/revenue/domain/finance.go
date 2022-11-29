package domain

import (
	"github.com/google/uuid"
	"time"
)

type Registry struct {
	Id           uuid.UUID
	Value        float64
	Status       string
	Description  string
	DueDate      time.Time
	CreationDate time.Time
}

type Revenue struct {
	Registry
}

type Expenditure struct {
	Registry
}
