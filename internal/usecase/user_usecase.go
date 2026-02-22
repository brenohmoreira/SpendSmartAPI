package usecase

import (
	"SpendSmartAPI/internal/domain"
	"SpendSmartAPI/internal/repository"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
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

	existUser, err := u.repository.FindByEmail(ctx, user.Email)

	if err != nil {
		return errors.New("Internal error while validating email existence")
	}

	if existUser != nil {
		return errors.New("User already exists")
	}

	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return errors.New("Internal encryption error")
	}

	user.Password = string(encryptPassword)

	return u.repository.Create(ctx, user)
}

func (u *UserUseCase) FindAll(ctx context.Context) ([]domain.User, error) {
	return u.repository.FindAll(ctx)
}

func (u *UserUseCase) FindById(ctx context.Context, id int) (*domain.User, error) {
	return u.repository.FindById(ctx, id)
}
