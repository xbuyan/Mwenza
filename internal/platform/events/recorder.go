package events

type Recorder struct {
	events []Event
}

func (r *Recorder) Record(event Event) {
	r.events = append(r.events, event)
}

func (r *Recorder) Pull() []Event {
	pending := r.events
	r.events = nil
	return pending
}

func (r *Recorder) HasEvents() bool {
	return len(r.events) > 0
}
