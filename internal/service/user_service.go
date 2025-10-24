package service

import (
	"context"
	"database/sql"
	"time"
	db "user-management-api/db/sqlc"
	"user-management-api/internal/models"
	"user-management-api/internal/repository"
)

type UserService struct {
	repo *repository.Database
}

func NewUserService(repo *repository.Database) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error) {
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return models.UserResponse{}, err
	}

	result, err := s.repo.Queries.CreateUser(ctx, db.CreateUserParams{
		Name: req.Name,
		Dob:  dob,
	})
	if err != nil {
		return models.UserResponse{}, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return models.UserResponse{}, err
	}

	user, err := s.repo.Queries.GetUser(ctx, int32(lastID))
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserFromDB(user.ID, user.Name, user.Dob)
}

func (s *UserService) GetUser(ctx context.Context, id int32) (models.UserResponse, error) {
	user, err := s.repo.Queries.GetUser(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.UserResponse{}, err
		}
		return models.UserResponse{}, err
	}

	return models.UserFromDB(user.ID, user.Name, user.Dob)
}

func (s *UserService) ListUsers(ctx context.Context, limit, offset int32) ([]models.UserResponse, error) {
	users, err := s.repo.Queries.ListUsers(ctx, db.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	var response []models.UserResponse
	for _, user := range users {
		userResp, err := models.UserFromDB(user.ID, user.Name, user.Dob)
		if err != nil {
			return nil, err
		}
		response = append(response, userResp)
	}
	return response, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error) {
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return models.UserResponse{}, err
	}

	err = s.repo.Queries.UpdateUser(ctx, db.UpdateUserParams{
		ID:   id,
		Name: req.Name,
		Dob:  dob,
	})
	if err != nil {
		return models.UserResponse{}, err
	}

	user, err := s.repo.Queries.GetUser(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserFromDB(user.ID, user.Name, user.Dob)
}

// Add this missing method
func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.repo.Queries.DeleteUser(ctx, id)
}
