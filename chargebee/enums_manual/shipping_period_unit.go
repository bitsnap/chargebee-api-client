package enums_manual

type ShippingPeriodUnitEnum string

const (
	DayShippingPeriod   ShippingPeriodUnitEnum = "day"
	WeekShippingPeriod                         = "week"
	MonthShippingPeriod                        = "month"
	YearShippingPeriod                         = "year"
)
