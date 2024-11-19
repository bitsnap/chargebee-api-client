package enums

import "slices"

type StatusEnum string

const (
	BusinessEntityActiveStatus   StatusEnum = "active"
	BusinessEntityInactiveStatus            = "inactive"

	EntitlementFeatureActiveStatus   = "active"
	EntitlementFeatureArchivedStatus = "archived"
	EntitlementFeatureDraftStatus    = "draft"

	CreditNoteAdjustedStatus  = "adjusted"
	CreditNoteRefundedStatus  = "refunded"
	CreditNoteRefundDueStatus = "refund_due"
	CreditNoteVoided          = "voided"

	InvoicePaidStatus       = "paid"
	InvoicePostedStatus     = "posted"
	InvoicePaymentDueStatus = "payment_due"
	InvoiceNotPaidStatus    = "not_paid"
	InvoiceVoidedStatus     = "voided"
	InvoicePendingStatus    = "pending"

	AttachedItemActiveStatus   = "active"
	AttachedItemArchivedStatus = "archived"
	AttachedItemDeletedStatus  = "deleted"

	CouponCodeNotRedeemedStatus = "not_redeemed"
	CouponCodeRedeemedStatus    = "redeemed"
	CouponCodeArchivedStatus    = "archived"

	CouponActiveStatus   = "active"
	CouponExpiredStatus  = "expired"
	CouponArchivedStatus = "archived"
	CouponDeletedStatus  = "deleted"

	ItemDifferentialPriceActive  = "active"
	ItemDifferentialPriceDeleted = "deleted"

	ItemFamilyActiveStatus  = "active"
	ItemFamilyDeletedStatus = "deleted"

	ItemPriceActiveStatus   = "active"
	ItemPriceArchivedStatus = "archived"
	ItemPriceDeletedStatus  = "deleted"

	ItemActiveStatus   = "active"
	ItemArchivedStatus = "archived"
	ItemDeletedStatus  = "deleted"

	OrderNewStatus                = "new"
	OrderProcessingStatus         = "processing"
	OrderCompleteStatus           = "complete"
	OrderCancelledStatus          = "canceled"
	OrderVoidedStatus             = "voided"
	OrderQueuedStatus             = "queued"
	OrderAwaitingShipmentStatus   = "awaiting_shipment"
	OrderOnHoldStatus             = "on_hold"
	OrderDeliveredStatus          = "delivered"
	OrderShippedStatus            = "shipped"
	OrderPartiallyDeliveredStatus = "partially_delivered"
	OrderReturnedStatus           = "returned"

	PaymentSourceValidStatus               = "valid"
	PaymentSourceExpiringStatus            = "expiring"
	PaymentSourceExpiredStatus             = "expired"
	PaymentSourceInvalidStatus             = "invalid"
	PaymentSourcePendingVerificationStatus = "pending_verification"

	PaymentVoucherActiveStatus   = "active"
	PaymentVoucherConsumedStatus = "consumed"
	PaymentVoucherExpiredStatus  = "expired"
	PaymentVoucherFailureStatus  = "failure"

	TransactionInProgressStatus     = "in_progress"
	TransactionSuccessStatus        = "success"
	TransactionVoidedStatus         = "voided"
	TransactionFailureStatus        = "failure"
	TransactionTimeoutStatus        = "timeout"
	TransactionNeedsAttentionStatus = "needs_attention"

	QuoteOpenStatus     = "open"
	QuoteAcceptedStatus = "accepted"
	QuoteDeclinedStatus = "declined"
	QuoteInvoicedStatus = "invoiced"
	QuoteClosedStatus   = "closed"

	GiftScheduledStatus = "scheduled"
	GiftUnclaimedStatus = "unclaimed"
	GiftClaimedStatus   = "claimed"
	GiftCancelledStatus = "canceled"
	GiftExpiredStatus   = "expired"

	RampScheduledStatus = "scheduled"
	RampSucceededStatus = "succeeded"
	RampFailedStatus    = "failed"
	RampDraftStatus     = "draft"

	SubscriptionFutureStatus      = "future"
	SubscriptionInTrialStatus     = "in_trial"
	SubscriptionActiveStatus      = "active"
	SubscriptionNonRenewingStatus = "non_renewing"
	SubscriptionPausedStatus      = "paused"
	SubscriptionCancelledStatus   = "canceled"
	SubscriptionTransferredStatus = "transferred"
)

func KnownBusinessEntityStatus(name string) bool {
	return slices.Contains([]StatusEnum{
		BusinessEntityActiveStatus,
		BusinessEntityInactiveStatus,
	}, StatusEnum(name))
}

func KnownEntitlementFeatureStatus(name string) bool {
	return slices.Contains([]StatusEnum{
		EntitlementFeatureActiveStatus,
		EntitlementFeatureArchivedStatus,
		EntitlementFeatureDraftStatus,
	}, StatusEnum(name))
}

