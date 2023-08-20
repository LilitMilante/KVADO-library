package api

import (
	"KVADO-library/gen/proto"
	"KVADO-library/internal/entity"

	"github.com/google/uuid"
)

// Mappers for books

func BooksByAuthorIDResponse(books []entity.Book) *proto.BooksByAuthorResponse {
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

// Mappers for authors

func AuthorsByBookIDResponse(authors []entity.Author) *proto.AuthorsByBookResponse {
	return &proto.AuthorsByBookResponse{
		Authors: authorsToAPI(authors),
	}
}

func authorToAPI(author entity.Author) *proto.Author {
	return &proto.Author{
		Id:        author.ID.String(),
		FirstName: author.FirstName,
		LastName:  author.LastName,
	}
}

func authorsToAPI(authors []entity.Author) []*proto.Author {
	apiAuthors := make([]*proto.Author, 0, len(authors))

	for _, v := range authors {
		apiAuthors = append(apiAuthors, authorToAPI(v))
	}

	return apiAuthors
}
