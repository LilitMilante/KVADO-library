package api

import (
	"KVADO-library/gen/proto"
	"KVADO-library/internal/entity"
)

func BooksByAuthorResponse(books []entity.Book) *proto.BooksByAuthorResponse {
	return &proto.BooksByAuthorResponse{
		Books: booksToAPI(books),
	}
}

func bookToAPI(book entity.Book) *proto.Book {
	return &proto.Book{
		Id:       book.ID.String(),
		AuthorId: book.AuthorID.String(),
		Title:    book.Title,
	}
}

func booksToAPI(books []entity.Book) []*proto.Book {
	apiBooks := make([]*proto.Book, 0, len(books))

	for _, v := range books {
		apiBooks = append(apiBooks, bookToAPI(v))
	}

	return apiBooks
}
