package api

import (
	"context"

	"KVADO-library/gen/proto"
	"KVADO-library/internal/entity"

	"github.com/google/uuid"
)

type Service interface {
	BooksByAuthor(ctx context.Context, authorID uuid.UUID) ([]entity.Book, error)
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

func (h *Handler) BooksByAuthor(ctx context.Context, r *proto.BooksByAuthorRequest) (*proto.BooksByAuthorResponse, error) {
	authorID, err := uuid.Parse(r.AuthorId)
	if err != nil {
		return nil, err
	}

	books, err := h.srv.BooksByAuthor(ctx, authorID)
	if err != nil {
		return nil, err
	}

	return BooksByAuthorResponse(books), nil
}
