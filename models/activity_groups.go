package models

import (
	"errors"
	"net/mail"
	"time"
)

type ActivityGroup struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (activityGroup ActivityGroup) Validate() error {
	if activityGroup.Title == "" {
		return errors.New("title cannot be null")
	}

	if activityGroup.Email != "" {
		_, err := mail.ParseAddress(activityGroup.Email)
		if err != nil {
			return errors.New("invalid email address")
		}
	}

	return nil
}
