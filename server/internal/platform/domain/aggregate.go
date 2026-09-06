package domain

import "github.com/mwenza/mwenza/internal/platform/events"

type AggregateRoot interface {
	Entity

	PendingEvents() []events.Event

	ClearEvents()
}
