package enums

type BillingDayOfWeekModeEnum string

const (
	BillingUsingDefaults BillingDayOfWeekModeEnum = "using_defaults"
	BillingManuallySet                            = "manually_set"
)

type BillingDateModeEnum string

const (
	BillingDateUsingDefaults BillingDayOfWeekModeEnum = "using_defaults"
	BillingDateManuallySet                            = "manually_set"
)

type BillingDayOfWeekEnum string

const (
	BillingOnMonday    BillingDayOfWeekEnum = "monday"
	BillingOnTuesday                        = "tuesday"
	BillingOnWednesday                      = "wednesday"
	BillingOnThursday                       = "thursday"
	BillingOnFriday                         = "friday"
	BillingOnSaturday                       = "saturday"
	BillingOnSunday                         = "sunday"
)
