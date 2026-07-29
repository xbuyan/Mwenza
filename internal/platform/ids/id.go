package ids

import (
	"errors"

	"github.com/google/uuid"
)

var ErrInvalidID = errors.New("invalid id")

type ID string

func New() ID {
	return ID(uuid.NewString())
}

func Parse(value string) (ID, error) {
	if value == "" {
		return "", ErrInvalidID
	}

	if _, err := uuid.Parse(value); err != nil {
		return "", ErrInvalidID
	}

	return ID(value), nil
}

func MustParse(value string) ID {
	id, err := Parse(value)
	if err != nil {
		panic(err)
	}

	return id
}

func (id ID) String() string {
	return string(id)
}

func (id ID) IsZero() bool {
	return id == ""
}
