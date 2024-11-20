package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"time"

)



func ListBusinessEntityTransfers(site string) ([]BusinessEntityTransfer, error) {
	return ListAll(site, DefaultSortBy(), ListBusinessEntityTransfersPageSortBy)
}

func ListBusinessEntityTransfersCreatedSince(site string, createdSince *time.Time) ([]BusinessEntityTransfer, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListBusinessEntityTransfersPageSortBy)
}

func ListBusinessEntityTransfersUpdatedSince(site string, updatedSince *time.Time) ([]BusinessEntityTransfer, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListBusinessEntityTransfersPageSortBy)
}


func ListCurrencies(site string) ([]Currency, error) {
	return ListAll(site, DefaultSortBy(), ListCurrenciesPageSortBy)
}

func ListCurrenciesCreatedSince(site string, createdSince *time.Time) ([]Currency, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListCurrenciesPageSortBy)
}

func ListCurrenciesUpdatedSince(site string, updatedSince *time.Time) ([]Currency, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListCurrenciesPageSortBy)
}


func ListCustomerContacts(site, id string) ([]Contact, error) {
	return ListAllOfId(site, id, DefaultSortBy(), ListCustomerContactsPageSortBy)
}

func ListCustomerContactsCreatedSince(site, id string, createdSince *time.Time) ([]Contact, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListCustomerContactsPageSortBy)
}

func ListCustomerContactsUpdatedSince(site, id string, updatedSince *time.Time) ([]Contact, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListCustomerContactsPageSortBy)
}


func ListCustomerHierarchies(site, id string) ([]Hierarchy, error) {
	return ListAllOfId(site, id, DefaultSortBy(), ListCustomerHierarchiesPageSortBy)
}

func ListCustomerHierarchiesCreatedSince(site, id string, createdSince *time.Time) ([]Hierarchy, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListCustomerHierarchiesPageSortBy)
}

func ListCustomerHierarchiesUpdatedSince(site, id string, updatedSince *time.Time) ([]Hierarchy, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListCustomerHierarchiesPageSortBy)
}


func ListCustomerPaymentVouchers(site, id string) ([]PaymentVoucher, error) {
	return ListAllOfId(site, id, DefaultSortBy(), ListCustomerPaymentVouchersPageSortBy)
}

func ListCustomerPaymentVouchersCreatedSince(site, id string, createdSince *time.Time) ([]PaymentVoucher, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListCustomerPaymentVouchersPageSortBy)
}

func ListCustomerPaymentVouchersUpdatedSince(site, id string, updatedSince *time.Time) ([]PaymentVoucher, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListCustomerPaymentVouchersPageSortBy)
}


func ListCustomers(site string) ([]Customer, error) {
	return ListAll(site, DefaultSortBy(), ListCustomersPageSortBy)
}

func ListCustomersCreatedSince(site string, createdSince *time.Time) ([]Customer, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListCustomersPageSortBy)
}

func ListCustomersUpdatedSince(site string, updatedSince *time.Time) ([]Customer, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListCustomersPageSortBy)
}


func ListEntitlements(site string) ([]Entitlement, error) {
	return ListAll(site, DefaultSortBy(), ListEntitlementsPageSortBy)
}

func ListEntitlementsCreatedSince(site string, createdSince *time.Time) ([]Entitlement, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListEntitlementsPageSortBy)
}

func ListEntitlementsUpdatedSince(site string, updatedSince *time.Time) ([]Entitlement, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListEntitlementsPageSortBy)
}


func ListEvents(site string) ([]Event, error) {
	return ListAll(site, DefaultSortBy(), ListEventsPageSortBy)
}

func ListEventsCreatedSince(site string, createdSince *time.Time) ([]Event, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListEventsPageSortBy)
}

func ListEventsUpdatedSince(site string, updatedSince *time.Time) ([]Event, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListEventsPageSortBy)
}


func ListFeatures(site string) ([]Feature, error) {
	return ListAll(site, DefaultSortBy(), ListFeaturesPageSortBy)
}

func ListFeaturesCreatedSince(site string, createdSince *time.Time) ([]Feature, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListFeaturesPageSortBy)
}

func ListFeaturesUpdatedSince(site string, updatedSince *time.Time) ([]Feature, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListFeaturesPageSortBy)
}


func ListCreditNotes(site string) ([]CreditNote, error) {
	return ListAll(site, DefaultSortBy(), ListCreditNotesPageSortBy)
}

func ListCreditNotesCreatedSince(site string, createdSince *time.Time) ([]CreditNote, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListCreditNotesPageSortBy)
}

