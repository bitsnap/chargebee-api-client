package enums_manual

type DiscountTypeEnum string

const (
	DiscountFixed      DiscountTypeEnum = "fixed_amount"
	DiscountPercentage                  = "percentage"
)
