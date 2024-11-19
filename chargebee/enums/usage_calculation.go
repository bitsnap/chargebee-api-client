package enums

type UsageCalculationEnum string

const (
	SumOfUsages UsageCalculationEnum = "sum_of_usages"
	LastUsage                        = "last_usage"
	MaxUsage                         = "max_usage"
)