func ListCreditNotesUpdatedSince(site string, updatedSince *time.Time) ([]CreditNote, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListCreditNotesPageSortBy)
}


func ListPaymentReferenceNumbers(site, id string) ([]PaymentReferenceNumber, error) {
	return ListAllOfId(site, id, DefaultSortBy(), ListPaymentReferenceNumbersPageSortBy)
}

func ListPaymentReferenceNumbersCreatedSince(site, id string, createdSince *time.Time) ([]PaymentReferenceNumber, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListPaymentReferenceNumbersPageSortBy)
}

func ListPaymentReferenceNumbersUpdatedSince(site, id string, updatedSince *time.Time) ([]PaymentReferenceNumber, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListPaymentReferenceNumbersPageSortBy)
}


func ListInvoicePayments(site, id string) ([]Transaction, error) {
	return ListAllOfId(site, id, DefaultSortBy(), ListInvoicePaymentsPageSortBy)
}

func ListInvoicePaymentsCreatedSince(site, id string, createdSince *time.Time) ([]Transaction, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListInvoicePaymentsPageSortBy)
}

func ListInvoicePaymentsUpdatedSince(site, id string, updatedSince *time.Time) ([]Transaction, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListInvoicePaymentsPageSortBy)
}


func ListPromotionalCredits(site string) ([]PromotionalCredit, error) {
	return ListAll(site, DefaultSortBy(), ListPromotionalCreditsPageSortBy)
}

func ListPromotionalCreditsCreatedSince(site string, createdSince *time.Time) ([]PromotionalCredit, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListPromotionalCreditsPageSortBy)
}

func ListPromotionalCreditsUpdatedSince(site string, updatedSince *time.Time) ([]PromotionalCredit, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListPromotionalCreditsPageSortBy)
}


func ListInvoiceUnbilledCharges(site string) ([]InvoiceUnbilledCharge, error) {
	return ListAll(site, DefaultSortBy(), ListInvoiceUnbilledChargesPageSortBy)
}

func ListInvoiceUnbilledChargesCreatedSince(site string, createdSince *time.Time) ([]InvoiceUnbilledCharge, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListInvoiceUnbilledChargesPageSortBy)
}

func ListInvoiceUnbilledChargesUpdatedSince(site string, updatedSince *time.Time) ([]InvoiceUnbilledCharge, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListInvoiceUnbilledChargesPageSortBy)
}


func ListInvoicePaymentVouchers(site, id string) ([]PaymentVoucher, error) {
	return ListAllOfId(site, id, DefaultSortBy(), ListInvoicePaymentVouchersPageSortBy)
}

func ListInvoicePaymentVouchersCreatedSince(site, id string, createdSince *time.Time) ([]PaymentVoucher, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListInvoicePaymentVouchersPageSortBy)
}

func ListInvoicePaymentVouchersUpdatedSince(site, id string, updatedSince *time.Time) ([]PaymentVoucher, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListInvoicePaymentVouchersPageSortBy)
}


func ListInvoices(site string) ([]Invoice, error) {
	return ListAll(site, DefaultSortBy(), ListInvoicesPageSortBy)
}

func ListInvoicesCreatedSince(site string, createdSince *time.Time) ([]Invoice, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListInvoicesPageSortBy)
}

func ListInvoicesUpdatedSince(site string, updatedSince *time.Time) ([]Invoice, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListInvoicesPageSortBy)
}


func ListItemsAttached(site, id string) ([]ItemAttached, error) {
	return ListAllOfId(site, id, DefaultSortBy(), ListItemsAttachedPageSortBy)
}

func ListItemsAttachedCreatedSince(site, id string, createdSince *time.Time) ([]ItemAttached, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListItemsAttachedPageSortBy)
}

func ListItemsAttachedUpdatedSince(site, id string, updatedSince *time.Time) ([]ItemAttached, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListItemsAttachedPageSortBy)
}


func ListCouponCodes(site string) ([]CouponCode, error) {
	return ListAll(site, DefaultSortBy(), ListCouponCodesPageSortBy)
}

func ListCouponCodesCreatedSince(site string, createdSince *time.Time) ([]CouponCode, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListCouponCodesPageSortBy)
}

func ListCouponCodesUpdatedSince(site string, updatedSince *time.Time) ([]CouponCode, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListCouponCodesPageSortBy)
}


func ListCouponSets(site string) ([]CouponSet, error) {
	return ListAll(site, DefaultSortBy(), ListCouponSetsPageSortBy)
}

func ListCouponSetsCreatedSince(site string, createdSince *time.Time) ([]CouponSet, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListCouponSetsPageSortBy)
}

