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
