package enums

import "slices"

type TypeEnum string

const (
	EntitlementFeatureTypeSwitch   TypeEnum = "switch"
	EntitlementFeatureTypeCustom            = "custom"
	EntitlementFeatureTypeQuantity          = "quantity"
	EntitlementFeatureTypeRange             = "range"

	CreditNoteTypeAdjustment = "adjustment"
	CreditNoteTypeRefundable = "refundable"

	PaymentReferenceNumberKIDType            = "kid"
	PaymentReferenceNumberOCRType            = "ocr"
	PaymentReferenceNumberFRNType            = "frn"
	PaymentReferenceNumberFIKType            = "fik"
	PaymentReferenceNumberSwissReferenceType = "swiss_reference"

	PromotionalCreditIncrementType = "increment"
	PromotionalCreditDecrementType = "decrement"

	AttachedItemRecommendedType = "recommended"
	AttachedItemMandatoryType   = "mandatory"
	AttachedItemOptionalType    = "optional"

	ItemPlanType   = "plan"
	ItemAddonType  = "addon"
	ItemChargeType = "charge"

	PaymentSourceCard                   = "card"
	PaymentSourcePaypalExpressCheckout  = "paypal_express_checkout"
	PaymentSourceAmazonPayments         = "amazon_payments"
	PaymentSourceDirectDebit            = "direct_debit"
	PaymentSourceGeneric                = "generic"
	PaymentSourceAliPay                 = "alipay"
	PaymentSourceUnionPay               = "unionpay"
	PaymentSourceApplePay               = "apple_pay"
	PaymentSourceWeChatPay              = "wechat_pay"
	PaymentSourceIdealPay               = "ideal"
	PaymentSourceGooglePay              = "google_pay"
	PaymentSourceSofort                 = "sofort"
	PaymentSourceBancontact             = "bancontact"
	PaymentSourceGiropay                = "giropay"
	PaymentSourceDotpay                 = "dotpay"
	PaymentSourceUPI                    = "upi"
	PaymentSourceNetbankingEmandates    = "netbanking_emandates"
	PaymentSourcesVenmo                 = "venmo"
	PaymentSourcesPayTo                 = "pay_to"
	PaymentSourcesFasterPayments        = "faster_payments"
	PaymentSourcesSepaInstantTransfer   = "sepa_instant_transfer"
	PaymentSourcesAutomatedBankTransfer = "automated_bank_transfer"
	PaymentSourcesKlarnaPayNow          = "klarna_pay_now"
	PaymentSourcesOnlineBankingPoland   = "online_banking_poland"

	TransactionAuthorizationType   = "authorization"
	TransactionPaymentType         = "payment"
	TransactionRefundType          = "refund"
	TransactionPaymentReversalType = "payment_reversal"

	DiscountFixedAmountType = "fixed_amount"
	DiscountPercentageType  = "percentage"
)

func KnownEntitlementFeatureType(name string) bool {
	return slices.Contains([]TypeEnum{
		EntitlementFeatureTypeSwitch,
		EntitlementFeatureTypeCustom,
		EntitlementFeatureTypeQuantity,
		EntitlementFeatureTypeRange,
	}, TypeEnum(name))
}

func KnownCreditNoteType(name string) bool {
	return slices.Contains([]TypeEnum{
		CreditNoteTypeAdjustment,
		CreditNoteTypeRefundable,
	}, TypeEnum(name))
}

func KnownPaymentReferenceNumberType(name string) bool {
	return slices.Contains([]TypeEnum{
		PaymentReferenceNumberKIDType,
		PaymentReferenceNumberOCRType,
		PaymentReferenceNumberFRNType,
		PaymentReferenceNumberFIKType,
		PaymentReferenceNumberSwissReferenceType,
	}, TypeEnum(name))
}

func KnownPromotionalCreditType(name string) bool {
	return slices.Contains([]TypeEnum{
		PromotionalCreditIncrementType,
		PromotionalCreditDecrementType,
	}, TypeEnum(name))
}

func KnownAttachedItemType(name string) bool {
	return slices.Contains([]TypeEnum{
		AttachedItemRecommendedType,
		AttachedItemMandatoryType,
		AttachedItemOptionalType,
	}, TypeEnum(name))
}

func KnownItemType(name string) bool {
	return slices.Contains([]TypeEnum{
		ItemPlanType,
		ItemAddonType,
		ItemChargeType,
	}, TypeEnum(name))
}

func KnownPaymentSourceType(name string) bool {
	return slices.Contains([]TypeEnum{
		PaymentSourceCard,
		PaymentSourcePaypalExpressCheckout,
		PaymentSourceAmazonPayments,
		PaymentSourceDirectDebit,
		PaymentSourceGeneric,
		PaymentSourceAliPay,
		PaymentSourceUnionPay,
		PaymentSourceApplePay,
		PaymentSourceWeChatPay,
		PaymentSourceIdealPay,
		PaymentSourceGooglePay,
		PaymentSourceSofort,
		PaymentSourceBancontact,
		PaymentSourceGiropay,
		PaymentSourceDotpay,
		PaymentSourceUPI,
		PaymentSourceNetbankingEmandates,
		PaymentSourcesVenmo,
		PaymentSourcesPayTo,
		PaymentSourcesFasterPayments,
		PaymentSourcesSepaInstantTransfer,
		PaymentSourcesAutomatedBankTransfer,
		PaymentSourcesKlarnaPayNow,
		PaymentSourcesOnlineBankingPoland,
	}, TypeEnum(name))
}

func KnownTransactionType(name string) bool {
	return slices.Contains([]TypeEnum{
		TransactionAuthorizationType,
		TransactionPaymentType,
		TransactionRefundType,
		TransactionPaymentReversalType,
	}, TypeEnum(name))
}

func KnownDiscountType(name string) bool {
	return slices.Contains([]TypeEnum{
		DiscountFixedAmountType,
		DiscountPercentageType,
	}, TypeEnum(name))
}
