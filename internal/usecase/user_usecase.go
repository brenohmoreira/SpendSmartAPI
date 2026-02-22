package usecase

import (
	"SpendSmartAPI/internal/domain"
	"SpendSmartAPI/internal/repository"
	"context"
	"errors"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repository repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		repository: repository,
	}
}

func (u *UserUseCase) Create(ctx context.Context, user *domain.User) error {
	if user.Name == "" {
		return errors.New("Name is required")
	}
	if user.Email == "" {
		return errors.New("Email is required")
	}
	if user.Phone == "" {
		return errors.New("Phone is required")
	}
	if user.Password == "" {
		return errors.New("Password is required")
	}

	// Criptogra senha
	// Verifica se existe

	return u.repository.Create(ctx, user)
}

func (u *UserUseCase) FindAll(ctx context.Context) ([]domain.User, error) {
	return u.repository.FindAll(ctx)
}

func (u *UserUseCase) FindById(ctx context.Context, id int) (*domain.User, error) {
	return u.repository.FindById(ctx, id)
}
