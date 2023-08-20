package entity

import (
	"github.com/google/uuid"
)

type Author struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
}
