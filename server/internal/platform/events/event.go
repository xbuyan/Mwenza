package events

import (
	"time"

	"github.com/mwenza/mwenza/internal/platform/ids"
)

type Event interface {
	EventID() ids.ID
	EventName() string
	OccurredAt() time.Time
}
