package enums_manual

import "slices"

type EventTypeEnum string

const (
	CouponCreatedEvent                             EventTypeEnum = "coupon_created"
	CouponUpdatedEvent                                           = "coupon_updated"
	CouponDeletedEvent                                           = "coupon_deleted"
	CouponSetCreatedEvent                                        = "coupon_set_created"
	CouponSetUpdatedEvent                                        = "coupon_set_updated"
	CouponSetDeletedEvent                                        = "coupon_set_deleted"
	CouponCodesAddedEvent                                        = "coupon_codes_added"
	CouponCodesDeletedEvent                                      = "coupon_codes_deleted"
	CouponCodesUpdatedEvent                                      = "coupon_codes_updated"
	CustomerCreatedEvent                                         = "customer_created"
	CustomerChangedEvent                                         = "customer_changed"
	CustomerDeletedEvent                                         = "customer_deleted"
	CustomerMovedOutEvent                                        = "customer_moved_out"
	CustomerMovedInEvent                                         = "customer_moved_in"
	PromotionalCreditsAddedEvent                                 = "promotional_credits_added"
	PromotionalCreditsDeducedEvent                               = "promotional_credits_deducted"
	SubscriptionCreatedEvent                                     = "subscription_created"
	SubscriptionCreatedWithBackdatingEvent                       = "subscription_created_with_backdating"
	SubscriptionStartedEvent                                     = "subscription_started"
	SubscriptionTrialEndReminderEvent                            = "subscription_trial_end_reminder"
	SubscriptionActivatedEvent                                   = "subscription_activated"
	SubscriptionActivatedWIthBackdatingEvent                     = "subscription_activated_with_backdating"
	SubscriptionChangedEvent                                     = "subscription_changed"
	SubscriptionTrialExtendedEvent                               = "subscription_trial_extended"
	MrrUpdatedEvent                                              = "mrr_updated"
	SubscriptionChangedWithBackdatingEvent                       = "subscription_changed_with_backdating"
	SubscriptionCancellationScheduledEvent                       = "subscription_cancellation_scheduled"
	SubscriptionCancellationReminderEvent                        = "subscription_cancellation_reminder"
	SubscriptionCancelledEvent                                   = "subscription_canceled"
	SubscriptionCancelledWithBackdatingEvent                     = "subscription_canceled_with_backdating"
	SubscriptionReactivatedEvent                                 = "subscription_reactivated"
	SubscriptionReactivatedWithBackdatingEvent                   = "subscription_reactivated_with_backdating"
	SubscriptionRenewedEvent                                     = "subscription_renewed"
	SubscriptionItemsRenewedEvent                                = "subscription_items_renewed"
	SubscriptionScheduledCancellationRemovedEvent                = "subscription_scheduled_cancellation_removed"
	SubscriptionChangesScheduledEvent                            = "subscription_changes_scheduled"
	SubscriptionScheduledChangesRemovedEvent                     = "subscription_scheduled_changes_removed"
	SubscriptionShippingAddressUpdatedEvent                      = "subscription_shipping_address_updated"
	SubscriptionDeletedEvent                                     = "subscription_deleted"
	SubscriptionPausedEvent                                      = "subscription_paused"
	SubscriptionPauseScheduledEvent                              = "subscription_pause_scheduled"
	SubscriptionScheduledPauseRemovedEvent                       = "subscription_scheduled_pause_removed"
	SubscriptionResumedEvent                                     = "subscription_resumed"
	SubscriptionMovedInEvent                                     = "subscription_moved_in"
	SubscriptionMovedOutEvent                                    = "subscription_moved_out"
	SubscriptionMovementFailedEvent                              = "subscription_movement_failed"
	SubscriptionResumeScheduledEvent                             = "subscription_resumption_scheduled"
	SubscriptionScheduledResumeRemovedEvent                      = "subscription_scheduled_resumption_removed"
	SubscriptionAdvanceInvoiceScheduleAddedEvent                 = "subscription_advance_invoice_schedule_added"
	SubscriptionAdvanceInvoiceScheduleUpdatedEvent               = "subscription_advance_invoice_schedule_updated"
	SubscriptionAdvanceInvoiceScheduleRemovedEvent               = "subscription_advance_invoice_schedule_removed"
	PendingInvoiceCreatedEvent                                   = "pending_invoice_created"
	PendingInvoiceUpdatedEvent                                   = "pending_invoice_updated"
	InvoiceGeneratedEvent                                        = "invoice_generated"
	InvoiceGeneratedWithBackdatingEvent                          = "invoice_generated_with_backdating"
	InvoiceUpdatedEvent                                          = "invoice_updated"
	InvoiceDeletedEvent                                          = "invoice_deleted"
	CreditNoteCreatedEvent                                       = "credit_note_created"
	CreditNoteCreatedWithBackdatingEvent                         = "credit_note_created_with_backdating"
	CreditNoteUpdatedEvent                                       = "credit_note_updated"
	CreditNoteDeletedEvent                                       = "credit_note_deleted"
	SubscriptionRenewalReminderEvent                             = "subscription_renewal_reminder"
	AddUsagesReminderEvent                                       = "add_usages_reminder"
	TransactionCreatedEvent                                      = "transaction_created"
	TransactionUpdatedEvent                                      = "transaction_updated"
	TransactionDeletedEvent                                      = "transaction_deleted"
	PaymentSucceededEvent                                        = "payment_succeeded"
	PaymentFailedEvent                                           = "payment_failed"
	PaymentRefundedEvent                                         = "payment_refunded"
	PaymentInitiatedEvent                                        = "payment_initiated"
	RefundInitiatedEvent                                         = "refund_initiated"
	AuthorizationSucceededEvent                                  = "authorization_succeeded"
	AuthorizationVoidedEvent                                     = "authorization_voided"
	CardAddedEvent                                               = "card_added"
	CardUpdatedEvent                                             = "card_updated"
	CardExpiryReminderEvent                                      = "card_expiry_reminder"
	CardExpiredEvent                                             = "card_expired"
	CardDeletedEvent                                             = "card_deleted"
	PaymentSourceAddedEvent                                      = "payment_source_added"
	PaymentSourceUpdatedEvent                                    = "payment_source_updated"
	PaymentSourceDeletedEvent                                    = "payment_source_deleted"
	PaymentSourceExpiringEvent                                   = "payment_source_expiring"
	PaymentSourceExpiredEvent                                    = "payment_source_expired"
	PaymentSourceLocallyDeletedEvent                             = "payment_source_locally_deleted"
	VirtualBankAccountAddedEvent                                 = "virtual_bank_account_added"
	VirtualBankAccountUpdatedEvent                               = "virtual_bank_account_updated"
	VirtualBankAccountDeletedEvent                               = "virtual_bank_account_deleted"
	TokenCreatedEvent                                            = "token_created"
	TokenConsumedEvent                                           = "token_consumed"
	TokenExpiredEvent                                            = "token_expired"
	UnBilledChargesCreatedEvent                                  = "unbilled_charges_created"
	UnBilledChargesVoidedEvent                                   = "unbilled_charges_voided"
	UnBilledChargesDeletedEvent                                  = "unbilled_charges_deleted"
	UnBilledChargesInvoicedEvent                                 = "unbilled_charges_invoiced"
	OrderCreatedEvent                                            = "order_created"
	OrderUpdatedEvent                                            = "order_updated"
	OrderCancelledEvent                                          = "order_canceled"
	OrderDeliveredEvent                                          = "order_delivered"
	OrderReturnedEvent                                           = "order_returned"
	OrderReadyToProcessEvent                                     = "order_ready_to_process"
	OrderReadyToShipEvent                                        = "order_ready_to_ship"
	OrderDeletedEvent                                            = "order_deleted"
	OrderResentEvent                                             = "order_resent"
	QuoteCreatedEvent                                            = "quote_created"
	QuoteUpdatedEvent                                            = "quote_updated"
	QuoteDeletedEvent                                            = "quote_deleted"
	TaxWithheldRecordedEvent                                     = "tax_withheld_recorded"
	TaxWithheldDeletedEvent                                      = "tax_withheld_deleted"
	TaxWithheldRefundedEvent                                     = "tax_withheld_refunded"
	GiftScheduledEvent                                           = "gift_scheduled"
	GiftUnclaimedEvent                                           = "gift_unclaimed"
	GiftClaimedEvent                                             = "gift_claimed"
	GiftExpiredEvent                                             = "gift_expired"
	GiftCancelledEvent                                           = "gift_canceled"
	GiftUpdatedEvent                                             = "gift_updated"
	HierarchyCreatedEvent                                        = "hierarchy_created"
	HierarchyDeletedEvent                                        = "hierarchy_deleted"
	PaymentIntentCreatedEvent                                    = "payment_intent_created"
	PaymentIntentUpdatedEvent                                    = "payment_intent_updated"
	ContractTermCreatedEvent                                     = "contract_term_created"
	ContractTermRenewedEvent                                     = "contract_term_renewed"
	ContractTermTerminatedEvent                                  = "contract_term_terminated"
	ContractTermCompletedEvent                                   = "contract_term_completed"
	ContractTermCancelledEvent                                   = "contract_term_canceled"
	ItemFamilyCreatedEvent                                       = "item_family_created"
	ItemFamilyUpdatedEvent                                       = "item_family_updated"
	ItemFamilyDeletedEvent                                       = "item_family_deleted"
	ItemCreatedEvent                                             = "item_created"
	ItemDeletedEvent                                             = "item_deleted"
	ItemPriceCreatedEvent                                        = "item_price_created"
	ItemPriceUpdatedEvent                                        = "item_price_updated"
	ItemPriceDeletedEvent                                        = "item_price_deleted"
	AttachedItemCreatedEvent                                     = "attached_item_created"
	AttachedItemUpdatedEvent                                     = "attached_item_updated"
	AttachedItemDeletedEvent                                     = "attached_item_deleted"
	DifferentialPriceCreatedEvent                                = "differential_price_created"
	DifferentialPriceUpdatedEvent                                = "differential_price_updated"
	DifferentialPriceDeletedEvent                                = "differential_price_deleted"
	FeatureCreatedEvent                                          = "feature_created"
	FeatureUpdatedEvent                                          = "feature_updated"
	FeatureDeletedEvent                                          = "feature_deleted"
	FeatureActivatedEvent                                        = "feature_activated"
	FeatureReactivatedEvent                                      = "feature_reactivated"
	FeatureArchivedEvent                                         = "feature_archived"
	ItemEntitlementsUpdatedEvent                                 = "item_entitlements_updated"
	ItemPriceEntitlementsUpdatedEvent                            = "item_price_entitlements_updated"
	EntitlementOverridesUpdatedEvent                             = "entitlement_overrides_updated"
	EntitlementOverridesRemovedEvent                             = "entitlement_overrides_removed"
	ItemEntitlementsRemovedEvent                                 = "item_entitlements_removed"
	ItemPriceEntitlementsRemovedEvent                            = "item_price_entitlements_removed"
	EntitlementOverridesAutoRemovedEvent                         = "entitlement_overrides_auto_removed"
	SubscriptionEntitlementsCreatedEvent                         = "subscription_entitlements_created"
	BusinessEntityCreatedEvent                                   = "business_entity_created"
	BusinessEntityUpdatedEvent                                   = "business_entity_updated"
	BusinessEntityDeletedEvent                                   = "business_entity_deleted"
	PurchaseCreatedEvent                                         = "purchase_created"
	VoucherCreatedEvent                                          = "voucher_created"
	VoucherExpiredEvent                                          = "voucher_expired"
	VoucherCreateFailed                                          = "voucher_create_failed"
	PriceVariantCreated                                          = "price_variant_created"
	PriceVariantUpdated                                          = "price_variant_updated"
	PriceVariantDeleted                                          = "price_variant_deleted"
	InstallmentConfigCreated                                     = "installment_config_created"
	InstallmentConfigDeleted                                     = "installment_config_deleted"
	InvoiceInstallmentsCreated                                   = "invoice_installments_created"
	InvoiceInstallmentUpdated                                    = "invoice_installment_updated"
	RampCreated                                                  = "ramp_created"
	RampDeleted                                                  = "ramp_deleted"
	RampApplied                                                  = "ramp_applied"
	SubscriptionRampUpdated                                      = "subscription_ramp_updated"
	SubscriptionRampDrafted                                      = "subscription_ramp_drafted"
	CustomerEntitlementsUpdated                                  = "customer_entitlements_updated"
)

