package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"KVADO-library/internal/entity"

	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) BooksByAuthorID(ctx context.Context, authorID uuid.UUID) ([]entity.Book, error) {
	q := `SELECT id, title FROM books
    	  JOIN authors_books ON books.id = authors_books.book_id
    	  WHERE authors_books.author_id = ?`

	rows, err := r.db.QueryContext(ctx, q, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []entity.Book

	for rows.Next() {
		var book entity.Book

		err := rows.Scan(&book.ID, &book.Title)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	bookAuthors, err := r.booksAuthorIDs(ctx, books)
	if err != nil {
		return nil, err
	}

	for i, v := range books {
		books[i].AuthorIDs = bookAuthors[v.ID]
	}

	return books, nil
}

func (r *Repository) booksAuthorIDs(ctx context.Context, books []entity.Book) (map[uuid.UUID][]uuid.UUID, error) {
	// Prepare placeholders for the IN clause
	placeholders := make([]string, len(books))
	for i := range books {
		placeholders[i] = "?"
	}

	// Construct the query with the placeholders
	q := fmt.Sprintf(`
		SELECT author_id, book_id
		FROM authors_books
		WHERE book_id IN (%s)
	`, strings.Join(placeholders, ","))

	args := make([]any, len(books))
	for i, v := range books {
		args[i] = v.ID
	}

	// Execute the query
	rows, err := r.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a map to store the results
	bookAuthors := make(map[uuid.UUID][]uuid.UUID)

	// Iterate over the result rows
	for rows.Next() {
		var authorID uuid.UUID
		var bookID uuid.UUID
		if err := rows.Scan(&authorID, &bookID); err != nil {
			return nil, err
		}
		bookAuthors[bookID] = append(bookAuthors[bookID], authorID)
	}

	return bookAuthors, nil
}

func (r *Repository) AuthorsByBookID(ctx context.Context, bookID uuid.UUID) ([]entity.Author, error) {
	q := `SELECT id, first_name, last_name FROM authors
    	  JOIN authors_books ON authors.id = authors_books.author_id
    	  WHERE authors_books.book_id = ?`

	rows, err := r.db.QueryContext(ctx, q, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []entity.Author

	for rows.Next() {
		var author entity.Author

		err := rows.Scan(&author.ID, &author.FirstName, &author.LastName)
		if err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	return authors, nil
}
