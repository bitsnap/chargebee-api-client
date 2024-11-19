package enums

type ReasonCodeEnum string

const (
	WriteOffCreditNoteReason                 ReasonCodeEnum = "write_off"
	SubscriptionChangeCreditNoteReason                      = "subscription_change"
	SubscriptionCancellationCreditNoteReason                = "subscription_cancellation"
	SubscriptionPauseCreditNoteReason                       = "subscription_pause"
	ChargeBackCreditNoteReason                              = "chargeback"
	ProductUnsatisfactoryCreditNoteReason                   = "product_unsatisfactory"
	ServiceUnsatisfactoryCreditNoteReason                   = "service_unsatisfactory"
	OrderChangeCreditNoteReason                             = "order_change"
	OrderCancellationCreditNoteReason                       = "order_cancellation"
	WaiverCreditNoteReason                                  = "waiver"
	OtherCreditNoteReason                                   = "other"
	FraudulentCreditNoteReason                              = "fraudulent"

	BusinessEntityTransferCorrectionReason = "correction"
)
