package enums_manual

type AuthorizationReasonEnum string

const (
	BlockingFunds AuthorizationReasonEnum = "blocking_funds"
	Verification                          = "verification"
)
