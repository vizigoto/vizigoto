package user

import "github.com/vizigoto/vizigoto/pkg/uuid"

type ID string

func NewID() ID {
	id := uuid.New()
	return ID(id)
}
