package enums

type CreditTypeEnum string

const (
	CreditLoyalty  CreditTypeEnum = "loyalty_credits"
	CreditReferral                = "referral_rewards"
	CreditGeneral                 = "general"
)
