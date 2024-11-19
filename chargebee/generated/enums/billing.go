package enums

// THIS IS GENERATED CODE. DO NOT EDIT.

type BillingDayOfWeekEnum string

const (
        Monday BillingDayOfWeekEnum = "monday"
        Sunday BillingDayOfWeekEnum = "sunday"
        Tuesday BillingDayOfWeekEnum = "tuesday"
        Wednesday BillingDayOfWeekEnum = "wednesday"
)


type BillingDateModeEnum string

const (
        ManuallySet BillingDateModeEnum = "manually_set"
        UsingDefaults BillingDateModeEnum = "using_defaults"
)


type BillingDayOfWeekModeEnum BillingDateModeEnum