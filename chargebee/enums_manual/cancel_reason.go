package enums_manual

type CancelReasonEnum string

const (
	CancelNotPaid                CancelReasonEnum = "not_paid"
	CancelNoCard                                  = "no_card"
	CancelFraudReviewFailed                       = "fraud_review_failed"
	CancelNotCompliantEUCustomer                  = "non_compliant_eu_customer"
	CancelTaxCalculationFailed                    = "tax_calculation_failed"
	CancelCurrencyIncompatible                    = "currency_incompatible_with_gateway"
	CancelNonCompliantCustomer                    = "non_compliant_customer"
)
