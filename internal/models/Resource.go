package models

import (
	"github.com/google/uuid"
)

type Resource struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name  string    `json:"name"`
	UcaID int       `json:"ucaId"`
}
