package service

import (
	"awasomeProject/internal/domain"
	"context"
)

type UsersRepository interface {
	Create(ctx context.Context, user domain.User) error
	GetByID(ctx context.Context, id int64) (domain.User, error)
	GetAll(ctx context.Context) ([]domain.User, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, inp domain.UpdateUserInput) error
}

type Users struct {
	repo UsersRepository
}

func NewUsers(repo UsersRepository) *Users {
	return &Users{
		repo: repo,
	}
}

func (b *Users) Create(ctx context.Context, user domain.User) error {
	return b.repo.Create(ctx, user)
}

func (b *Users) GetByID(ctx context.Context, id int64) (domain.User, error) {
	return b.repo.GetByID(ctx, id)
}

func (b *Users) GetAll(ctx context.Context) ([]domain.User, error) {
	return b.repo.GetAll(ctx)
}

func (b *Users) Delete(ctx context.Context, id int64) error {
	return b.repo.Delete(ctx, id)
}

func (b *Users) Update(ctx context.Context, id int64, inp domain.UpdateUserInput) error {
	return b.repo.Update(ctx, id, inp)
}
