package enums

import "github.com/bitsnap/chargebee-api-client/codegen/templates"

type ScrappedEnum struct {
	Name       string
	CustomType string
	Variants   map[string]string
}

func (en *ScrappedEnum) String() string {
	t := "string"
	if en.CustomType != "" {
		t = en.CustomType
	}

	return `
type ` + en.Name + ` ` + t + `
const (
	
)
`
}

func EnumHeader(content string) string {
	return `package enums

// THIS IS GENERATED CODE. DO NOT EDIT.
`
}

var content = map[string][]func() string{
	"api_version.go": {func() string {
		return templates.GenerateStringEnum("ApiVersionEnum", "api_version", "https://apidocs.chargebee.com/docs/api/events")
	}},
	"apply_on.go": {func() string {
		return templates.GenerateStringEnum("ApplyOnEnum", "apply_on", "https://apidocs.chargebee.com/docs/api/coupons")
	}},
	"authorization_reason.go": {func() string {
		return templates.GenerateStringEnum("AuthorizationReasonEnum", "authorization_reason", "https://apidocs.chargebee.com/docs/api/transactions")
	}},
	"autocollection.go": {func() string {
		return templates.GenerateStringEnum("AutoCollectionEnum", "auto_collection", "https://apidocs.chargebee.com/docs/api/customers")
	}},
	"billing.go": {func() string {
		return templates.GenerateStringEnum("BillingDayOfWeekEnum", "billing_day_of_week", "https://apidocs.chargebee.com/docs/api/customers")
	}, func() string {
		return templates.GenerateStringEnum("BillingDateModeEnum", "billing_date_mode", "https://apidocs.chargebee.com/docs/api/customers")
	}, func() string {
		return templates.GenerateAliasEnum("BillingDayOfWeekModeEnum", "BillingDateModeEnum")
	}},
	"cancellation_reason.go": {func() string {
		return templates.GeneratePreffixedStringEnum("CancellationReasonEnum", "Order", "cancellation_reason", "https://apidocs.chargebee.com/docs/api/orders")
	}},
	"cancel_reason.go": {func() string {
		return templates.GeneratePreffixedStringEnum("CancelReasonEnum", "Subscription", "cancel_reason", "https://apidocs.chargebee.com/docs/api/subscriptions")
	}},
	"change_option.go": {func() string {
		return templates.GeneratePreffixedStringEnum("ChangeOptionEnum", "QuotedSubscription", "change_option", "https://apidocs.chargebee.com/docs/api/quoted_subscriptions")
	}},
	"channel.go": {func() string {
		return templates.GeneratePreffixedStringEnum("ChannelEnum", "Item", "channel", "https://apidocs.chargebee.com/docs/api/items")
	}},
	"charge_event.go": {func() string {
		return templates.GeneratePreffixedStringEnum("ChargeEventEnum", "QuoteLineGroup", "charge_event", "https://apidocs.chargebee.com/docs/api/quote_line_groups")
	}},
	"charge_on_event.go": {func() string {
		return templates.GeneratePreffixedStringEnum("ChargeOnEventEnum", "AttachedItem", "charge_on_event", "https://apidocs.chargebee.com/docs/api/attached_items")
	}},
	"credit_type.go": {func() string {
		return templates.GenerateStringEnum("CreditTypeEnum", "credit_type", "https://apidocs.chargebee.com/docs/api/promotional_credits")
	}},
	"customer_type.go": {func() string {
		return templates.GenerateStringEnum("CustomerTypeEnum", "customer_type", "https://apidocs.chargebee.com/docs/api/customers")
	}},
	"discount_type.go": {func() string {
		return templates.GenerateStringEnum("DiscountTypeEnum", "discount_type", "https://apidocs.chargebee.com/docs/api/coupons")
	}},
	"dunning_status.go": {func() string {
		return templates.GeneratePreffixedStringEnum("DunningStatusEnum", "Invoice", "dunning_status", "https://apidocs.chargebee.com/docs/api/invoices")
	}},
	"duration_type.go": {func() string {
		return templates.GeneratePreffixedStringEnum("DurationTypeEnum", "Coupon", "duration_type", "https://apidocs.chargebee.com/docs/api/coupons")
	}},
	"einvoicing_method.go": {func() string {
		return templates.GeneratePreffixedStringEnum("EinvoicingMethodEnum", "Invoicing", "einvoicing_method", "https://apidocs.chargebee.com/docs/api/customers")
	}},
	"entity_code.go": {func() string {
		return templates.GenerateStringEnum("EntityCodeEnum", "entity_code", "https://apidocs.chargebee.com/docs/api/customers")
	}},
	"entity_type.go": {func() string {
		return templates.GeneratePreffixedStringEnum("EntityTypeEnum", "Entity", "entity_type", "https://apidocs.chargebee.com/docs/api/entitlements")
	}},
	"event_type.go": {func() string {
		return templates.GenerateStringEnum("EventTypeEnum", "event_type", "https://apidocs.chargebee.com/docs/api/events")
	}},
	"forex_type.go": {func() string {
		return templates.GeneratePreffixedStringEnum("ForexTypeEnum", "Forex", "forex_type", "https://apidocs.chargebee.com/docs/api/currencies")
	}},
	"fraud_flag.go": {func() string {
		return templates.GenerateStringEnum("FraudFlagEnum", "fraud_flag", "https://apidocs.chargebee.com/docs/api/transactions")
	}},
	"free_period_unit.go": {func() string {
		return templates.GenerateAliasEnum("FreePeriodUnitEnum", "PeriodUnitEnum")
	}},
	"gateway.go": {func() string {
		return templates.GenerateStringEnum("GatewayEnum", "gateway", "https://apidocs.chargebee.com/docs/api/transactions")
	}},
	"initiator_type.go": {func() string {
		return templates.GenerateStringEnum("InitiatorTypeEnum", "initiator_type", "https://apidocs.chargebee.com/docs/api/transactions")
	}},
	"item_applicability.go": {func() string {
		return templates.GenerateStringEnum("ItemApplicabilityEnum", "item_applicability", "https://apidocs.chargebee.com/docs/api/items")
	}},
	"item_type.go": {func() string {
		return templates.GeneratePreffixedStringEnum("ItemTypeEnum", "Item", "item_type", "https://apidocs.chargebee.com/docs/api/item_entitlements")
	}},
	"operation_type.go": {func() string {
		return templates.GenerateStringEnum("OperationTypeEnum", "operation_type", "https://apidocs.chargebee.com/docs/api/quotes")
	}},
	"order_type.go": {func() string {
		return templates.GeneratePreffixedStringEnum("OrderTypeEnum", "Order", "order_type", "https://apidocs.chargebee.com/docs/api/orders")
	}},
	"payment_method.go": {func() string {
		return templates.GenerateStringEnum("PaymentMethodEnum", "payment_method", "https://apidocs.chargebee.com/docs/api/transactions")
	}},
	"payment_status.go": {func() string {
		return templates.GenerateStringEnum("PaymentStatusEnum", "payment_status", "https://apidocs.chargebee.com/docs/api/orders")
	}},
	"payment_voucher_type.go": {func() string {
		return templates.GenerateStringEnum("PaymentVoucherTypeEnum", "payment_voucher_type", "https://apidocs.chargebee.com/docs/api/payment_vouchers")
	}},
	"period_unit.go": {func() string {
		return templates.GenerateStringEnum("PeriodUnitEnum", "period_unit", "https://apidocs.chargebee.com/docs/api/payment_schedule_schemes")
	}},
	"pii_cleared.go": {func() string {
		return templates.GenerateStringEnum("PiiClearedEnum", "pii_cleared", "https://apidocs.chargebee.com/docs/api/customers")
	}},
	"price_type.go": {func() string {
		return templates.GenerateStringEnum("PriceTypeEnum", "price_type", "https://apidocs.chargebee.com/docs/api/orders")
	}},
	"pricing_model.go": {func() string {
		return templates.GenerateStringEnum("PricingModelEnum", "pricing_model", "https://apidocs.chargebee.com/docs/api/item_prices")
	}},
	"proration_type.go": {func() string {
		return templates.GenerateStringEnum("ProrationTypeEnum", "proration_type", "https://apidocs.chargebee.com/docs/api/item_prices")
	}},
	"reason_code.go": {func() string {
		return templates.GeneratePreffixedStringEnum("ReasonCodeEnum", "CreditNote", "reason_code", "https://apidocs.chargebee.com/docs/api/credit_notes")
	}},
	"reason_code_be.go": {func() string {
		return templates.GeneratePreffixedStringEnum("BusinessEntityReasonCodeEnum", "BusinessEntityTransfer", "reason_code", "https://apidocs.chargebee.com/docs/api/business_entity_transfers")
	}},
	"resent_status.go": {func() string {
		return templates.GenerateStringEnum("ResentStatusEnum", "resent_status", "https://apidocs.chargebee.com/docs/api/orders")
	}},
	"resource_type.go": {func() string {
		return templates.GeneratePreffixedStringEnum("ResourceTypeEnum", "Resource", "resource_type", "https://apidocs.chargebee.com/docs/api/business_entity_transfers")
	}},
	"schedule_type.go": {func() string {
		return templates.GenerateStringEnum("ScheduleTypeEnum", "schedule_type", "https://apidocs.chargebee.com/docs/api/advance_invoice_schedules")
	}},
	"offline_payment_method.go": {func() string {
		return templates.GenerateStringEnum("OfflinePaymentMethodEnum", "offline_payment_method", "https://apidocs.chargebee.com/docs/api/subscriptions")
	}},
	"shipping_period_unit.go": {func() string {
		return templates.GeneratePreffixedStringEnum("ShippingPeriodUnitEnum", "ShippingPeriod", "shipping_period_unit", "https://apidocs.chargebee.com/docs/api/item_prices")
	}},
	"source.go": {func() string {
		return templates.GenerateStringEnum("SourceEnum", "source", "https://apidocs.chargebee.com/docs/api/events")
	}},
	//"status.go": {func() string {
	//	return templates.GenerateStringEnum("StatusEnum", "status", "https://apidocs.chargebee.com/docs/api/events")
	//}},
	"taxability.go": {func() string {
		return templates.GeneratePreffixedStringEnum("TaxabilityEnum", "Taxes", "taxability", "https://apidocs.chargebee.com/docs/api/customers")
	}},
	"trial_end_action.go": {func() string {
		return templates.GenerateStringEnum("TrialEndActionEnum", "trial_end_action", "https://apidocs.chargebee.com/docs/api/subscriptions")
	}},
	"trial_period_unit.go": {func() string {
		return templates.GeneratePreffixedStringEnum("TrialPeriodUnitEnum", "Trial", "trial_period_unit", "https://apidocs.chargebee.com/docs/api/item_prices")
	}},
	//"type.go": {func() string {
	//	return templates.GeneratePreffixedStringEnum("TypeEnum", "Type", "type", "https://apidocs.chargebee.com/docs/api/item_prices")
	//}},
	"usage_accumulation_reset_frequency.go": {func() string {
		return templates.GenerateStringEnum("UsageAccumulationResetFrequencyEnum", "usage_accumulation_reset_frequency", "https://apidocs.chargebee.com/docs/api/item_prices")
	}},
	"usage_calculation.go": {func() string {
		return templates.GenerateStringEnum("UsageCalculationEnum", "usage_calculation", "https://apidocs.chargebee.com/docs/api/items")
	}},
	"validation_status.go": {func() string {
		return templates.GenerateStringEnum("ValidationStatusEnum", "validation_status", "https://apidocs.chargebee.com/docs/api/addresses")
	}},
	"vat_number.go": {func() string {
		return templates.GenerateStringEnum("VatNumberStatusEnum", "vat_number_status", "https://apidocs.chargebee.com/docs/api/customers")
	}},
	"webhook_status.go": {func() string {
		return templates.GenerateStringEnum("WebhookStatusEnum", "webhook_status", "https://apidocs.chargebee.com/docs/api/events")
	}},
}

func GenerateEnumsContent() map[string][]func() string {
	return content
}
