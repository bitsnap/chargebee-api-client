package templates

import (
	"embed"
	"strings"

	. "github.com/bitsnap/go-util"
)

//go:embed list_convenience.go.tpl
var listConvTemplateFS embed.FS

var listConvTemplate = MustParseTemplate(listConvTemplateFS, "list_convenience.go.tpl")

type convenienceMethod struct {
	FileName     string
	Name         string
	NameSingular string
	WithID       bool
	ByID         bool
}

type listConvenience struct {
	Methods []convenienceMethod
}

var convenience = listConvenience{Methods: []convenienceMethod{{
	"business_entities.go",
	"BusinessEntityTransfers",
	"BusinessEntityTransfer",
	false,
	false,
}, {
	"currencies.go",
	"Currencies",
	"Currency",
	false,
	true,
}, {
	"customer_contacts.go",
	"CustomerContacts",
	"Contact",
	true,
	false,
}, {
	"customer_hierarchies.go",
	"CustomerHierarchies",
	"Hierarchy",
	true,
	false,
}, {
	"customer_vouchers.go",
	"CustomerPaymentVouchers",
	"PaymentVoucher",
	true,
	false,
}, {
	"customers.go",
	"Customers",
	"Customer",
	false,
	true,
}, {
	"entitlements.go",
	"Entitlements",
	"Entitlement",
	false,
	false,
}, {
	"events.go",
	"Events",
	"Event",
	false,
	true,
}, {
	"features.go",
	"Features",
	"Feature",
	false,
	true,
}, {
	"invoice_credit_notes.go",
	"CreditNotes",
	"CreditNote",
	false,
	true,
}, {
	"invoice_payment_reference_numbers.go",
	"PaymentReferenceNumbers",
	"PaymentReferenceNumber",
	true,
	false,
}, {
	"invoice_payments.go",
	"InvoicePayments",
	"Transaction",
	true,
	false,
}, {
	"invoice_promotional_credits.go",
	"PromotionalCredits",
	"PromotionalCredit",
	false,
	true,
}, {
	"invoice_unbilled_charges.go",
	"InvoiceUnbilledCharges",
	"InvoiceUnbilledCharge",
	false,
	false,
}, {
	"invoice_vouchers.go",
	"InvoicePaymentVouchers",
	"PaymentVoucher",
	true,
	false,
}, {
	"invoices.go",
	"Invoices",
	"Invoice",
	false,
	true,
}, {
	"item_attached.go",
	"ItemsAttached",
	"ItemAttached",
	true,
	true,
}, {
	"item_coupon_codes.go",
	"CouponCodes",
	"CouponCode",
	false,
	false,
}, {
	"item_coupon_sets.go",
	"CouponSets",
	"CouponSet",
	false,
	true,
}, {
	"item_coupons.go",
	"Coupons",
	"Coupon",
	false,
	true,
}, {
	"item_differential_prices.go",
	"DifferentialPrices",
	"DifferentialPrice",
	false,
	true,
}, {
	"item_families.go",
	"ItemFamilies",
	"ItemFamily",
	false,
	true,
}, {
	"item_prices.go",
	"ItemPrices",
	"ItemPrice",
	false,
	true,
}, {
	"items.go",
	"Items",
	"Item",
	false,
	true,
}, {
	"orders.go",
	"Orders",
	"Order",
	false,
	true,
}, {
	"payment_sources.go",
	"PaymentSources",
	"PaymentSource",
	false,
	true,
}, {
	"payment_virtual_bank_accounts.go",
	"VirtualBankAccounts",
	"VirtualBankAccount",
	false,
	true,
}, {
	"payments.go",
	"Transactions",
	"Transaction",
	false,
	true,
}, {
	"quote_line_groups.go",
	"QuoteLineGroups",
	"QuoteLineGroup",
	true,
	false,
}, {
	"quotes.go",
	"Quotes",
	"Quote",
	false,
	true,
}, {
	"subscription_contract_terms.go",
	"ContractTerms",
	"ContractTerm",
	true,
	false,
}, {
	"subscription_discounts.go",
	"Discounts",
	"Discount",
	true,
	false,
}, {
	"subscription_entitlements.go",
	"SubscriptionEntitlements",
	"SubscriptionEntitlement",
	true,
	false,
}, {
	"subscription_gifts.go",
	"Gifts",
	"Gift",
	false,
	true,
}, {
	"subscription_invoice_schedules.go",
	"AdvancedInvoiceSchedules",
	"AdvancedInvoiceSchedule",
	true,
	false,
}, {
	"subscription_ramps.go",
	"Ramps",
	"Ramp",
	false,
	false,
}, {
	"subscription_usages.go",
	"Usages",
	"Usage",
	false,
	false,
}, {
	"subscriptions.go",
	"Subscriptions",
	"Subscription",
	false,
	false,
}}}

func GenerateListConvenienceMethods() string {
	output := strings.Builder{}
	err := listConvTemplate.Execute(&output, convenience)
	if err != nil {
		panic(err)
	}

	return output.String()
}
