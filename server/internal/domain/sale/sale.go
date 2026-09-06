package sale

import (
	"errors"

	"github.com/mwenza/mwenza/internal/platform/ids"
)

var ErrEmptySaleID = errors.New("sale id cannot be empty")

type Sale struct {
	id     ids.ID
	status Status
}

func New(id ids.ID) (*Sale, error) {
	if id.IsZero() {
		return nil, ErrEmptySaleID
	}

	return &Sale{
		id:     id,
		status: StatusDraft,
	}, nil
}
