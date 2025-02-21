package models

import (
	"github.com/google/uuid"
)

type Alert struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Email      string    `json:"email"`
	ResourceID uuid.UUID `json:"resource_id" gorm:"type:uuid"`
	IsAll      bool      `json:"isAll"` // Renommer 'all' en 'IsAll'
}
