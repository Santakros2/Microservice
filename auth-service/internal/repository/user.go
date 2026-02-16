package repository

import (
	"auth-service/internal/domain"
	"context"
	"database/sql"
)

type UserRepo struct {
	DB *sql.DB
}

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByID(ctx context.Context, id string) (*domain.User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

func NewRepo(db *sql.DB) UserRepository {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Create(ctx context.Context, user *domain.User) error {

	query := `
	INSERT INTO auth_users (id, email, role, password_hash, created_at)
	VALUES (?, ?, ?, ?, ?)`

	_, err := r.DB.ExecContext(
		ctx,
		query,
		user.ID,
		user.Email,
		user.Role,
		user.PasswordHash,
		user.CreatedAt,
	)

	return err
}
func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {

	query := `
	SELECT id, email, role, password_hash, created_at
	FROM auth_users
	WHERE email = ?`

	row := r.DB.QueryRowContext(ctx, query, email)

	var u domain.User
	err := row.Scan(
		&u.ID,
		&u.Email,
		&u.Role,
		&u.PasswordHash,
		&u.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserRepo) FindByID(ctx context.Context, id string) (*domain.User, error) {

	query := `
	SELECT id, email, role, password_hash, created_at
	FROM auth_users
	WHERE id = ?`

	row := r.DB.QueryRowContext(ctx, query, id)

	var u domain.User
	err := row.Scan(
		&u.ID,
		&u.Email,
		&u.Role,
		&u.PasswordHash,
		&u.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserRepo) ExistsByEmail(ctx context.Context, email string) (bool, error) {

	query := `SELECT EXISTS(SELECT 1 FROM auth_users WHERE email = ?)`

	var exists bool
	err := r.DB.QueryRowContext(ctx, query, email).Scan(&exists)

	return exists, err
}
