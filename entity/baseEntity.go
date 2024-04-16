package entity

import (
	"github.com/google/uuid"
	"time"
)

type IModel interface {
	CompareUniqueColumns(model IModel) error
	ValidationModel() error
}

type AbstractEntity struct {
	IModel
	Id          int        `json:"id" db:"id"`
	UserId      string     `json:"userId" db:"user_id"`
	CreatedDate *time.Time `json:"createdDate" db:"created_date"`
	UpdatedBy   uuid.UUID  `json:"updatedBy" db:"updated_by"`
	UpdatedDate *time.Time `json:"updatedDate" db:"updated_date"`
}
