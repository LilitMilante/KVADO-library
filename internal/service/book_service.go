package service

import (
	"context"

	"KVADO-library/internal/entity"

	"github.com/google/uuid"
)

type Repository interface {
	BooksByAuthor(ctx context.Context, authorID uuid.UUID) ([]entity.Book, error)
}

type BookService struct {
	repo Repository
}

func NewBookService(repo Repository) *BookService {
	return &BookService{repo: repo}
}

func (b *BookService) BooksByAuthor(ctx context.Context, authorID uuid.UUID) ([]entity.Book, error) {
	return b.repo.BooksByAuthor(ctx, authorID)
}
