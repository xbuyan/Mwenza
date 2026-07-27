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
