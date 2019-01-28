package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type Daily struct {
	ID         uuid.UUID `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	Name       string    `json:"name" db:"name"`
	ResetAt    time.Time `json:"reset_at" db:"reset_at"`
	Recurrence int       `json:"recurrence" db:"recurrence"`
}

// String is not required by pop and may be deleted
func (d Daily) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Dailies is not required by pop and may be deleted
type Dailies []Daily

// String is not required by pop and may be deleted
func (d Dailies) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (d *Daily) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (d *Daily) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	if (d.ResetAt == time.Time{}) {
		d.ResetAt = time.Now()
	}
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (d *Daily) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
