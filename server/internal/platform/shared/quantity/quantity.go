package quantity

import "errors"

var (
	ErrNegativeQuantity     = errors.New("quantity cannot be negative")
	ErrInsufficientQuantity = errors.New("insufficient quantity")
)

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

func (q Quantity) Add(other Quantity) Quantity {
	return Quantity{
		value: q.value + other.value,
	}
}

func (q Quantity) Subtract(other Quantity) (Quantity, error) {
	if other.value > q.value {
		return Quantity{}, ErrInsufficientQuantity
	}

	return Quantity{
		value: q.value - other.value,
	}, nil
}

func (q Quantity) IsZero() bool {
	return q.value == 0
}
