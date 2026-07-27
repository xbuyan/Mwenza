package ids

import "github.com/google/uuid"

type ID string

func New() ID {
	return ID(uuid.NewString())
}

func Parse(value string) ID {
	return ID(value)
}

func (id ID) String() string {
	return string(id)
}