func KnownCreditNoteStatus(name string) bool {
	return slices.Contains([]StatusEnum{
		CreditNoteAdjustedStatus,
		CreditNoteRefundedStatus,
		CreditNoteRefundDueStatus,
		CreditNoteVoided,
	}, StatusEnum(name))
}

func KnownInvoiceStatus(name string) bool {
	return slices.Contains([]StatusEnum{
		InvoicePaidStatus,
		InvoicePostedStatus,
		InvoicePaymentDueStatus,
		InvoiceNotPaidStatus,
		InvoiceVoidedStatus,
		InvoicePendingStatus,
	}, StatusEnum(name))
}

func KnownAttachedItemStatus(name string) bool {
	return slices.Contains([]StatusEnum{
		AttachedItemActiveStatus,
		AttachedItemArchivedStatus,
		AttachedItemDeletedStatus,
	}, StatusEnum(name))
}

func KnownCouponCodeStatus(name string) bool {
	return slices.Contains([]StatusEnum{
		CouponCodeNotRedeemedStatus,
		CouponCodeRedeemedStatus,
		CouponCodeArchivedStatus,
	}, StatusEnum(name))
}

func KnownCouponStatus(name string) bool {
	return slices.Contains([]StatusEnum{
		CouponActiveStatus,
		CouponExpiredStatus,
		CouponArchivedStatus,
		CouponDeletedStatus,
	}, StatusEnum(name))
}

func KnownItemDifferentialPriceStatus(name string) bool {
	return slices.Contains([]StatusEnum{
		ItemDifferentialPriceActive,
		ItemDifferentialPriceDeleted,
	}, StatusEnum(name))
}

func KnownItemFamilyStatus(name string) bool {
	return slices.Contains([]StatusEnum{
		ItemFamilyActiveStatus,
		ItemFamilyDeletedStatus,
	}, StatusEnum(name))
}

func KnownItemPriceStatus(name string) bool {
	return slices.Contains([]StatusEnum{
		ItemPriceActiveStatus,
		ItemPriceArchivedStatus,
		ItemPriceDeletedStatus,
	}, StatusEnum(name))
}

func KnownItemStatus(name string) bool {
	return slices.Contains([]StatusEnum{
		ItemActiveStatus,
		ItemArchivedStatus,
		ItemDeletedStatus,
	}, StatusEnum(name))
}

func KnownOrderStatus(name string) bool {
	return slices.Contains([]StatusEnum{
		OrderNewStatus,
		OrderProcessingStatus,
		OrderCompleteStatus,
		OrderCancelledStatus,
		OrderVoidedStatus,
		OrderQueuedStatus,
		OrderAwaitingShipmentStatus,
		OrderOnHoldStatus,
		OrderDeliveredStatus,
		OrderShippedStatus,
		OrderPartiallyDeliveredStatus,
		OrderReturnedStatus,
	}, StatusEnum(name))
}

func KnownPaymentSourceStatus(name string) bool {
	return slices.Contains([]StatusEnum{
		PaymentSourceValidStatus,
		PaymentSourceExpiringStatus,
		PaymentSourceExpiredStatus,
		PaymentSourceInvalidStatus,
		PaymentSourcePendingVerificationStatus,
	}, StatusEnum(name))
}

func KnownPaymentVoucherStatus(name string) bool {
	return slices.Contains([]StatusEnum{
		PaymentVoucherActiveStatus,
		PaymentVoucherConsumedStatus,
		PaymentVoucherExpiredStatus,
		PaymentVoucherFailureStatus,
	}, StatusEnum(name))
}

func KnownTransactionStatus(name string) bool {
	return slices.Contains([]StatusEnum{
		TransactionInProgressStatus,
		TransactionSuccessStatus,
		TransactionVoidedStatus,
		TransactionFailureStatus,
		TransactionTimeoutStatus,
		TransactionNeedsAttentionStatus,
	}, StatusEnum(name))
}

func KnownQuoteStatus(name string) bool {
	return slices.Contains([]StatusEnum{
		QuoteOpenStatus,
		QuoteAcceptedStatus,
		QuoteDeclinedStatus,
		QuoteInvoicedStatus,
		QuoteClosedStatus,
	}, StatusEnum(name))
}

func KnownGiftStatus(name string) bool {
	return slices.Contains([]StatusEnum{
		GiftScheduledStatus,
		GiftUnclaimedStatus,
		GiftClaimedStatus,
		GiftCancelledStatus,
		GiftExpiredStatus,
	}, StatusEnum(name))
}

func KnownRampStatus(name string) bool {
	return slices.Contains([]StatusEnum{
		RampScheduledStatus,
		RampSucceededStatus,
		RampFailedStatus,
		RampDraftStatus,
	}, StatusEnum(name))
}

func KnownSubscriptionStatus(name string) bool {
	return slices.Contains([]StatusEnum{
		SubscriptionFutureStatus,
		SubscriptionInTrialStatus,
		SubscriptionActiveStatus,
		SubscriptionNonRenewingStatus,
		SubscriptionPausedStatus,
		SubscriptionCancelledStatus,
		SubscriptionTransferredStatus,
	}, StatusEnum(name))
}
