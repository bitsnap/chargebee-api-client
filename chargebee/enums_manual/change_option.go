package enums_manual

type ChangeOptionEnum string

const (
	ChangeEndOfTerm    ChangeOptionEnum = "end_of_term"
	ChangeSpecificDate                  = "specific_date"
	ChangeImmediately                   = "immediately"
)
