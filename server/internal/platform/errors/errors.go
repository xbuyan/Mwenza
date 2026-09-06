package errors

import stderrors "errors"

var (
	ErrNotFound     = stderrors.New("not found")
	ErrUnauthorized = stderrors.New("unauthorized")
	ErrConflict     = stderrors.New("conflict")
	ErrValidation   = stderrors.New("validation failed")
)
