package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/juanfer2/go-thrullo-api.git/src/password_hash"
)

// User is used by pop to map your Users database table to your go code.
type User struct {
	ID           int       `json:"id" db:"id"`
	FirstName    string    `json:"first_name" db:"first_name"`
	LastName     string    `json:"last_name" db:"first_name"`
	Username     string    `json:"username" db:"username"`
	Email        string    `json:"email" db:"email"`
	Token        string    `json:"token" db:"token"`
	PasswordHash string    `json:"password_hash" db:"password_hash"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// TableName overrides the table name used by Pop.
func (u User) TableName() string {
	return "Users"
}

// String is not required by pop and may be deleted
func (user User) String() string {
	jt, _ := json.Marshal(user)
	return string(jt)
}

// Useres is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (user Users) String() string {
	jt, _ := json.Marshal(user)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (user *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (user *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (user *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (user *User) HassPassword(password string) {
	user.PasswordHash, _ = password_hash.HashPassword(password)
}

func (user *User) CheckPassword(password string) bool {
	passwordHash := user.PasswordHash

	return password_hash.CheckPasswordHash(password, passwordHash)
}
