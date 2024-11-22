package enums

// THIS IS GENERATED CODE. DO NOT EDIT.

type TypeEnum string

const (
	EntitlementFeatureTypeCustom   TypeEnum = "custom"
	EntitlementFeatureTypeQuantity TypeEnum = "quantity"
	EntitlementFeatureTypeRange    TypeEnum = "range"
	EntitlementFeatureTypeSwitch   TypeEnum = "switch"
)

const (
	CreditNoteTypeAdjustment TypeEnum = "adjustment"
	CreditNoteTypeRefundable TypeEnum = "refundable"
)

const (
	PaymentReferenceNumberTypeFik TypeEnum = "fik"
	PaymentReferenceNumberTypeFrn TypeEnum = "frn"
	PaymentReferenceNumberTypeKid TypeEnum = "kid"
	PaymentReferenceNumberTypeOcr TypeEnum = "ocr"
)

const (
	PromotionalCreditTypeDecrement TypeEnum = "decrement"
	PromotionalCreditTypeIncrement TypeEnum = "increment"
)

const (
	AttachedItemTypeMandatory   TypeEnum = "mandatory"
	AttachedItemTypeOptional    TypeEnum = "optional"
	AttachedItemTypeRecommended TypeEnum = "recommended"
)

const (
	ItemTypeAddon  TypeEnum = "addon"
	ItemTypeCharge TypeEnum = "charge"
	ItemTypePlan   TypeEnum = "plan"
)

const (
	PaymentSourceTypeAmazonPayments        TypeEnum = "amazon_payments"
	PaymentSourceTypeCard                  TypeEnum = "card"
	PaymentSourceTypeDirectDebit           TypeEnum = "direct_debit"
	PaymentSourceTypePaypalExpressCheckout TypeEnum = "paypal_express_checkout"
)

const (
	TransactionTypeAuthorization   TypeEnum = "authorization"
	TransactionTypePayment         TypeEnum = "payment"
	TransactionTypePaymentReversal TypeEnum = "payment_reversal"
	TransactionTypeRefund          TypeEnum = "refund"
)

const (
	DiscountTypeFixedAmount TypeEnum = "fixed_amount"
	DiscountTypePercentage  TypeEnum = "percentage"
)
