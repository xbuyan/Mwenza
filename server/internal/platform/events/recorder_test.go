package events

import "testing"

type testEvent struct {
	BaseEvent
	name string
}

func newTestEvent(name string) testEvent {
	return testEvent{
		BaseEvent: NewBaseEvent(),
		name:      name,
	}
}

func (e testEvent) EventName() string {
	return e.name
}

func TestRecorder(t *testing.T) {
	var recorder Recorder

	if recorder.HasEvents() {
		t.Fatal("expected empty recorder")
	}

	recorder.Record(newTestEvent("test.created"))
	recorder.Record(newTestEvent("test.updated"))

	if !recorder.HasEvents() {
		t.Fatal("expected recorder to contain events")
	}

	events := recorder.Pull()

	if len(events) != 2 {
		t.Fatalf("expected 2 events, got %d", len(events))
	}

	if recorder.HasEvents() {
		t.Fatal("expected recorder to be empty after pull")
	}
}

func TestPullEmptyRecorder(t *testing.T) {
	var recorder Recorder

	events := recorder.Pull()

	if len(events) != 0 {
		t.Fatalf("expected no events, got %d", len(events))
	}
}
