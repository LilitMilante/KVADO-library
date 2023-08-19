package api

import (
	"KVADO-library/gen/proto"
	"KVADO-library/internal/entity"

	"github.com/google/uuid"
)

func BooksByAuthorResponse(books []entity.Book) *proto.BooksByAuthorResponse {
	return &proto.BooksByAuthorResponse{
		Books: booksToAPI(books),
	}
}

func bookToAPI(book entity.Book) *proto.Book {
	return &proto.Book{
		Id:        book.ID.String(),
		Title:     book.Title,
		AuthorIds: authorIDsToAPI(book.AuthorIDs),
	}
}

func authorIDsToAPI(authorIDs []uuid.UUID) []string {
	strID := make([]string, 0, len(authorIDs))
	for _, v := range authorIDs {
		strID = append(strID, v.String())
	}

	return strID
}

func booksToAPI(books []entity.Book) []*proto.Book {
	apiBooks := make([]*proto.Book, 0, len(books))

	for _, v := range books {
		apiBooks = append(apiBooks, bookToAPI(v))
	}

	return apiBooks
}
