#!/usr/bin/env bash
set -euo pipefail

echo "======================================"
echo "Mwenza Sprint 002 - Domain Foundation"
echo "======================================"

mkdir -p internal/platform/domain

########################################
# eventrecorder.go
########################################

cat > internal/platform/domain/eventrecorder.go <<'EOGO'
package domain

import "github.com/mwenza/mwenza/internal/platform/events"

type EventRecorder struct {
	events []events.Event
}

func (r *EventRecorder) Record(event events.Event) {
	r.events = append(r.events, event)
}

func (r *EventRecorder) PendingEvents() []events.Event {
	out := make([]events.Event, len(r.events))
	copy(out, r.events)
	return out
}

func (r *EventRecorder) ClearEvents() {
	r.events = nil
}
EOGO

########################################
# baseaggregate.go
########################################

cat > internal/platform/domain/baseaggregate.go <<'EOGO'
package domain

import (
	"github.com/mwenza/mwenza/internal/platform/events"
	"github.com/mwenza/mwenza/internal/platform/ids"
)

type BaseAggregate struct {
	id ids.ID
	EventRecorder
}

func NewBaseAggregate(id ids.ID) BaseAggregate {
	return BaseAggregate{
		id: id,
	}
}

func (a BaseAggregate) ID() ids.ID {
	return a.id
}

func (a *BaseAggregate) Record(event events.Event) {
	a.EventRecorder.Record(event)
}
EOGO

########################################
# README.md
########################################

cat > internal/platform/domain/README.md <<'EOGO'
# Domain Package

This package contains the fundamental building blocks used by every
bounded context in Mwenza.

Rules:

- Business logic lives in the domain.
- Infrastructure depends on the domain.
- The domain never depends on infrastructure.
- Aggregates emit domain events.
- Aggregates enforce invariants.
EOGO

########################################
# Tests
########################################

cat > internal/platform/domain/eventrecorder_test.go <<'EOGO'
package domain

import (
	"testing"
	"time"

	"github.com/mwenza/mwenza/internal/platform/ids"
)

type fakeEvent struct {
	id ids.ID
}

func (f fakeEvent) EventID() ids.ID {
	return f.id
}

func (f fakeEvent) EventName() string {
	return "fake"
}

func (f fakeEvent) OccurredAt() time.Time {
	return time.Now()
}

func TestEventRecorder(t *testing.T) {
	var recorder EventRecorder

	recorder.Record(fakeEvent{
		id: ids.New(),
	})

	if len(recorder.PendingEvents()) != 1 {
		t.Fatalf("expected 1 event")
	}

	recorder.ClearEvents()

	if len(recorder.PendingEvents()) != 0 {
		t.Fatalf("expected no events")
	}
}
EOGO

gofmt -w internal

echo
echo "Sprint 002 completed successfully."
