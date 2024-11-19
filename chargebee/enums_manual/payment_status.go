package enums_manual

type PaymentStatusEnum string

const (
	PaymentPaid    PaymentStatusEnum = "paid"
	PaymentNotPaid                   = "not_paid"
)
