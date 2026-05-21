package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	QueryTimeoutDuration = time.Second * 5
	ErrDuplicateEmail    = errors.New("a user with that email already exist")
	ErrDuplicateUsername = errors.New("a user with that username already exist")
)

type Storage struct {
	Users interface {
		Create(context.Context, *sql.Tx, *User) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users: &UserStore{db},
	}
}
