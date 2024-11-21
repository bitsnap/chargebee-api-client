package enums_manual

type PiiClearedEnum string

const (
	PiiClearActive    PiiClearedEnum = "active"
	PiiClearScheduled                = "scheduled_for_clear"
	PiiCleared                       = "cleared"
)
