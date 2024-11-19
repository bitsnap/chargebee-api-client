package enums_manual

type PriceTypeEnum string

const (
	TaxInclusivePrice PriceTypeEnum = "tax_inclusive"
	TaxExclusivePrice               = "tax_exclusive"
)
