package enums_manual

type FraudFlagEnum string

const (
	FraudSafe       FraudFlagEnum = "safe"
	FraudSuspicious               = "suspicious"
	FraudFraudulent               = "fraudulent"
)
