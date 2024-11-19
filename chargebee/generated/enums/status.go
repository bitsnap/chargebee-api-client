package enums

// THIS IS GENERATED CODE. DO NOT EDIT.

type StatusEnum string

const (
	EntitlementFeatureStatusActive   StatusEnum = "active"
	EntitlementFeatureStatusArchived StatusEnum = "archived"
	EntitlementFeatureStatusDraft    StatusEnum = "draft"
)

const (
	BusinessEntityStatusActive   StatusEnum = "active"
	BusinessEntityStatusInactive StatusEnum = "inactive"
)

const (
	CreditNoteStatusAdjusted  StatusEnum = "adjusted"
	CreditNoteStatusRefundDue StatusEnum = "refund_due"
	CreditNoteStatusRefunded  StatusEnum = "refunded"
	CreditNoteStatusVoided    StatusEnum = "voided"
)

const (
	InvoiceStatusNotPaid    StatusEnum = "not_paid"
	InvoiceStatusPaid       StatusEnum = "paid"
	InvoiceStatusPaymentDue StatusEnum = "payment_due"
	InvoiceStatusPosted     StatusEnum = "posted"
)

const (
	AttachedItemStatusActive   StatusEnum = "active"
	AttachedItemStatusArchived StatusEnum = "archived"
	AttachedItemStatusDeleted  StatusEnum = "deleted"
)

const (
	CouponCodeStatusArchived    StatusEnum = "archived"
	CouponCodeStatusNotRedeemed StatusEnum = "not_redeemed"
	CouponCodeStatusRedeemed    StatusEnum = "redeemed"
)

const (
	CouponStatusActive   StatusEnum = "active"
	CouponStatusArchived StatusEnum = "archived"
	CouponStatusDeleted  StatusEnum = "deleted"
	CouponStatusExpired  StatusEnum = "expired"
)

const (
	ItemDifferentialPriceStatusActive  StatusEnum = "active"
	ItemDifferentialPriceStatusDeleted StatusEnum = "deleted"
)

const (
	ItemFamilyStatusActive  StatusEnum = "active"
	ItemFamilyStatusDeleted StatusEnum = "deleted"
)

const (
	ItemPriceStatusActive   StatusEnum = "active"
	ItemPriceStatusArchived StatusEnum = "archived"
	ItemPriceStatusDeleted  StatusEnum = "deleted"
)

const (
	ItemStatusActive   StatusEnum = "active"
	ItemStatusArchived StatusEnum = "archived"
	ItemStatusDeleted  StatusEnum = "deleted"
)

const (
	OrderStatusCancelled  StatusEnum = "cancelled"
	OrderStatusComplete   StatusEnum = "complete"
	OrderStatusNew        StatusEnum = "new"
	OrderStatusProcessing StatusEnum = "processing"
)

const (
	PaymentSourceStatusExpired  StatusEnum = "expired"
	PaymentSourceStatusExpiring StatusEnum = "expiring"
	PaymentSourceStatusInvalid  StatusEnum = "invalid"
	PaymentSourceStatusValid    StatusEnum = "valid"
)

const (
	PaymentVoucherStatusActive   StatusEnum = "active"
	PaymentVoucherStatusConsumed StatusEnum = "consumed"
	PaymentVoucherStatusExpired  StatusEnum = "expired"
	PaymentVoucherStatusFailure  StatusEnum = "failure"
)

const (
	TransactionStatusFailure    StatusEnum = "failure"
	TransactionStatusInProgress StatusEnum = "in_progress"
	TransactionStatusSuccess    StatusEnum = "success"
	TransactionStatusVoided     StatusEnum = "voided"
)

const (
	QuoteStatusAccepted StatusEnum = "accepted"
	QuoteStatusDeclined StatusEnum = "declined"
	QuoteStatusInvoiced StatusEnum = "invoiced"
	QuoteStatusOpen     StatusEnum = "open"
)

const (
	GiftStatusCancelled StatusEnum = "cancelled"
	GiftStatusClaimed   StatusEnum = "claimed"
	GiftStatusScheduled StatusEnum = "scheduled"
	GiftStatusUnclaimed StatusEnum = "unclaimed"
)

const (
	RampsStatusDraft     StatusEnum = "draft"
	RampsStatusFailed    StatusEnum = "failed"
	RampsStatusScheduled StatusEnum = "scheduled"
	RampsStatusSucceeded StatusEnum = "succeeded"
)

const (
	SubscriptionStatusActive      StatusEnum = "active"
	SubscriptionStatusFuture      StatusEnum = "future"
	SubscriptionStatusInTrial     StatusEnum = "in_trial"
	SubscriptionStatusNonRenewing StatusEnum = "non_renewing"
)
