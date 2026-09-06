package sale

type Status string

const (
	StatusDraft     Status = "draft"
	StatusConfirmed Status = "confirmed"
	StatusCancelled Status = "cancelled"
	StatusCompleted Status = "completed"
)
