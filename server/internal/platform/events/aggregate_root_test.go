package events

import "testing"

type testAggregate struct {
	AggregateRoot
}

func TestAggregateRootRecordsEvents(t *testing.T) {
	var agg testAggregate

	if agg.HasEvents() {
		t.Fatal("expected no events")
	}

	agg.Record(newTestEvent("aggregate.created"))

	if !agg.HasEvents() {
		t.Fatal("expected recorded event")
	}

	events := agg.Pull()

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}

	if agg.HasEvents() {
		t.Fatal("expected no remaining events")
	}
}
