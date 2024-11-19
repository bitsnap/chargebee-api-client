package enums_manual

type ScheduleTypeEnum string

const (
	FixedIntervals         ScheduleTypeEnum = "fixed_intervals"
	SpecificDatesIntervals                  = "specific_dates"
)
