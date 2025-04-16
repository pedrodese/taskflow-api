package schema

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID          string `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Title       string `json:"title" gorm:"not null" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (t *Task) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == "" {
		t.ID = uuid.New().String()
	}
	return nil
}
