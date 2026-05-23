package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	QueryTimeoutDuration = time.Second * 5
	ErrDuplicateEmail    = errors.New("a user with that email already exist")
	ErrDuplicateUsername = errors.New("a user with that username already exist")
	ErrDuplicateName     = errors.New("an organization with that name already exist")
	ErrDuplicateOrgEmail = errors.New("an organization with that email already exist")
)

type Storage struct {
	Users interface {
		Create(context.Context, *sql.Tx, *User) error
	}
	Organizations interface {
		Create(context.Context, *sql.Tx, *Organization) error
		GetByID(context.Context, uuid.UUID) (*Organization, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users:         &UserStore{db},
		Organizations: &OrganizationStore{db},
	}
}
