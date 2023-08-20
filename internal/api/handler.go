package api

import (
	"context"

	"KVADO-library/gen/proto"
	"KVADO-library/internal/entity"

	"github.com/google/uuid"
)

type Service interface {
	BooksByAuthorID(ctx context.Context, authorID uuid.UUID) ([]entity.Book, error)
	AuthorsByBookID(ctx context.Context, bookID uuid.UUID) ([]entity.Author, error)
}

type Handler struct {
	proto.UnimplementedLibraryServer

	srv Service
}

func NewHandler(srv Service) *Handler {
	return &Handler{
		srv: srv,
	}
}

// Get list books by author ID

func (h *Handler) BooksByAuthorID(ctx context.Context, r *proto.BooksByAuthorRequest) (*proto.BooksByAuthorResponse, error) {
	authorID, err := uuid.Parse(r.AuthorId) // parse string ID to UUID
	if err != nil {
		return nil, err
	}

	books, err := h.srv.BooksByAuthorID(ctx, authorID)
	if err != nil {
		return nil, err
	}

	return BooksByAuthorIDResponse(books), nil
}

// Get list authors by book ID

func (h *Handler) AuthorsByBookID(ctx context.Context, r *proto.AuthorsByBookRequest) (*proto.AuthorsByBookResponse, error) {
	bookID, err := uuid.Parse(r.BookId) // parse string ID to UUID
	if err != nil {
		return nil, err
	}

	authors, err := h.srv.AuthorsByBookID(ctx, bookID)
	if err != nil {
		return nil, err
	}

	return AuthorsByBookIDResponse(authors), nil
}
