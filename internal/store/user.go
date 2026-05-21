package store

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  password  `json:"-"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type password struct {
	text *string
	hash []byte
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, tx *sql.Tx, user *User) error {
	query := `
	INSERT INTO users (email, password_hash, name) 
	VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()


	err := tx.QueryRowContext(
		ctx,
		query,
		user.Email,
		user.Password.hash,
		user.Name,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		switch err.Error() {
		case `pq: duplicate key value violates unique constraint "users_email_key`:
			return ErrDuplicateEmail

		case `pq: duplicate key value violates unique constraint "users_username_key`:
			return ErrDuplicateUsername
		default:
			return err
		}
	}
	return nil
}