func ListCouponSetsUpdatedSince(site string, updatedSince *time.Time) ([]CouponSet, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListCouponSetsPageSortBy)
}


func ListCoupons(site string) ([]Coupon, error) {
	return ListAll(site, DefaultSortBy(), ListCouponsPageSortBy)
}

func ListCouponsCreatedSince(site string, createdSince *time.Time) ([]Coupon, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListCouponsPageSortBy)
}

func ListCouponsUpdatedSince(site string, updatedSince *time.Time) ([]Coupon, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListCouponsPageSortBy)
}


func ListDifferentialPrices(site string) ([]DifferentialPrice, error) {
	return ListAll(site, DefaultSortBy(), ListDifferentialPricesPageSortBy)
}

func ListDifferentialPricesCreatedSince(site string, createdSince *time.Time) ([]DifferentialPrice, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListDifferentialPricesPageSortBy)
}

func ListDifferentialPricesUpdatedSince(site string, updatedSince *time.Time) ([]DifferentialPrice, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListDifferentialPricesPageSortBy)
}


func ListItemFamilies(site string) ([]ItemFamily, error) {
	return ListAll(site, DefaultSortBy(), ListItemFamiliesPageSortBy)
}

func ListItemFamiliesCreatedSince(site string, createdSince *time.Time) ([]ItemFamily, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListItemFamiliesPageSortBy)
}

func ListItemFamiliesUpdatedSince(site string, updatedSince *time.Time) ([]ItemFamily, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListItemFamiliesPageSortBy)
}


func ListItemPrices(site string) ([]ItemPrice, error) {
	return ListAll(site, DefaultSortBy(), ListItemPricesPageSortBy)
}

func ListItemPricesCreatedSince(site string, createdSince *time.Time) ([]ItemPrice, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListItemPricesPageSortBy)
}

func ListItemPricesUpdatedSince(site string, updatedSince *time.Time) ([]ItemPrice, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListItemPricesPageSortBy)
}


func ListItems(site string) ([]Item, error) {
	return ListAll(site, DefaultSortBy(), ListItemsPageSortBy)
}

func ListItemsCreatedSince(site string, createdSince *time.Time) ([]Item, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListItemsPageSortBy)
}

func ListItemsUpdatedSince(site string, updatedSince *time.Time) ([]Item, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListItemsPageSortBy)
}


func ListOrders(site string) ([]Order, error) {
	return ListAll(site, DefaultSortBy(), ListOrdersPageSortBy)
}

func ListOrdersCreatedSince(site string, createdSince *time.Time) ([]Order, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListOrdersPageSortBy)
}

func ListOrdersUpdatedSince(site string, updatedSince *time.Time) ([]Order, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListOrdersPageSortBy)
}


func ListPaymentSources(site string) ([]PaymentSource, error) {
	return ListAll(site, DefaultSortBy(), ListPaymentSourcesPageSortBy)
}

func ListPaymentSourcesCreatedSince(site string, createdSince *time.Time) ([]PaymentSource, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListPaymentSourcesPageSortBy)
}

func ListPaymentSourcesUpdatedSince(site string, updatedSince *time.Time) ([]PaymentSource, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListPaymentSourcesPageSortBy)
}


func ListVirtualBankAccounts(site string) ([]VirtualBankAccount, error) {
	return ListAll(site, DefaultSortBy(), ListVirtualBankAccountsPageSortBy)
}

func ListVirtualBankAccountsCreatedSince(site string, createdSince *time.Time) ([]VirtualBankAccount, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListVirtualBankAccountsPageSortBy)
}

func ListVirtualBankAccountsUpdatedSince(site string, updatedSince *time.Time) ([]VirtualBankAccount, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListVirtualBankAccountsPageSortBy)
}


func ListTransactions(site string) ([]Transaction, error) {
	return ListAll(site, DefaultSortBy(), ListTransactionsPageSortBy)
}

func ListTransactionsCreatedSince(site string, createdSince *time.Time) ([]Transaction, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListTransactionsPageSortBy)
}

func ListTransactionsUpdatedSince(site string, updatedSince *time.Time) ([]Transaction, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListTransactionsPageSortBy)
}


func ListQuoteLineGroups(site, id string) ([]QuoteLineGroup, error) {
	return ListAllOfId(site, id, DefaultSortBy(), ListQuoteLineGroupsPageSortBy)
}

func ListQuoteLineGroupsCreatedSince(site, id string, createdSince *time.Time) ([]QuoteLineGroup, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListQuoteLineGroupsPageSortBy)
}

