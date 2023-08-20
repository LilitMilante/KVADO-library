package service

import (
	"context"

	"KVADO-library/internal/entity"

	"github.com/google/uuid"
)

type Repository interface {
	BooksByAuthorID(ctx context.Context, authorID uuid.UUID) ([]entity.Book, error)
	AuthorsByBookID(ctx context.Context, bookID uuid.UUID) ([]entity.Author, error)
}

type Service struct {
	repo Repository
}

// create new service

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// get all books by author ID

func (s *Service) BooksByAuthorID(ctx context.Context, authorID uuid.UUID) ([]entity.Book, error) {
	return s.repo.BooksByAuthorID(ctx, authorID)
}

// get all authors by book ID

func (s *Service) AuthorsByBookID(ctx context.Context, bookID uuid.UUID) ([]entity.Author, error) {
	return s.repo.AuthorsByBookID(ctx, bookID)
}
