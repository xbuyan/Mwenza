package events

import "testing"

func TestNewBaseEvent(t *testing.T) {
	e := NewBaseEvent()

	if e.EventID().String() == "" {
		t.Fatal("expected event id")
	}

	if e.OccurredAt().IsZero() {
		t.Fatal("expected occurrence time")
	}
}
