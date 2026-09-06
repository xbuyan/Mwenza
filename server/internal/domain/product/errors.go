package product

import "errors"

var (
	ErrEmptyID     = errors.New("product id cannot be empty")
	ErrEmptyName   = errors.New("product name cannot be empty")
	ErrEmptySKU    = errors.New("product SKU cannot be empty")
	ErrInvalidUnit = errors.New("invalid product unit")
)