func ListQuoteLineGroupsUpdatedSince(site, id string, updatedSince *time.Time) ([]QuoteLineGroup, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListQuoteLineGroupsPageSortBy)
}


func ListQuotes(site string) ([]Quote, error) {
	return ListAll(site, DefaultSortBy(), ListQuotesPageSortBy)
}

func ListQuotesCreatedSince(site string, createdSince *time.Time) ([]Quote, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListQuotesPageSortBy)
}

func ListQuotesUpdatedSince(site string, updatedSince *time.Time) ([]Quote, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListQuotesPageSortBy)
}


func ListContractTerms(site, id string) ([]ContractTerm, error) {
	return ListAllOfId(site, id, DefaultSortBy(), ListContractTermsPageSortBy)
}

func ListContractTermsCreatedSince(site, id string, createdSince *time.Time) ([]ContractTerm, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListContractTermsPageSortBy)
}

func ListContractTermsUpdatedSince(site, id string, updatedSince *time.Time) ([]ContractTerm, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListContractTermsPageSortBy)
}


func ListDiscounts(site, id string) ([]Discount, error) {
	return ListAllOfId(site, id, DefaultSortBy(), ListDiscountsPageSortBy)
}

func ListDiscountsCreatedSince(site, id string, createdSince *time.Time) ([]Discount, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListDiscountsPageSortBy)
}

func ListDiscountsUpdatedSince(site, id string, updatedSince *time.Time) ([]Discount, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListDiscountsPageSortBy)
}


func ListSubscriptionEntitlements(site, id string) ([]SubscriptionEntitlement, error) {
	return ListAllOfId(site, id, DefaultSortBy(), ListSubscriptionEntitlementsPageSortBy)
}

func ListSubscriptionEntitlementsCreatedSince(site, id string, createdSince *time.Time) ([]SubscriptionEntitlement, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListSubscriptionEntitlementsPageSortBy)
}

func ListSubscriptionEntitlementsUpdatedSince(site, id string, updatedSince *time.Time) ([]SubscriptionEntitlement, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListSubscriptionEntitlementsPageSortBy)
}


func ListGifts(site string) ([]Gift, error) {
	return ListAll(site, DefaultSortBy(), ListGiftsPageSortBy)
}

func ListGiftsCreatedSince(site string, createdSince *time.Time) ([]Gift, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListGiftsPageSortBy)
}

func ListGiftsUpdatedSince(site string, updatedSince *time.Time) ([]Gift, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListGiftsPageSortBy)
}


func ListAdvancedInvoiceSchedules(site, id string) ([]AdvancedInvoiceSchedule, error) {
	return ListAllOfId(site, id, DefaultSortBy(), ListAdvancedInvoiceSchedulesPageSortBy)
}

func ListAdvancedInvoiceSchedulesCreatedSince(site, id string, createdSince *time.Time) ([]AdvancedInvoiceSchedule, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListAdvancedInvoiceSchedulesPageSortBy)
}

func ListAdvancedInvoiceSchedulesUpdatedSince(site, id string, updatedSince *time.Time) ([]AdvancedInvoiceSchedule, error) {
	return ListAllOfId(site, id, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListAdvancedInvoiceSchedulesPageSortBy)
}


func ListRamps(site string) ([]Ramp, error) {
	return ListAll(site, DefaultSortBy(), ListRampsPageSortBy)
}

func ListRampsCreatedSince(site string, createdSince *time.Time) ([]Ramp, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListRampsPageSortBy)
}

func ListRampsUpdatedSince(site string, updatedSince *time.Time) ([]Ramp, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListRampsPageSortBy)
}


func ListUsages(site string) ([]Usage, error) {
	return ListAll(site, DefaultSortBy(), ListUsagesPageSortBy)
}

func ListUsagesCreatedSince(site string, createdSince *time.Time) ([]Usage, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListUsagesPageSortBy)
}

func ListUsagesUpdatedSince(site string, updatedSince *time.Time) ([]Usage, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListUsagesPageSortBy)
}


func ListSubscriptions(site string) ([]Subscription, error) {
	return ListAll(site, DefaultSortBy(), ListSubscriptionsPageSortBy)
}

func ListSubscriptionsCreatedSince(site string, createdSince *time.Time) ([]Subscription, error) {
	return ListAll(site, &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, ListSubscriptionsPageSortBy)
}

func ListSubscriptionsUpdatedSince(site string, updatedSince *time.Time) ([]Subscription, error) {
	return ListAll(site, &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, ListSubscriptionsPageSortBy)
}