package quantity

import "errors"

var ErrNegativeQuantity = errors.New("quantity cannot be negative")

type Quantity struct {
	value int64
}

func New(v int64) (Quantity, error) {
	if v < 0 {
		return Quantity{}, ErrNegativeQuantity
	}

	return Quantity{value: v}, nil
}

func (q Quantity) Value() int64 {
	return q.value
}
