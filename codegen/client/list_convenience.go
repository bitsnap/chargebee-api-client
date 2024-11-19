package client

import (
	"embed"
	"strings"

	. "github.com/bitsnap/go-util"
)

//go:embed list_convenience.go.tpl
var listConvTemplateFS embed.FS

var listConvTemplate = MustParseTemplate(listConvTemplateFS, "list_convenience.go.tpl")

func GenerateListConvenienceMethods() string {
	type Method struct {
		Name         string
		NameSingular string
		WithID       bool
	}

	type ListConvenience struct {
		Methods []Method
	}

	output := strings.Builder{}
	err := listConvTemplate.Execute(&output, ListConvenience{
		Methods: []Method{{
			"BusinessEntityTransfers",
			"BusinessEntityTransfer",
			false,
		}, {
			"Currencies",
			"Currency",
			false,
		}, {
			"CustomerContacts",
			"Contact",
			true,
		}, {
			"CustomerHierarchies",
			"Hierarchy",
			true,
		}, {
			"CustomerPaymentVouchers",
			"PaymentVoucher",
			true,
		}, {
			"Customers",
			"Customer",
			false,
		}, {
			"Entitlements",
			"Entitlement",
			false,
		}, {
			"Events",
			"Event",
			false,
		}, {
			"Features",
			"Feature",
			false,
		}, {
			"CreditNotes",
			"CreditNote",
			false,
		}, {
			"PaymentReferenceNumbers",
			"PaymentReferenceNumber",
			true,
		}, {
			"InvoicePayments",
			"Transaction",
			true,
		}, {
			"PromotionalCredits",
			"PromotionalCredit",
			false,
		}, {
			"InvoiceUnbilledCharges",
			"InvoiceUnbilledCharge",
			false,
		}, {
			"Invoices",
			"Invoice",
			false,
		}, {
			"ItemsAttached",
			"ItemAttached",
			true,
		}, {
			"CouponCodes",
			"CouponCode",
			false,
		}, {
			"CouponSets",
			"CouponSet",
			false,
		}, {
			"Coupons",
			"Coupon",
			false,
		}, {
			"DifferentialPrices",
			"DifferentialPrice",
			false,
		}, {
			"ItemFamilies",
			"ItemFamily",
			false,
		}, {
			"ItemPrices",
			"ItemPrice",
			false,
		}, {
			"Items",
			"Item",
			false,
		}, {
			"Orders",
			"Order",
			false,
		}, {
			"PaymentSources",
			"PaymentSource",
			false,
		}, {
			"VirtualBankAccounts",
			"VirtualBankAccount",
			false,
		}, {
			"Transactions",
			"Transaction",
			false,
		}, {
			"QuoteLineGroups",
			"QuoteLineGroup",
			true,
		}, {
			"Quotes",
			"Quote",
			false,
		}, {
			"ContractTerms",
			"ContractTerm",
			true,
		}, {
			"Discounts",
			"Discount",
			true,
		}, {
			"SubscriptionEntitlements",
			"SubscriptionEntitlement",
			true,
		}, {
			"Gifts",
			"Gift",
			false,
		}, {
			"AdvancedInvoiceSchedules",
			"AdvancedInvoiceSchedule",
			true,
		}, {
			"Ramps",
			"Ramp",
			false,
		}, {
			"Usages",
			"Usage",
			false,
		}, {
			"Subscriptions",
			"Subscription",
			false,
		}},
	})
	if err != nil {
		panic(err)
	}

	return output.String()
}
