package enums_manual

type InitiatorTypeEnum string

const (
	CustomerTransaction InitiatorTypeEnum = "customer"
	MerchantTransaction                   = "merchant"
)
