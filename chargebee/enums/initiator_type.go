package enums

type InitiatorTypeEnum string

const (
	CustomerTransaction InitiatorTypeEnum = "customer"
	MerchantTransaction                   = "merchant"
)
