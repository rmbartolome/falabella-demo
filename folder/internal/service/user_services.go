package service

import (
	"context"
	"folder/internal/repository"

	"github.com/go-kit/kit/log"
)

type UserServices interface {
	FindAll(ctx context.Context) ([]repository.User, error)
	GetByID(ctx context.Context, id string) (repository.User, error)
	Create(ctx context.Context, User repository.User) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) (repository.User, error)
}

type userService struct {
	repository repository.UserRepository
	logger     log.Logger
}

func NewUserService(repository repository.UserRepository, logger log.Logger) UserServices {
	return &userService{
		repository: repository,
		logger:     logger,
	}
}

func (s *userService) FindAll(ctx context.Context) ([]repository.User, error) {
	users, err := s.repository.FindAll(ctx)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return users, nil
}

func (s *userService) GetByID(ctx context.Context, id string) (repository.User, error) {
	users, err := s.repository.GetByID(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return users, err
	}
	s.logger.Log("getbyid:", "success")
	return users, nil
}

func (s *userService) Create(ctx context.Context, user repository.User) error {
	user.Prepare()
	err := user.Validate()
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}

	if err := s.repository.Create(ctx, user); err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("create:", "success")

	return nil
}

func (s *userService) Delete(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("delete:", "success")
	return nil
}
