package models

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
)

// User is used by pop to map your Users database table to your go code.
type Board struct {
	ID          int         `json:"id" db:"id"`
	Title       string      `json:"title" db:"title"`
	Description string      `json:"description" db:"description"`
	BoardUsers  []BoardUser `json:"board_users,omitempty" has_many:"board_users"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
}

func (board Board) Type() {
	r := reflect.TypeOf(board)
	fmt.Println(r)
}

// TableName overrides the table name used by Pop.
func (u Board) TableName() string {
	return "boards"
}

// String is not required by pop and may be deleted
func (board Board) String() string {
	jt, _ := json.Marshal(board)
	return string(jt)
}

// Useres is not required by pop and may be deleted
type Boards []Board

// String is not required by pop and may be deleted
func (boards Boards) String() string {
	jt, _ := json.Marshal(boards)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (board *Board) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (board *Board) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (board *Board) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
