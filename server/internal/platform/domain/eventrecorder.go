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
