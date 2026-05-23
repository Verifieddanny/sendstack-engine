package store

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Organization struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type OrganizationStore struct {
	db *sql.DB
}

func (s *OrganizationStore) Create(ctx context.Context, tx *sql.Tx, org *Organization) error {
	query := `
	INSERT INTO organization (name, email) 
	VALUES ($1, $2) RETURNING id, created_at, updated_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := tx.QueryRowContext(
		ctx,
		query,
		org.Name,
		org.Email,
	).Scan(
		&org.ID,
		&org.CreatedAt,
		&org.UpdatedAt,
	)
	if err != nil {
		switch err.Error() {
		case `pq: duplicate key value violates unique constraint "organization_email_key`:
			return ErrDuplicateOrgEmail

		case `pq: duplicate key value violates unique constraint "organization_name_key`:
			return ErrDuplicateName
		default:
			return err
		}
	}
	return nil
}

func (s *OrganizationStore) GetByID(ctx context.Context, id uuid.UUID) (*Organization, error) {
	query := `
	SELECT id, name, email, created_at, updated_at
	FROM organization
	WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var org Organization

	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&org.ID,
		&org.Name,
		&org.Email,
		&org.CreatedAt,
		&org.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &org, nil
}
