package repository

import (
	"SpendSmartAPI/internal/domain"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindAll(ctx context.Context) ([]domain.User, error)
	FindById(ctx context.Context, id int) (*domain.User, error)
}
