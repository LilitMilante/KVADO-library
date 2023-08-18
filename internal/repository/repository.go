package repository

import (
	"context"
	"database/sql"

	"KVADO-library/internal/entity"

	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) BooksByAuthor(ctx context.Context, authorID uuid.UUID) ([]entity.Book, error) {
	q := `SELECT id, author_id, title FROM books WHERE author_id = ?`

	rows, err := r.db.QueryContext(ctx, q, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []entity.Book

	for rows.Next() {
		var book entity.Book

		err := rows.Scan(&book.ID, &book.AuthorID, &book.Title)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}
