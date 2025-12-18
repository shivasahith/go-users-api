package repository

import (
	"context"
	"database/sql"
	"time"

	db "go-users-api/db/sqlc"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(dbConn *sql.DB) *UserRepository {
	return &UserRepository{
		queries: db.New(dbConn),
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, name string, dob time.Time) (int64, error) {

	result, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (db.User, error) {
	return r.queries.GetUserByID(ctx, int32(id))
}
func (r *UserRepository) UpdateUser(ctx context.Context, id int64, name string, dob time.Time) error {
	return r.queries.UpdateUser(ctx, db.UpdateUserParams{
		ID:   int32(id),
		Name: name,
		Dob:  dob,
	})
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int64) error {
	return r.queries.DeleteUser(ctx, int32(id))
}

func (r *UserRepository) ListUsers(ctx context.Context) ([]db.User, error) {
	return r.queries.ListUsers(ctx)
}
