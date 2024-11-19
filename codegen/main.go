package main

import (
	"os"
	"path"
	"strings"

	"github.com/bitsnap/chargebee-api-client/codegen/templates"

	"github.com/bitsnap/chargebee-api-client/codegen/client"
	"github.com/bitsnap/chargebee-api-client/codegen/client/customer"
	"github.com/bitsnap/chargebee-api-client/codegen/client/events"
	"github.com/bitsnap/chargebee-api-client/codegen/client/invoices"
	"github.com/bitsnap/chargebee-api-client/codegen/client/items"
	"github.com/bitsnap/chargebee-api-client/codegen/client/payments"
	"github.com/bitsnap/chargebee-api-client/codegen/client/quotes"
	"github.com/bitsnap/chargebee-api-client/codegen/client/subscriptions"

	. "github.com/bitsnap/go-util"
)

var content = map[string][]func() string{
	"business_entities.go": {
		client.GenerateBusinessEntitiesModel,
		client.GenerateBusinessEntityTransferModel,
		client.GenerateListBusinessEntityTransfers,
	},
	"contacts.go": {
		client.GenerateContactsModel,
	},
	"currencies.go": {
		client.GenerateCurrenciesModel,
		client.GenerateListCurrencies,
		client.GenerateRetrieveCurrency,
	},
	"customers.go": {
		client.GenerateCustomersModel,
		client.GenerateListCustomers,
		client.GenerateRetrieveCustomer,
	},
	"customer_contacts.go": {
		customer.GenerateListCustomerContacts,
	},
	"customer_hierarchies.go": {
		customer.GenerateRetrieveCustomerAccountHierarchy,
		customer.GenerateHierarchyModel,
	},
	"customer_vouchers.go": {
		customer.GenerateListCustomerVouchers,
	},
	"entitlements.go": {
		client.GenerateEntitlementsModel,
		client.GenerateListEntitlements,
	},
	"events.go": {
		client.GenerateEventsModel,
		client.GenerateListEvents,
		client.GenerateRetrieveEvent,
	},
	"events_items.go": {
		events.GenerateImpactedItemsModel,
	},
	"events_subscriptions.go": {
		events.GenerateImpactedSubscriptionsModel,
	},
	"features.go": {
		client.GenerateFeaturesModel,
		client.GenerateListFeatures,
		client.GenerateRetrieveFeature,
	},
	"invoices.go": {
		client.GenerateInvoicesModel,
		client.GenerateListInvoices,
		client.GenerateRetrieveInvoice,
	},
	"invoice_credit_notes.go": {
		invoices.GenerateInvoiceCreditNoteModel,
		invoices.GenerateListInvoiceCreditNotes,
		invoices.GenerateRetrieveInvoiceCreditNote,
	},
	"invoice_payment_reference_numbers.go": {
		invoices.GenerateInvoicePaymentReferenceNumberModel,
		invoices.GenerateListInvoicePaymentReferenceNumbers,
	},
	"invoice_payments.go": {
		invoices.GenerateListInvoicePayments,
	},
	"invoice_promotional_credits.go": {
		invoices.GenerateInvoicePromotionalCreditModel,
		invoices.GenerateListInvoicePromotionalCredits,
		invoices.GenerateRetrieveInvoicePromotionalCredit,
	},
	"invoice_un_billed_charges.go": {
		invoices.GenerateInvoiceUnBilledChargeModel,
		invoices.GenerateListInvoiceUnBilledCharges,
	},
	"invoice_vouchers.go": {
		invoices.GenerateListInvoiceVouchers,
	},
	"items.go": {
		client.GenerateItemsModel,
		client.GenerateListItems,
		client.GenerateRetrieveItem,
	},
	"item_coupons.go": {
		items.GenerateItemCouponModel,
		items.GenerateListItemCoupons,
		items.GenerateRetrieveItemCoupon,
	},
	"item_coupon_codes.go": {
		items.GenerateItemCouponCodeModel,
		items.GenerateListItemCouponCodes,
		items.GenerateRetrieveItemCouponCode,
	},
	"item_coupon_sets.go": {
		items.GenerateItemCouponSetModel,
		items.GenerateListItemCouponSets,
		items.GenerateRetrieveItemCouponSet,
	},
	"item_differential_prices.go": {
		items.GenerateItemDifferentialPriceModel,
		items.GenerateListItemDifferentialPrices,
		items.GenerateRetrieveItemDifferentialPrice,
	},
	"item_families.go": {
		items.GenerateItemFamilyModel,
		items.GenerateListItemFamilies,
		items.GenerateRetrieveItemFamily,
	},
	"item_prices.go": {
		items.GenerateItemPriceModel,
		items.GenerateListItemPrices,
		items.GenerateRetrieveItemPrice,
	},
	"item_attached.go": {
		items.GenerateItemAttachedModel,
		items.GenerateListItemAttached,
		items.GenerateRetrieveItemAttached,
	},
	"orders.go": {
		client.GenerateOrdersModel,
		client.GenerateListOrders,
		client.GenerateRetrieveOrder,
	},
	"payments.go": {
		client.GenerateTransactionsModel,
		client.GenerateListTransactions,
		client.GenerateRetrieveTransaction,
	},
	"payment_sources.go": {
		payments.GeneratePaymentSourcesModel,
		payments.GenerateListPaymentSources,
		payments.GenerateRetrievePaymentSource,
	},
	"payment_virtual_bank_accounts.go": {
		payments.GenerateVirtualBankAccountsModel,
		payments.GenerateListVirtualBankAccounts,
		payments.GenerateRetrieveVirtualBankAccount,
	},
	"payment_vouchers.go": {
		payments.GenerateVouchersModel,
	},
	"quotes.go": {
		client.GenerateQuotesModel,
		client.GenerateListQuotes,
		client.GenerateRetrieveQuote,
	},
	"quote_line_groups.go": {
		quotes.GenerateListQuoteLineGroups,
		quotes.GenerateListAllQuoteLineGroups,
		quotes.GenerateRetrieveQuoteLineGroup,
	},
	"quoted_charges.go": {
		quotes.GenerateQuotedChargeModel,
	},
	"quoted_subscription.go": {
		quotes.GenerateQuotedSubscriptionModel,
	},
	"subscriptions.go": {
		client.GenerateSubscriptionsModel,
		client.GenerateListSubscriptions,
		client.GenerateRetrieveSubscription,
		client.GenerateRetrieveSubscriptionWithScheduledChanges,
	},
	// TODO: Subscription address takes both subscription_id and address label parameters
	// "subscription_address.go": {
	// subscriptions.GenerateSubscriptionAddressesModel,
	// subscriptions.GenerateRetrieveSubscriptionAddress,
	//},
	"subscription_contract_terms.go": {
		subscriptions.GenerateSubscriptionContractTermsModel,
		subscriptions.GenerateListSubscriptionContractTerms,
	},
	"subscription_discounts.go": {
		subscriptions.GenerateSubscriptionDiscountsModel,
		subscriptions.GenerateListSubscriptionDiscounts,
	},
	"subscription_entitlements.go": {
		subscriptions.GenerateSubscriptionEntitlementsModel,
		subscriptions.GenerateListSubscriptionEntitlements,
	},
	"subscription_gifts.go": {
		subscriptions.GenerateSubscriptionGiftsModel,
		subscriptions.GenerateListSubscriptionGifts,
		subscriptions.GenerateRetrieveSubscriptionGift,
	},
	"subscription_invoice_schedules.go": {
		subscriptions.GenerateSubscriptionAdvancedInvoiceScheduleModel,
		subscriptions.GenerateListSubscriptionAdvancedInvoiceSchedule,
	},
	// TODO: alpha API endpoint, there's not much to do
	// "subscription_purchases.go": {
	//	subscriptions.GenerateSubscriptionPurchasesModel,
	// },
	"subscription_ramps.go": {
		subscriptions.GenerateSubscriptionRampsModel,
		subscriptions.GenerateListRamps,
	},
	"subscription_usages.go": {
		subscriptions.GenerateSubscriptionUsagesModel,
		subscriptions.GenerateListSubscriptionUsages,
		subscriptions.GenerateRetrieveSubscriptionUsage,
	},
	"list_convenience.go": {
		templates.GenerateListConvenienceMethods,
	},
}

