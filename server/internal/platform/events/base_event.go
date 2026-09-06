package events

import (
	"time"

	"github.com/mwenza/mwenza/internal/platform/ids"
)

type BaseEvent struct {
	id         ids.ID
	occurredAt time.Time
}

func NewBaseEvent() BaseEvent {
	return BaseEvent{
		id:         ids.New(),
		occurredAt: time.Now().UTC(),
	}
}

func (e BaseEvent) EventID() ids.ID {
	return e.id
}

func (e BaseEvent) OccurredAt() time.Time {
	return e.occurredAt
}
