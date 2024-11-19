package enums_manual

type TaxabilityEnum string

const (
	CustomerTaxable   TaxabilityEnum = "taxable"
	CustomerTaxExempt                = "exempt"
)
