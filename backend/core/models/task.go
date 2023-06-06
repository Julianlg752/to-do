package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Task struct {
	ID          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Status      bool      `json:"status" db:"status"`
	UserID      int64     `json:"userId" db:"user_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

func (t *Task) Validate() error {
	return validation.ValidateStruct(t,
		validation.Field(&t.Title, validation.Required),
		validation.Field(&t.Description, validation.Required),
		validation.Field(&t.Status, validation.NotNil),
	)
}