func header(content string) string {
	common := "\t. \"github.com/bitsnap/chargebee-api-client/chargebee/client\"\n"
	enums := "\t\"github.com/bitsnap/chargebee-api-client/chargebee/enums\"\n"
	models := "\t\"github.com/bitsnap/chargebee-api-client/chargebee/models\"\n"
	decimal := "\t\"github.com/shopspring/decimal\"\n"
	time := "\t\"time\"\n"

	imports := "\t\"github.com/goccy/go-json\"\n\t\"github.com/gofiber/fiber/v2\"\n\t\"net/url\""

	if !strings.Contains(content, "GetQuery") && !strings.Contains(content, "SortBy") {
		common = ""
	}

	if !strings.Contains(content, "models.") {
		models = ""
	}

	if !strings.Contains(content, "enums.") {
		enums = ""
	}

	if !strings.Contains(content, "time.") {
		time = ""
	}

	if !strings.Contains(content, "decimal.") {
		decimal = ""
	}

	if !strings.Contains(content, "fiber") {
		imports = ""
	}

	return `package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
` + common + enums + models + decimal + time + imports + `
)

`
}

func main() {
	rootDir, ok := os.LookupEnv("PWD")
	if !ok || len(os.Args) == 1 {
		panic("missing $PWD env variable and root dir")
	}

	rootDir = path.Join(rootDir, os.Args[1])
	targetDir := path.Join(rootDir, "chargebee", "generated")

	syncLogger := SetGlobalLogger()
	defer syncLogger()

	Logger().Infow("starting bitsnap chargebee codegen",
		"rootDir", rootDir,
		"targetDir", targetDir,
	)

	GenerateInto(rootDir, func(_ string) string {
		return ""
	}, templates.GeneratePublicAPIContent())

	GenerateInto(targetDir, header, content)
}
