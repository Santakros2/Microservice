package repository

import (
	"context"
	"database/sql"
	"profile-service/internal/domain"
)

type ProfileRepo struct {
	DB *sql.DB
}

func (r *ProfileRepo) FindByEmail(
	ctx context.Context,
	email string,
) (*domain.Profile, error) {

	query := `
	SELECT id, email, name, bio, created_at
	FROM profiles
	WHERE email = ?`

	row := r.DB.QueryRowContext(ctx, query, email)

	var p domain.Profile
	err := row.Scan(
		&p.ID,
		&p.Email,
		&p.Name,
		&p.Bio,
		&p.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *ProfileRepo) Create(ctx context.Context, user *domain.Profile) error {
	query := `
	INSERT INTO profiles (id, email, name, bio, created_at)
	VALUES (?, ?, ?, ?, ?)`

	_, err := r.DB.ExecContext(
		ctx,
		query,
		user.ID,
		user.Email,
		user.Name,
		user.Bio,
		user.CreatedAt,
	)

	return err
}
