#!/usr/bin/env bash
set -euo pipefail

echo "======================================"
echo "Mwenza Sprint 001 - Platform Kernel"
echo "======================================"

mkdir -p \
internal/platform/domain \
internal/platform/events \
internal/platform/errors \
internal/platform/ids \
internal/platform/clock

########################################
# ids
########################################

cat > internal/platform/ids/id.go <<'EOGO'
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
EOGO

########################################
# clock
########################################

cat > internal/platform/clock/clock.go <<'EOGO'
package clock

import "time"

type Clock interface {
	Now() time.Time
}

type SystemClock struct{}

func (SystemClock) Now() time.Time {
	return time.Now().UTC()
}
EOGO

########################################
# errors
########################################

cat > internal/platform/errors/errors.go <<'EOGO'
package errors

import stderrors "errors"

var (
	ErrNotFound      = stderrors.New("not found")
	ErrUnauthorized  = stderrors.New("unauthorized")
	ErrConflict      = stderrors.New("conflict")
	ErrValidation    = stderrors.New("validation failed")
)
EOGO

########################################
# events
########################################

cat > internal/platform/events/event.go <<'EOGO'
package events

import (
	"time"

	"github.com/mwenza/mwenza/internal/platform/ids"
)

type Event interface {
	EventID() ids.ID
	EventName() string
	OccurredAt() time.Time
}
EOGO

cat > internal/platform/events/publisher.go <<'EOGO'
package events

import "context"

type Publisher interface {
	Publish(ctx context.Context, events ...Event) error
}
EOGO

########################################
# domain
########################################

cat > internal/platform/domain/entity.go <<'EOGO'
package domain

import "github.com/mwenza/mwenza/internal/platform/ids"

type Entity interface {
	ID() ids.ID
}
EOGO

cat > internal/platform/domain/aggregate.go <<'EOGO'
package domain

import "github.com/mwenza/mwenza/internal/platform/events"

type AggregateRoot interface {
	Entity

	PendingEvents() []events.Event

	ClearEvents()
}
EOGO

cat > internal/platform/domain/value_object.go <<'EOGO'
package domain

type ValueObject interface {
	isValueObject()
}
EOGO

cat > internal/platform/domain/repository.go <<'EOGO'
package domain

import "context"

type Repository[T any] interface {
	Save(ctx context.Context, aggregate T) error
}
EOGO

########################################
# dependencies
########################################

go get github.com/google/uuid

########################################
# formatting
########################################

gofmt -w internal

echo
echo "Sprint 001 completed successfully."