func KnownEventType(name string) bool {
	return slices.Contains([]EventTypeEnum{
		CouponCreatedEvent,
		CouponUpdatedEvent,
		CouponDeletedEvent,
		CouponSetCreatedEvent,
		CouponSetUpdatedEvent,
		CouponSetDeletedEvent,
		CouponCodesAddedEvent,
		CouponCodesDeletedEvent,
		CouponCodesUpdatedEvent,
		CustomerCreatedEvent,
		CustomerChangedEvent,
		CustomerDeletedEvent,
		CustomerMovedOutEvent,
		CustomerMovedInEvent,
		PromotionalCreditsAddedEvent,
		PromotionalCreditsDeducedEvent,
		SubscriptionCreatedEvent,
		SubscriptionCreatedWithBackdatingEvent,
		SubscriptionStartedEvent,
		SubscriptionTrialEndReminderEvent,
		SubscriptionActivatedEvent,
		SubscriptionActivatedWIthBackdatingEvent,
		SubscriptionChangedEvent,
		SubscriptionTrialExtendedEvent,
		MrrUpdatedEvent,
		SubscriptionChangedWithBackdatingEvent,
		SubscriptionCancellationScheduledEvent,
		SubscriptionCancellationReminderEvent,
		SubscriptionCancelledEvent,
		SubscriptionCancelledWithBackdatingEvent,
		SubscriptionReactivatedEvent,
		SubscriptionReactivatedWithBackdatingEvent,
		SubscriptionRenewedEvent,
		SubscriptionItemsRenewedEvent,
		SubscriptionScheduledCancellationRemovedEvent,
		SubscriptionChangesScheduledEvent,
		SubscriptionScheduledChangesRemovedEvent,
		SubscriptionShippingAddressUpdatedEvent,
		SubscriptionDeletedEvent,
		SubscriptionPausedEvent,
		SubscriptionPauseScheduledEvent,
		SubscriptionScheduledPauseRemovedEvent,
		SubscriptionResumedEvent,
		SubscriptionMovedInEvent,
		SubscriptionMovedOutEvent,
		SubscriptionMovementFailedEvent,
		SubscriptionResumeScheduledEvent,
		SubscriptionScheduledResumeRemovedEvent,
		SubscriptionAdvanceInvoiceScheduleAddedEvent,
		SubscriptionAdvanceInvoiceScheduleUpdatedEvent,
		SubscriptionAdvanceInvoiceScheduleRemovedEvent,
		PendingInvoiceCreatedEvent,
		PendingInvoiceUpdatedEvent,
		InvoiceGeneratedEvent,
		InvoiceGeneratedWithBackdatingEvent,
		InvoiceUpdatedEvent,
		InvoiceDeletedEvent,
		CreditNoteCreatedEvent,
		CreditNoteCreatedWithBackdatingEvent,
		CreditNoteUpdatedEvent,
		CreditNoteDeletedEvent,
		SubscriptionRenewalReminderEvent,
		AddUsagesReminderEvent,
		TransactionCreatedEvent,
		TransactionUpdatedEvent,
		TransactionDeletedEvent,
		PaymentSucceededEvent,
		PaymentFailedEvent,
		PaymentRefundedEvent,
		PaymentInitiatedEvent,
		RefundInitiatedEvent,
		AuthorizationSucceededEvent,
		AuthorizationVoidedEvent,
		CardAddedEvent,
		CardUpdatedEvent,
		CardExpiryReminderEvent,
		CardExpiredEvent,
		CardDeletedEvent,
		PaymentSourceAddedEvent,
		PaymentSourceUpdatedEvent,
		PaymentSourceDeletedEvent,
		PaymentSourceExpiringEvent,
		PaymentSourceExpiredEvent,
		PaymentSourceLocallyDeletedEvent,
		VirtualBankAccountAddedEvent,
		VirtualBankAccountUpdatedEvent,
		VirtualBankAccountDeletedEvent,
		TokenCreatedEvent,
		TokenConsumedEvent,
		TokenExpiredEvent,
		UnBilledChargesCreatedEvent,
		UnBilledChargesVoidedEvent,
		UnBilledChargesDeletedEvent,
		UnBilledChargesInvoicedEvent,
		OrderCreatedEvent,
		OrderUpdatedEvent,
		OrderCancelledEvent,
		OrderDeliveredEvent,
		OrderReturnedEvent,
		OrderReadyToProcessEvent,
		OrderReadyToShipEvent,
		OrderDeletedEvent,
		OrderResentEvent,
		QuoteCreatedEvent,
		QuoteUpdatedEvent,
		QuoteDeletedEvent,
		TaxWithheldRecordedEvent,
		TaxWithheldDeletedEvent,
		TaxWithheldRefundedEvent,
		GiftScheduledEvent,
		GiftUnclaimedEvent,
		GiftClaimedEvent,
		GiftExpiredEvent,
		GiftCancelledEvent,
		GiftUpdatedEvent,
		HierarchyCreatedEvent,
		HierarchyDeletedEvent,
		PaymentIntentCreatedEvent,
		PaymentIntentUpdatedEvent,
		ContractTermCreatedEvent,
		ContractTermRenewedEvent,
		ContractTermTerminatedEvent,
		ContractTermCompletedEvent,
		ContractTermCancelledEvent,
		ItemFamilyCreatedEvent,
		ItemFamilyUpdatedEvent,
		ItemFamilyDeletedEvent,
		ItemCreatedEvent,
		ItemDeletedEvent,
		ItemPriceCreatedEvent,
		ItemPriceUpdatedEvent,
		ItemPriceDeletedEvent,
		AttachedItemCreatedEvent,
		AttachedItemUpdatedEvent,
		AttachedItemDeletedEvent,
		DifferentialPriceCreatedEvent,
		DifferentialPriceUpdatedEvent,
		DifferentialPriceDeletedEvent,
		FeatureCreatedEvent,
		FeatureUpdatedEvent,
		FeatureDeletedEvent,
		FeatureActivatedEvent,
		FeatureReactivatedEvent,
		FeatureArchivedEvent,
		ItemEntitlementsUpdatedEvent,
		ItemPriceEntitlementsUpdatedEvent,
		EntitlementOverridesUpdatedEvent,
		EntitlementOverridesRemovedEvent,
		ItemEntitlementsRemovedEvent,
		ItemPriceEntitlementsRemovedEvent,
		EntitlementOverridesAutoRemovedEvent,
		SubscriptionEntitlementsCreatedEvent,
		BusinessEntityCreatedEvent,
		BusinessEntityUpdatedEvent,
		BusinessEntityDeletedEvent,
		PurchaseCreatedEvent,
		VoucherCreatedEvent,
		VoucherExpiredEvent,
		VoucherCreateFailed,
		PriceVariantCreated,
		PriceVariantUpdated,
		PriceVariantDeleted,
		InstallmentConfigCreated,
		InstallmentConfigDeleted,
		InvoiceInstallmentsCreated,
		InvoiceInstallmentUpdated,
		RampCreated,
		RampDeleted,
		RampApplied,
		SubscriptionRampUpdated,
		SubscriptionRampDrafted,
		CustomerEntitlementsUpdated,
	}, EventTypeEnum(name))
}
