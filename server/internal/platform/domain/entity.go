package domain

import "github.com/mwenza/mwenza/internal/platform/ids"

type Entity interface {
	ID() ids.ID
}
