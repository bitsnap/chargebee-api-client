package codegen

import (
	"strings"

	"github.com/bitsnap/chargebee-api-client/codegen/chargebee"
	"github.com/bitsnap/chargebee-api-client/codegen/chargebee/customer"
	"github.com/bitsnap/chargebee-api-client/codegen/chargebee/events"
	"github.com/bitsnap/chargebee-api-client/codegen/chargebee/invoices"
	"github.com/bitsnap/chargebee-api-client/codegen/chargebee/items"
	"github.com/bitsnap/chargebee-api-client/codegen/chargebee/payments"
	"github.com/bitsnap/chargebee-api-client/codegen/chargebee/quotes"
	"github.com/bitsnap/chargebee-api-client/codegen/chargebee/subscriptions"
	"github.com/bitsnap/chargebee-api-client/codegen/client"

	goutil "github.com/bitsnap/go-util"
)

var content = map[string][]func() string{
	"business_entities.go": {
		chargebee.GenerateBusinessEntitiesModel,
		chargebee.GenerateBusinessEntityTransferModel,
		chargebee.GenerateListBusinessEntityTransfers,
	},
	"contacts.go": {
		chargebee.GenerateContactsModel,
	},
	"currencies.go": {
		chargebee.GenerateCurrenciesModel,
		chargebee.GenerateListCurrencies,
		chargebee.GenerateRetrieveCurrency,
	},
	"customers.go": {
		chargebee.GenerateCustomersModel,
		chargebee.GenerateListCustomers,
		chargebee.GenerateRetrieveCustomer,
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
		chargebee.GenerateEntitlementsModel,
		chargebee.GenerateListEntitlements,
	},
	"events.go": {
		chargebee.GenerateEventsModel,
		chargebee.GenerateListEvents,
		chargebee.GenerateRetrieveEvent,
	},
	"events_items.go": {
		events.GenerateImpactedItemsModel,
	},
	"events_subscriptions.go": {
		events.GenerateImpactedSubscriptionsModel,
	},
	"features.go": {
		chargebee.GenerateFeaturesModel,
		chargebee.GenerateListFeatures,
		chargebee.GenerateRetrieveFeature,
	},
	"invoices.go": {
		chargebee.GenerateInvoicesModel,
		chargebee.GenerateListInvoices,
		chargebee.GenerateRetrieveInvoice,
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
		chargebee.GenerateItemsModel,
		chargebee.GenerateListItems,
		chargebee.GenerateRetrieveItem,
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
		chargebee.GenerateOrdersModel,
		chargebee.GenerateListOrders,
		chargebee.GenerateRetrieveOrder,
	},
	"payments.go": {
		chargebee.GenerateTransactionsModel,
		chargebee.GenerateListTransactions,
		chargebee.GenerateRetrieveTransaction,
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
		chargebee.GenerateQuotesModel,
		chargebee.GenerateListQuotes,
		chargebee.GenerateRetrieveQuote,
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
		chargebee.GenerateSubscriptionsModel,
		chargebee.GenerateListSubscriptions,
		chargebee.GenerateRetrieveSubscription,
		chargebee.GenerateRetrieveSubscriptionWithScheduledChanges,
	},
	// TODO: Subscription address takes both subscription_id and address label parameters
	//"subscription_address.go": {
	//subscriptions.GenerateSubscriptionAddressesModel,
	//subscriptions.GenerateRetrieveSubscriptionAddress,
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
	//"subscription_purchases.go": {
	//	subscriptions.GenerateSubscriptionPurchasesModel,
	//},
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
		client.GenerateListConvenienceMethods,
	},
}

func header(content string) string {
	enums := "\t\"github.com/bitsnap/chargebee-api-client/chargebee/enums\"\n"
	models := "\t\"github.com/bitsnap/chargebee-api-client/chargebee/models\"\n"
	decimal := "\t\"github.com/shopspring/decimal\"\n"
	time := "\t\"time\"\n"

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

	return `package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
` + enums + models + decimal + time + `
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

`
}

func GenerateInto(targetDir string) {
	goutil.GenerateInto(targetDir, header, content)
}
