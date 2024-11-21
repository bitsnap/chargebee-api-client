package enums_manual

type DurationTypeEnum string

const (
	DurationOneTime       DurationTypeEnum = "one_time"
	DurationForever                        = "forever"
	DurationLimitedPeriod                  = "limited_period"
)