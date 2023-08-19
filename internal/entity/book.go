package entity

import (
	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID
	Title     string
	AuthorIDs []uuid.UUID
}
