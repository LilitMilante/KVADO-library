package entity

import (
	"github.com/google/uuid"
)

type Book struct {
	ID       uuid.UUID
	AuthorID uuid.UUID
	Title    string
}
