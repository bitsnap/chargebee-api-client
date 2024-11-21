package enums

// THIS IS GENERATED CODE. DO NOT EDIT.

type BillingDayOfWeekEnum string

const (
        Monday BillingDayOfWeekEnum = "monday"
        Sunday = "sunday"
        Tuesday = "tuesday"
        Wednesday = "wednesday"
)

type BillingDateModeEnum string

const (
        ManuallySet BillingDateModeEnum = "manually_set"
        UsingDefaults = "using_defaults"
)

type BillingDayOfWeekModeEnum BillingDateModeEnum