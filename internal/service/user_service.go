package service

import (
	"context"
	"time"

	"go-users-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, name string, dob string) (int64, error) {
	parsedDob, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return 0, err
	}

	return s.repo.CreateUser(ctx, name, parsedDob)
}

func (s *UserService) GetUserByID(ctx context.Context, id int64) (int64, string, string, int, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return 0, "", "", 0, err
	}

	age := CalculateAge(user.Dob)

	return int64(user.ID), user.Name, user.Dob.Format("2006-01-02"), age, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int64, name, dob string) error {
	parsedDob, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return err
	}

	return s.repo.UpdateUser(ctx, id, name, parsedDob)
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context) ([]map[string]interface{}, error) {
	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, 0)

	for _, u := range users {
		result = append(result, map[string]interface{}{
			"id":   u.ID,
			"name": u.Name,
			"dob":  u.Dob.Format("2006-01-02"),
			"age":  CalculateAge(u.Dob),
		})
	}

	return result, nil
}
