package sale

import "errors"

var ErrEmptySaleID = errors.New("sale id cannot be empty")

type Sale struct {
	id     string
	status Status
}

func New(id string) (*Sale, error) {
	if id == "" {
		return nil, ErrEmptySaleID
	}

	return &Sale{
		id:     id,
		status: StatusDraft,
	}, nil
}
