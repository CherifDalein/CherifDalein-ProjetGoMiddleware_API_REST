package models

import (
	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID   `json:"id" db:"id"`
	Summary     string      `json:"summary" db:"summary"`
	Description string      `json:"description" db:"description"`
	Location    string      `json:"location" db:"location"`
	Start       string      `json:"start" db:"start"`
	End         string      `json:"end" db:"end"`
	UID         string      `json:"uid" db:"uid"`
	ResourceId  []uuid.UUID `json:"resource_id" db:"resource_id"`
}
