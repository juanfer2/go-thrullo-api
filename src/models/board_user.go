package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
)

// User is used by pop to map your Users database table to your go code.
type BoardUser struct {
	ID        int       `json:"id" db:"id"`
	BoardID   int       `db:"board_id"`
	UserID    int       `db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// TableName overrides the table name used by Pop.
func (u BoardUser) TableName() string {
	return "BoardUsers"
}

// String is not required by pop and may be deleted
func (boardUser BoardUser) String() string {
	jt, _ := json.Marshal(boardUser)
	return string(jt)
}

// Useres is not required by pop and may be deleted
type BoardUsers []BoardUser

// String is not required by pop and may be deleted
func (boardUsers BoardUsers) String() string {
	jt, _ := json.Marshal(boardUsers)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (boardUser *BoardUser) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (boardUser *BoardUser) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (boardUser *BoardUser) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
