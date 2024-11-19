package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"time"

	"github.com/bitsnap/chargebee-api-client/chargebee/client"
	chargebee "github.com/bitsnap/chargebee-api-client/chargebee/generated"
)

// UseSite sets custom Chargebee site name overriding CHARGEBEE_SITE environment variable value
func UseSite(site string) {
	client.UseSite(site)
}

// UseToken sets custom Chargebee API token, overriding CHARGEBEE_API_TOKEN environment variable value
func UseToken(token string) {
	client.UseToken(token)
}

// business_entities.go

func ListAllBusinessEntityTransfers() ([]chargebee.BusinessEntityTransfer, error) {
	return ListAllSiteBusinessEntityTransfers(client.Site())
}

func ListAllSiteBusinessEntityTransfers(site string) ([]chargebee.BusinessEntityTransfer, error) {
	return chargebee.ListBusinessEntityTransfers(site)
}

func ListBusinessEntityTransfersCreatedSince(createdSince *time.Time) ([]chargebee.BusinessEntityTransfer, error) {
	return ListSiteBusinessEntityTransfersCreatedSince(client.Site(), createdSince)
}

func ListSiteBusinessEntityTransfersCreatedSince(site string, createdSince *time.Time) ([]chargebee.BusinessEntityTransfer, error) {
	return chargebee.ListBusinessEntityTransfersCreatedSince(site, createdSince)
}

func ListBusinessEntityTransfersUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.BusinessEntityTransfer, error) {
	return ListSiteBusinessEntityTransfersUpdatedSince(client.Site(), updatedSince)
}

func ListSiteBusinessEntityTransfersUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.BusinessEntityTransfer, error) {
	return chargebee.ListBusinessEntityTransfersUpdatedSince(site, updatedSince)
}

func ListBusinessEntityTransfers() ([]chargebee.BusinessEntityTransfer, string, error) {
	return ListSiteBusinessEntityTransfers(client.Site())
}

func ListSiteBusinessEntityTransfers(site string) ([]chargebee.BusinessEntityTransfer, string, error) {
	return chargebee.ListBusinessEntityTransfersPage(site, "")
}

func ListBusinessEntityTransfersPage(offset string) ([]chargebee.BusinessEntityTransfer, string, error) {
	return ListSiteBusinessEntityTransfersPage(client.Site(), offset)
}

func ListSiteBusinessEntityTransfersPage(site string, offset string) ([]chargebee.BusinessEntityTransfer, string, error) {
	return chargebee.ListBusinessEntityTransfersPage(site, offset)
}

func ListBusinessEntityTransfersPageSortBy(offset string, field string) ([]chargebee.BusinessEntityTransfer, string, error) {
	return ListSiteBusinessEntityTransfersPageSortBy(client.Site(), offset, field)
}

func ListSiteBusinessEntityTransfersPageSortBy(site string, offset string, field string) ([]chargebee.BusinessEntityTransfer, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListBusinessEntityTransfersPageSortBy(site, offset, sortBy)
}

// currencies.go

func ListAllCurrencies() ([]chargebee.Currency, error) {
	return ListAllSiteCurrencies(client.Site())
}

func ListAllSiteCurrencies(site string) ([]chargebee.Currency, error) {
	return chargebee.ListCurrencies(site)
}

func ListCurrenciesCreatedSince(createdSince *time.Time) ([]chargebee.Currency, error) {
	return ListSiteCurrenciesCreatedSince(client.Site(), createdSince)
}

func ListSiteCurrenciesCreatedSince(site string, createdSince *time.Time) ([]chargebee.Currency, error) {
	return chargebee.ListCurrenciesCreatedSince(site, createdSince)
}

func ListCurrenciesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Currency, error) {
	return ListSiteCurrenciesUpdatedSince(client.Site(), updatedSince)
}

func ListSiteCurrenciesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Currency, error) {
	return chargebee.ListCurrenciesUpdatedSince(site, updatedSince)
}

func ListCurrencies() ([]chargebee.Currency, string, error) {
	return ListSiteCurrencies(client.Site())
}

func ListSiteCurrencies(site string) ([]chargebee.Currency, string, error) {
	return chargebee.ListCurrenciesPage(site, "")
}

func ListCurrenciesPage(offset string) ([]chargebee.Currency, string, error) {
	return ListSiteCurrenciesPage(client.Site(), offset)
}

func ListSiteCurrenciesPage(site string, offset string) ([]chargebee.Currency, string, error) {
	return chargebee.ListCurrenciesPage(site, offset)
}

func ListCurrenciesPageSortBy(offset string, field string) ([]chargebee.Currency, string, error) {
	return ListSiteCurrenciesPageSortBy(client.Site(), offset, field)
}

func ListSiteCurrenciesPageSortBy(site string, offset string, field string) ([]chargebee.Currency, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListCurrenciesPageSortBy(site, offset, sortBy)
}

func RetrieveCurrencyById(id string) (*chargebee.Currency, error) {
	return RetrieveSiteCurrencyById(client.Site(), id)
}

func RetrieveSiteCurrencyById(site string, id string) (*chargebee.Currency, error) {
	return chargebee.RetrieveCurrencyById(site, id)
}

// customer_contacts.go

func ListAllCustomerContacts(id string) ([]chargebee.Contact, error) {
	return ListAllSiteCustomerContacts(client.Site(), id)
}

func ListAllSiteCustomerContacts(site, id string) ([]chargebee.Contact, error) {
	return chargebee.ListCustomerContacts(site, id)
}

func ListCustomerContactsCreatedSince(id string, createdSince *time.Time) ([]chargebee.Contact, error) {
	return ListSiteCustomerContactsCreatedSince(client.Site(), id, createdSince)
}

func ListSiteCustomerContactsCreatedSince(site, id string, createdSince *time.Time) ([]chargebee.Contact, error) {
	return chargebee.ListCustomerContactsCreatedSince(site, id, createdSince)
}

func ListCustomerContactsUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.Contact, error) {
	return ListSiteCustomerContactsUpdatedSince(client.Site(), id, updatedSince)
}

func ListSiteCustomerContactsUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.Contact, error) {
	return chargebee.ListCustomerContactsUpdatedSince(site, id, updatedSince)
}

func ListCustomerContacts(id string) ([]chargebee.Contact, string, error) {
	return ListSiteCustomerContacts(client.Site(), id)
}

func ListSiteCustomerContacts(site, id string) ([]chargebee.Contact, string, error) {
	return chargebee.ListCustomerContactsPage(site, id, "")
}

func ListCustomerContactsPage(id, offset string) ([]chargebee.Contact, string, error) {
	return ListSiteCustomerContactsPage(client.Site(), id, offset)
}

func ListSiteCustomerContactsPage(site, id string, offset string) ([]chargebee.Contact, string, error) {
	return chargebee.ListCustomerContactsPage(site, id, offset)
}

func ListCustomerContactsPageSortBy(id, offset string, field string) ([]chargebee.Contact, string, error) {
	return ListSiteCustomerContactsPageSortBy(client.Site(), id, offset, field)
}

func ListSiteCustomerContactsPageSortBy(site, id string, offset string, field string) ([]chargebee.Contact, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListCustomerContactsPageSortBy(site, id, offset, sortBy)
}

// customer_hierarchies.go

func ListAllCustomerHierarchies(id string) ([]chargebee.Hierarchy, error) {
	return ListAllSiteCustomerHierarchies(client.Site(), id)
}

func ListAllSiteCustomerHierarchies(site, id string) ([]chargebee.Hierarchy, error) {
	return chargebee.ListCustomerHierarchies(site, id)
}

func ListCustomerHierarchiesCreatedSince(id string, createdSince *time.Time) ([]chargebee.Hierarchy, error) {
	return ListSiteCustomerHierarchiesCreatedSince(client.Site(), id, createdSince)
}

func ListSiteCustomerHierarchiesCreatedSince(site, id string, createdSince *time.Time) ([]chargebee.Hierarchy, error) {
	return chargebee.ListCustomerHierarchiesCreatedSince(site, id, createdSince)
}

func ListCustomerHierarchiesUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.Hierarchy, error) {
	return ListSiteCustomerHierarchiesUpdatedSince(client.Site(), id, updatedSince)
}

func ListSiteCustomerHierarchiesUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.Hierarchy, error) {
	return chargebee.ListCustomerHierarchiesUpdatedSince(site, id, updatedSince)
}

func ListCustomerHierarchies(id string) ([]chargebee.Hierarchy, string, error) {
	return ListSiteCustomerHierarchies(client.Site(), id)
}

func ListSiteCustomerHierarchies(site, id string) ([]chargebee.Hierarchy, string, error) {
	return chargebee.ListCustomerHierarchiesPage(site, id, "")
}

func ListCustomerHierarchiesPage(id, offset string) ([]chargebee.Hierarchy, string, error) {
	return ListSiteCustomerHierarchiesPage(client.Site(), id, offset)
}

func ListSiteCustomerHierarchiesPage(site, id string, offset string) ([]chargebee.Hierarchy, string, error) {
	return chargebee.ListCustomerHierarchiesPage(site, id, offset)
}

func ListCustomerHierarchiesPageSortBy(id, offset string, field string) ([]chargebee.Hierarchy, string, error) {
	return ListSiteCustomerHierarchiesPageSortBy(client.Site(), id, offset, field)
}

func ListSiteCustomerHierarchiesPageSortBy(site, id string, offset string, field string) ([]chargebee.Hierarchy, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListCustomerHierarchiesPageSortBy(site, id, offset, sortBy)
}

// customer_vouchers.go

func ListAllCustomerPaymentVouchers(id string) ([]chargebee.PaymentVoucher, error) {
	return ListAllSiteCustomerPaymentVouchers(client.Site(), id)
}

func ListAllSiteCustomerPaymentVouchers(site, id string) ([]chargebee.PaymentVoucher, error) {
	return chargebee.ListCustomerPaymentVouchers(site, id)
}

func ListCustomerPaymentVouchersCreatedSince(id string, createdSince *time.Time) ([]chargebee.PaymentVoucher, error) {
	return ListSiteCustomerPaymentVouchersCreatedSince(client.Site(), id, createdSince)
}

func ListSiteCustomerPaymentVouchersCreatedSince(site, id string, createdSince *time.Time) ([]chargebee.PaymentVoucher, error) {
	return chargebee.ListCustomerPaymentVouchersCreatedSince(site, id, createdSince)
}

func ListCustomerPaymentVouchersUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.PaymentVoucher, error) {
	return ListSiteCustomerPaymentVouchersUpdatedSince(client.Site(), id, updatedSince)
}

func ListSiteCustomerPaymentVouchersUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.PaymentVoucher, error) {
	return chargebee.ListCustomerPaymentVouchersUpdatedSince(site, id, updatedSince)
}

func ListCustomerPaymentVouchers(id string) ([]chargebee.PaymentVoucher, string, error) {
	return ListSiteCustomerPaymentVouchers(client.Site(), id)
}

func ListSiteCustomerPaymentVouchers(site, id string) ([]chargebee.PaymentVoucher, string, error) {
	return chargebee.ListCustomerPaymentVouchersPage(site, id, "")
}

func ListCustomerPaymentVouchersPage(id, offset string) ([]chargebee.PaymentVoucher, string, error) {
	return ListSiteCustomerPaymentVouchersPage(client.Site(), id, offset)
}

func ListSiteCustomerPaymentVouchersPage(site, id string, offset string) ([]chargebee.PaymentVoucher, string, error) {
	return chargebee.ListCustomerPaymentVouchersPage(site, id, offset)
}

func ListCustomerPaymentVouchersPageSortBy(id, offset string, field string) ([]chargebee.PaymentVoucher, string, error) {
	return ListSiteCustomerPaymentVouchersPageSortBy(client.Site(), id, offset, field)
}

func ListSiteCustomerPaymentVouchersPageSortBy(site, id string, offset string, field string) ([]chargebee.PaymentVoucher, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListCustomerPaymentVouchersPageSortBy(site, id, offset, sortBy)
}

// customers.go

func ListAllCustomers() ([]chargebee.Customer, error) {
	return ListAllSiteCustomers(client.Site())
}

func ListAllSiteCustomers(site string) ([]chargebee.Customer, error) {
	return chargebee.ListCustomers(site)
}

func ListCustomersCreatedSince(createdSince *time.Time) ([]chargebee.Customer, error) {
	return ListSiteCustomersCreatedSince(client.Site(), createdSince)
}

func ListSiteCustomersCreatedSince(site string, createdSince *time.Time) ([]chargebee.Customer, error) {
	return chargebee.ListCustomersCreatedSince(site, createdSince)
}

func ListCustomersUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Customer, error) {
	return ListSiteCustomersUpdatedSince(client.Site(), updatedSince)
}

func ListSiteCustomersUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Customer, error) {
	return chargebee.ListCustomersUpdatedSince(site, updatedSince)
}

func ListCustomers() ([]chargebee.Customer, string, error) {
	return ListSiteCustomers(client.Site())
}

func ListSiteCustomers(site string) ([]chargebee.Customer, string, error) {
	return chargebee.ListCustomersPage(site, "")
}

func ListCustomersPage(offset string) ([]chargebee.Customer, string, error) {
	return ListSiteCustomersPage(client.Site(), offset)
}

func ListSiteCustomersPage(site string, offset string) ([]chargebee.Customer, string, error) {
	return chargebee.ListCustomersPage(site, offset)
}

func ListCustomersPageSortBy(offset string, field string) ([]chargebee.Customer, string, error) {
	return ListSiteCustomersPageSortBy(client.Site(), offset, field)
}

func ListSiteCustomersPageSortBy(site string, offset string, field string) ([]chargebee.Customer, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListCustomersPageSortBy(site, offset, sortBy)
}

func RetrieveCustomerById(id string) (*chargebee.Customer, error) {
	return RetrieveSiteCustomerById(client.Site(), id)
}

func RetrieveSiteCustomerById(site string, id string) (*chargebee.Customer, error) {
	return chargebee.RetrieveCustomerById(site, id)
}

// entitlements.go

func ListAllEntitlements() ([]chargebee.Entitlement, error) {
	return ListAllSiteEntitlements(client.Site())
}

func ListAllSiteEntitlements(site string) ([]chargebee.Entitlement, error) {
	return chargebee.ListEntitlements(site)
}

func ListEntitlementsCreatedSince(createdSince *time.Time) ([]chargebee.Entitlement, error) {
	return ListSiteEntitlementsCreatedSince(client.Site(), createdSince)
}

func ListSiteEntitlementsCreatedSince(site string, createdSince *time.Time) ([]chargebee.Entitlement, error) {
	return chargebee.ListEntitlementsCreatedSince(site, createdSince)
}

func ListEntitlementsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Entitlement, error) {
	return ListSiteEntitlementsUpdatedSince(client.Site(), updatedSince)
}

func ListSiteEntitlementsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Entitlement, error) {
	return chargebee.ListEntitlementsUpdatedSince(site, updatedSince)
}

func ListEntitlements() ([]chargebee.Entitlement, string, error) {
	return ListSiteEntitlements(client.Site())
}

func ListSiteEntitlements(site string) ([]chargebee.Entitlement, string, error) {
	return chargebee.ListEntitlementsPage(site, "")
}

func ListEntitlementsPage(offset string) ([]chargebee.Entitlement, string, error) {
	return ListSiteEntitlementsPage(client.Site(), offset)
}

func ListSiteEntitlementsPage(site string, offset string) ([]chargebee.Entitlement, string, error) {
	return chargebee.ListEntitlementsPage(site, offset)
}

func ListEntitlementsPageSortBy(offset string, field string) ([]chargebee.Entitlement, string, error) {
	return ListSiteEntitlementsPageSortBy(client.Site(), offset, field)
}

func ListSiteEntitlementsPageSortBy(site string, offset string, field string) ([]chargebee.Entitlement, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListEntitlementsPageSortBy(site, offset, sortBy)
}

// events.go

func ListAllEvents() ([]chargebee.Event, error) {
	return ListAllSiteEvents(client.Site())
}

func ListAllSiteEvents(site string) ([]chargebee.Event, error) {
	return chargebee.ListEvents(site)
}

func ListEventsCreatedSince(createdSince *time.Time) ([]chargebee.Event, error) {
	return ListSiteEventsCreatedSince(client.Site(), createdSince)
}

func ListSiteEventsCreatedSince(site string, createdSince *time.Time) ([]chargebee.Event, error) {
	return chargebee.ListEventsCreatedSince(site, createdSince)
}

func ListEventsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Event, error) {
	return ListSiteEventsUpdatedSince(client.Site(), updatedSince)
}

func ListSiteEventsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Event, error) {
	return chargebee.ListEventsUpdatedSince(site, updatedSince)
}

func ListEvents() ([]chargebee.Event, string, error) {
	return ListSiteEvents(client.Site())
}

func ListSiteEvents(site string) ([]chargebee.Event, string, error) {
	return chargebee.ListEventsPage(site, "")
}

func ListEventsPage(offset string) ([]chargebee.Event, string, error) {
	return ListSiteEventsPage(client.Site(), offset)
}

func ListSiteEventsPage(site string, offset string) ([]chargebee.Event, string, error) {
	return chargebee.ListEventsPage(site, offset)
}

func ListEventsPageSortBy(offset string, field string) ([]chargebee.Event, string, error) {
	return ListSiteEventsPageSortBy(client.Site(), offset, field)
}

func ListSiteEventsPageSortBy(site string, offset string, field string) ([]chargebee.Event, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListEventsPageSortBy(site, offset, sortBy)
}

func RetrieveEventById(id string) (*chargebee.Event, error) {
	return RetrieveSiteEventById(client.Site(), id)
}

func RetrieveSiteEventById(site string, id string) (*chargebee.Event, error) {
	return chargebee.RetrieveEventById(site, id)
}

// features.go

func ListAllFeatures() ([]chargebee.Feature, error) {
	return ListAllSiteFeatures(client.Site())
}

func ListAllSiteFeatures(site string) ([]chargebee.Feature, error) {
	return chargebee.ListFeatures(site)
}

func ListFeaturesCreatedSince(createdSince *time.Time) ([]chargebee.Feature, error) {
	return ListSiteFeaturesCreatedSince(client.Site(), createdSince)
}

func ListSiteFeaturesCreatedSince(site string, createdSince *time.Time) ([]chargebee.Feature, error) {
	return chargebee.ListFeaturesCreatedSince(site, createdSince)
}

func ListFeaturesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Feature, error) {
	return ListSiteFeaturesUpdatedSince(client.Site(), updatedSince)
}

func ListSiteFeaturesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Feature, error) {
	return chargebee.ListFeaturesUpdatedSince(site, updatedSince)
}

func ListFeatures() ([]chargebee.Feature, string, error) {
	return ListSiteFeatures(client.Site())
}

func ListSiteFeatures(site string) ([]chargebee.Feature, string, error) {
	return chargebee.ListFeaturesPage(site, "")
}

func ListFeaturesPage(offset string) ([]chargebee.Feature, string, error) {
	return ListSiteFeaturesPage(client.Site(), offset)
}

func ListSiteFeaturesPage(site string, offset string) ([]chargebee.Feature, string, error) {
	return chargebee.ListFeaturesPage(site, offset)
}

func ListFeaturesPageSortBy(offset string, field string) ([]chargebee.Feature, string, error) {
	return ListSiteFeaturesPageSortBy(client.Site(), offset, field)
}

func ListSiteFeaturesPageSortBy(site string, offset string, field string) ([]chargebee.Feature, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListFeaturesPageSortBy(site, offset, sortBy)
}

func RetrieveFeatureById(id string) (*chargebee.Feature, error) {
	return RetrieveSiteFeatureById(client.Site(), id)
}

func RetrieveSiteFeatureById(site string, id string) (*chargebee.Feature, error) {
	return chargebee.RetrieveFeatureById(site, id)
}

// invoice_credit_notes.go

func ListAllCreditNotes() ([]chargebee.CreditNote, error) {
	return ListAllSiteCreditNotes(client.Site())
}

func ListAllSiteCreditNotes(site string) ([]chargebee.CreditNote, error) {
	return chargebee.ListCreditNotes(site)
}

func ListCreditNotesCreatedSince(createdSince *time.Time) ([]chargebee.CreditNote, error) {
	return ListSiteCreditNotesCreatedSince(client.Site(), createdSince)
}

func ListSiteCreditNotesCreatedSince(site string, createdSince *time.Time) ([]chargebee.CreditNote, error) {
	return chargebee.ListCreditNotesCreatedSince(site, createdSince)
}

func ListCreditNotesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.CreditNote, error) {
	return ListSiteCreditNotesUpdatedSince(client.Site(), updatedSince)
}

func ListSiteCreditNotesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.CreditNote, error) {
	return chargebee.ListCreditNotesUpdatedSince(site, updatedSince)
}

func ListCreditNotes() ([]chargebee.CreditNote, string, error) {
	return ListSiteCreditNotes(client.Site())
}

func ListSiteCreditNotes(site string) ([]chargebee.CreditNote, string, error) {
	return chargebee.ListCreditNotesPage(site, "")
}

func ListCreditNotesPage(offset string) ([]chargebee.CreditNote, string, error) {
	return ListSiteCreditNotesPage(client.Site(), offset)
}

func ListSiteCreditNotesPage(site string, offset string) ([]chargebee.CreditNote, string, error) {
	return chargebee.ListCreditNotesPage(site, offset)
}

func ListCreditNotesPageSortBy(offset string, field string) ([]chargebee.CreditNote, string, error) {
	return ListSiteCreditNotesPageSortBy(client.Site(), offset, field)
}

func ListSiteCreditNotesPageSortBy(site string, offset string, field string) ([]chargebee.CreditNote, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListCreditNotesPageSortBy(site, offset, sortBy)
}

func RetrieveCreditNoteById(id string) (*chargebee.CreditNote, error) {
	return RetrieveSiteCreditNoteById(client.Site(), id)
}

func RetrieveSiteCreditNoteById(site string, id string) (*chargebee.CreditNote, error) {
	return chargebee.RetrieveCreditNoteById(site, id)
}

// invoice_payment_reference_numbers.go

func ListAllPaymentReferenceNumbers(id string) ([]chargebee.PaymentReferenceNumber, error) {
	return ListAllSitePaymentReferenceNumbers(client.Site(), id)
}

func ListAllSitePaymentReferenceNumbers(site, id string) ([]chargebee.PaymentReferenceNumber, error) {
	return chargebee.ListPaymentReferenceNumbers(site, id)
}

func ListPaymentReferenceNumbersCreatedSince(id string, createdSince *time.Time) ([]chargebee.PaymentReferenceNumber, error) {
	return ListSitePaymentReferenceNumbersCreatedSince(client.Site(), id, createdSince)
}

func ListSitePaymentReferenceNumbersCreatedSince(site, id string, createdSince *time.Time) ([]chargebee.PaymentReferenceNumber, error) {
	return chargebee.ListPaymentReferenceNumbersCreatedSince(site, id, createdSince)
}

func ListPaymentReferenceNumbersUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.PaymentReferenceNumber, error) {
	return ListSitePaymentReferenceNumbersUpdatedSince(client.Site(), id, updatedSince)
}

func ListSitePaymentReferenceNumbersUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.PaymentReferenceNumber, error) {
	return chargebee.ListPaymentReferenceNumbersUpdatedSince(site, id, updatedSince)
}

func ListPaymentReferenceNumbers(id string) ([]chargebee.PaymentReferenceNumber, string, error) {
	return ListSitePaymentReferenceNumbers(client.Site(), id)
}

func ListSitePaymentReferenceNumbers(site, id string) ([]chargebee.PaymentReferenceNumber, string, error) {
	return chargebee.ListPaymentReferenceNumbersPage(site, id, "")
}

func ListPaymentReferenceNumbersPage(id, offset string) ([]chargebee.PaymentReferenceNumber, string, error) {
	return ListSitePaymentReferenceNumbersPage(client.Site(), id, offset)
}

func ListSitePaymentReferenceNumbersPage(site, id string, offset string) ([]chargebee.PaymentReferenceNumber, string, error) {
	return chargebee.ListPaymentReferenceNumbersPage(site, id, offset)
}

func ListPaymentReferenceNumbersPageSortBy(id, offset string, field string) ([]chargebee.PaymentReferenceNumber, string, error) {
	return ListSitePaymentReferenceNumbersPageSortBy(client.Site(), id, offset, field)
}

func ListSitePaymentReferenceNumbersPageSortBy(site, id string, offset string, field string) ([]chargebee.PaymentReferenceNumber, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListPaymentReferenceNumbersPageSortBy(site, id, offset, sortBy)
}

// invoice_payments.go

func ListAllInvoicePayments(id string) ([]chargebee.Transaction, error) {
	return ListAllSiteInvoicePayments(client.Site(), id)
}

func ListAllSiteInvoicePayments(site, id string) ([]chargebee.Transaction, error) {
	return chargebee.ListInvoicePayments(site, id)
}

func ListInvoicePaymentsCreatedSince(id string, createdSince *time.Time) ([]chargebee.Transaction, error) {
	return ListSiteInvoicePaymentsCreatedSince(client.Site(), id, createdSince)
}

func ListSiteInvoicePaymentsCreatedSince(site, id string, createdSince *time.Time) ([]chargebee.Transaction, error) {
	return chargebee.ListInvoicePaymentsCreatedSince(site, id, createdSince)
}

func ListInvoicePaymentsUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.Transaction, error) {
	return ListSiteInvoicePaymentsUpdatedSince(client.Site(), id, updatedSince)
}

func ListSiteInvoicePaymentsUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.Transaction, error) {
	return chargebee.ListInvoicePaymentsUpdatedSince(site, id, updatedSince)
}

func ListInvoicePayments(id string) ([]chargebee.Transaction, string, error) {
	return ListSiteInvoicePayments(client.Site(), id)
}

func ListSiteInvoicePayments(site, id string) ([]chargebee.Transaction, string, error) {
	return chargebee.ListInvoicePaymentsPage(site, id, "")
}

func ListInvoicePaymentsPage(id, offset string) ([]chargebee.Transaction, string, error) {
	return ListSiteInvoicePaymentsPage(client.Site(), id, offset)
}

func ListSiteInvoicePaymentsPage(site, id string, offset string) ([]chargebee.Transaction, string, error) {
	return chargebee.ListInvoicePaymentsPage(site, id, offset)
}

func ListInvoicePaymentsPageSortBy(id, offset string, field string) ([]chargebee.Transaction, string, error) {
	return ListSiteInvoicePaymentsPageSortBy(client.Site(), id, offset, field)
}

func ListSiteInvoicePaymentsPageSortBy(site, id string, offset string, field string) ([]chargebee.Transaction, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListInvoicePaymentsPageSortBy(site, id, offset, sortBy)
}

// invoice_promotional_credits.go

func ListAllPromotionalCredits() ([]chargebee.PromotionalCredit, error) {
	return ListAllSitePromotionalCredits(client.Site())
}

func ListAllSitePromotionalCredits(site string) ([]chargebee.PromotionalCredit, error) {
	return chargebee.ListPromotionalCredits(site)
}

func ListPromotionalCreditsCreatedSince(createdSince *time.Time) ([]chargebee.PromotionalCredit, error) {
	return ListSitePromotionalCreditsCreatedSince(client.Site(), createdSince)
}

func ListSitePromotionalCreditsCreatedSince(site string, createdSince *time.Time) ([]chargebee.PromotionalCredit, error) {
	return chargebee.ListPromotionalCreditsCreatedSince(site, createdSince)
}

func ListPromotionalCreditsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.PromotionalCredit, error) {
	return ListSitePromotionalCreditsUpdatedSince(client.Site(), updatedSince)
}

func ListSitePromotionalCreditsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.PromotionalCredit, error) {
	return chargebee.ListPromotionalCreditsUpdatedSince(site, updatedSince)
}

func ListPromotionalCredits() ([]chargebee.PromotionalCredit, string, error) {
	return ListSitePromotionalCredits(client.Site())
}

func ListSitePromotionalCredits(site string) ([]chargebee.PromotionalCredit, string, error) {
	return chargebee.ListPromotionalCreditsPage(site, "")
}

func ListPromotionalCreditsPage(offset string) ([]chargebee.PromotionalCredit, string, error) {
	return ListSitePromotionalCreditsPage(client.Site(), offset)
}

func ListSitePromotionalCreditsPage(site string, offset string) ([]chargebee.PromotionalCredit, string, error) {
	return chargebee.ListPromotionalCreditsPage(site, offset)
}

func ListPromotionalCreditsPageSortBy(offset string, field string) ([]chargebee.PromotionalCredit, string, error) {
	return ListSitePromotionalCreditsPageSortBy(client.Site(), offset, field)
}

func ListSitePromotionalCreditsPageSortBy(site string, offset string, field string) ([]chargebee.PromotionalCredit, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListPromotionalCreditsPageSortBy(site, offset, sortBy)
}

func RetrievePromotionalCreditById(id string) (*chargebee.PromotionalCredit, error) {
	return RetrieveSitePromotionalCreditById(client.Site(), id)
}

func RetrieveSitePromotionalCreditById(site string, id string) (*chargebee.PromotionalCredit, error) {
	return chargebee.RetrievePromotionalCreditById(site, id)
}

// invoice_unbilled_charges.go

func ListAllInvoiceUnbilledCharges() ([]chargebee.InvoiceUnbilledCharge, error) {
	return ListAllSiteInvoiceUnbilledCharges(client.Site())
}

func ListAllSiteInvoiceUnbilledCharges(site string) ([]chargebee.InvoiceUnbilledCharge, error) {
	return chargebee.ListInvoiceUnbilledCharges(site)
}

func ListInvoiceUnbilledChargesCreatedSince(createdSince *time.Time) ([]chargebee.InvoiceUnbilledCharge, error) {
	return ListSiteInvoiceUnbilledChargesCreatedSince(client.Site(), createdSince)
}

func ListSiteInvoiceUnbilledChargesCreatedSince(site string, createdSince *time.Time) ([]chargebee.InvoiceUnbilledCharge, error) {
	return chargebee.ListInvoiceUnbilledChargesCreatedSince(site, createdSince)
}

func ListInvoiceUnbilledChargesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.InvoiceUnbilledCharge, error) {
	return ListSiteInvoiceUnbilledChargesUpdatedSince(client.Site(), updatedSince)
}

func ListSiteInvoiceUnbilledChargesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.InvoiceUnbilledCharge, error) {
	return chargebee.ListInvoiceUnbilledChargesUpdatedSince(site, updatedSince)
}

func ListInvoiceUnbilledCharges() ([]chargebee.InvoiceUnbilledCharge, string, error) {
	return ListSiteInvoiceUnbilledCharges(client.Site())
}

func ListSiteInvoiceUnbilledCharges(site string) ([]chargebee.InvoiceUnbilledCharge, string, error) {
	return chargebee.ListInvoiceUnbilledChargesPage(site, "")
}

func ListInvoiceUnbilledChargesPage(offset string) ([]chargebee.InvoiceUnbilledCharge, string, error) {
	return ListSiteInvoiceUnbilledChargesPage(client.Site(), offset)
}

func ListSiteInvoiceUnbilledChargesPage(site string, offset string) ([]chargebee.InvoiceUnbilledCharge, string, error) {
	return chargebee.ListInvoiceUnbilledChargesPage(site, offset)
}

func ListInvoiceUnbilledChargesPageSortBy(offset string, field string) ([]chargebee.InvoiceUnbilledCharge, string, error) {
	return ListSiteInvoiceUnbilledChargesPageSortBy(client.Site(), offset, field)
}

func ListSiteInvoiceUnbilledChargesPageSortBy(site string, offset string, field string) ([]chargebee.InvoiceUnbilledCharge, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListInvoiceUnbilledChargesPageSortBy(site, offset, sortBy)
}

// invoice_vouchers.go

func ListAllInvoicePaymentVouchers(id string) ([]chargebee.PaymentVoucher, error) {
	return ListAllSiteInvoicePaymentVouchers(client.Site(), id)
}

func ListAllSiteInvoicePaymentVouchers(site, id string) ([]chargebee.PaymentVoucher, error) {
	return chargebee.ListInvoicePaymentVouchers(site, id)
}

func ListInvoicePaymentVouchersCreatedSince(id string, createdSince *time.Time) ([]chargebee.PaymentVoucher, error) {
	return ListSiteInvoicePaymentVouchersCreatedSince(client.Site(), id, createdSince)
}

func ListSiteInvoicePaymentVouchersCreatedSince(site, id string, createdSince *time.Time) ([]chargebee.PaymentVoucher, error) {
	return chargebee.ListInvoicePaymentVouchersCreatedSince(site, id, createdSince)
}

func ListInvoicePaymentVouchersUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.PaymentVoucher, error) {
	return ListSiteInvoicePaymentVouchersUpdatedSince(client.Site(), id, updatedSince)
}

func ListSiteInvoicePaymentVouchersUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.PaymentVoucher, error) {
	return chargebee.ListInvoicePaymentVouchersUpdatedSince(site, id, updatedSince)
}

func ListInvoicePaymentVouchers(id string) ([]chargebee.PaymentVoucher, string, error) {
	return ListSiteInvoicePaymentVouchers(client.Site(), id)
}

func ListSiteInvoicePaymentVouchers(site, id string) ([]chargebee.PaymentVoucher, string, error) {
	return chargebee.ListInvoicePaymentVouchersPage(site, id, "")
}

func ListInvoicePaymentVouchersPage(id, offset string) ([]chargebee.PaymentVoucher, string, error) {
	return ListSiteInvoicePaymentVouchersPage(client.Site(), id, offset)
}

func ListSiteInvoicePaymentVouchersPage(site, id string, offset string) ([]chargebee.PaymentVoucher, string, error) {
	return chargebee.ListInvoicePaymentVouchersPage(site, id, offset)
}

func ListInvoicePaymentVouchersPageSortBy(id, offset string, field string) ([]chargebee.PaymentVoucher, string, error) {
	return ListSiteInvoicePaymentVouchersPageSortBy(client.Site(), id, offset, field)
}

func ListSiteInvoicePaymentVouchersPageSortBy(site, id string, offset string, field string) ([]chargebee.PaymentVoucher, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListInvoicePaymentVouchersPageSortBy(site, id, offset, sortBy)
}

// invoices.go

func ListAllInvoices() ([]chargebee.Invoice, error) {
	return ListAllSiteInvoices(client.Site())
}

func ListAllSiteInvoices(site string) ([]chargebee.Invoice, error) {
	return chargebee.ListInvoices(site)
}

func ListInvoicesCreatedSince(createdSince *time.Time) ([]chargebee.Invoice, error) {
	return ListSiteInvoicesCreatedSince(client.Site(), createdSince)
}

func ListSiteInvoicesCreatedSince(site string, createdSince *time.Time) ([]chargebee.Invoice, error) {
	return chargebee.ListInvoicesCreatedSince(site, createdSince)
}

func ListInvoicesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Invoice, error) {
	return ListSiteInvoicesUpdatedSince(client.Site(), updatedSince)
}

func ListSiteInvoicesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Invoice, error) {
	return chargebee.ListInvoicesUpdatedSince(site, updatedSince)
}

func ListInvoices() ([]chargebee.Invoice, string, error) {
	return ListSiteInvoices(client.Site())
}

func ListSiteInvoices(site string) ([]chargebee.Invoice, string, error) {
	return chargebee.ListInvoicesPage(site, "")
}

func ListInvoicesPage(offset string) ([]chargebee.Invoice, string, error) {
	return ListSiteInvoicesPage(client.Site(), offset)
}

func ListSiteInvoicesPage(site string, offset string) ([]chargebee.Invoice, string, error) {
	return chargebee.ListInvoicesPage(site, offset)
}

func ListInvoicesPageSortBy(offset string, field string) ([]chargebee.Invoice, string, error) {
	return ListSiteInvoicesPageSortBy(client.Site(), offset, field)
}

func ListSiteInvoicesPageSortBy(site string, offset string, field string) ([]chargebee.Invoice, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListInvoicesPageSortBy(site, offset, sortBy)
}

func RetrieveInvoiceById(id string) (*chargebee.Invoice, error) {
	return RetrieveSiteInvoiceById(client.Site(), id)
}

func RetrieveSiteInvoiceById(site string, id string) (*chargebee.Invoice, error) {
	return chargebee.RetrieveInvoiceById(site, id)
}

// item_attached.go

func ListAllItemsAttached(id string) ([]chargebee.ItemAttached, error) {
	return ListAllSiteItemsAttached(client.Site(), id)
}

func ListAllSiteItemsAttached(site, id string) ([]chargebee.ItemAttached, error) {
	return chargebee.ListItemsAttached(site, id)
}

func ListItemsAttachedCreatedSince(id string, createdSince *time.Time) ([]chargebee.ItemAttached, error) {
	return ListSiteItemsAttachedCreatedSince(client.Site(), id, createdSince)
}

func ListSiteItemsAttachedCreatedSince(site, id string, createdSince *time.Time) ([]chargebee.ItemAttached, error) {
	return chargebee.ListItemsAttachedCreatedSince(site, id, createdSince)
}

func ListItemsAttachedUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.ItemAttached, error) {
	return ListSiteItemsAttachedUpdatedSince(client.Site(), id, updatedSince)
}

func ListSiteItemsAttachedUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.ItemAttached, error) {
	return chargebee.ListItemsAttachedUpdatedSince(site, id, updatedSince)
}

func ListItemsAttached(id string) ([]chargebee.ItemAttached, string, error) {
	return ListSiteItemsAttached(client.Site(), id)
}

func ListSiteItemsAttached(site, id string) ([]chargebee.ItemAttached, string, error) {
	return chargebee.ListItemsAttachedPage(site, id, "")
}

func ListItemsAttachedPage(id, offset string) ([]chargebee.ItemAttached, string, error) {
	return ListSiteItemsAttachedPage(client.Site(), id, offset)
}

func ListSiteItemsAttachedPage(site, id string, offset string) ([]chargebee.ItemAttached, string, error) {
	return chargebee.ListItemsAttachedPage(site, id, offset)
}

func ListItemsAttachedPageSortBy(id, offset string, field string) ([]chargebee.ItemAttached, string, error) {
	return ListSiteItemsAttachedPageSortBy(client.Site(), id, offset, field)
}

func ListSiteItemsAttachedPageSortBy(site, id string, offset string, field string) ([]chargebee.ItemAttached, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListItemsAttachedPageSortBy(site, id, offset, sortBy)
}

func RetrieveItemAttachedById(id string) (*chargebee.ItemAttached, error) {
	return RetrieveSiteItemAttachedById(client.Site(), id)
}

func RetrieveSiteItemAttachedById(site string, id string) (*chargebee.ItemAttached, error) {
	return chargebee.RetrieveItemAttachedById(site, id)
}

// item_coupon_codes.go

func ListAllCouponCodes() ([]chargebee.CouponCode, error) {
	return ListAllSiteCouponCodes(client.Site())
}

func ListAllSiteCouponCodes(site string) ([]chargebee.CouponCode, error) {
	return chargebee.ListCouponCodes(site)
}

func ListCouponCodesCreatedSince(createdSince *time.Time) ([]chargebee.CouponCode, error) {
	return ListSiteCouponCodesCreatedSince(client.Site(), createdSince)
}

func ListSiteCouponCodesCreatedSince(site string, createdSince *time.Time) ([]chargebee.CouponCode, error) {
	return chargebee.ListCouponCodesCreatedSince(site, createdSince)
}

func ListCouponCodesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.CouponCode, error) {
	return ListSiteCouponCodesUpdatedSince(client.Site(), updatedSince)
}

func ListSiteCouponCodesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.CouponCode, error) {
	return chargebee.ListCouponCodesUpdatedSince(site, updatedSince)
}

func ListCouponCodes() ([]chargebee.CouponCode, string, error) {
	return ListSiteCouponCodes(client.Site())
}

func ListSiteCouponCodes(site string) ([]chargebee.CouponCode, string, error) {
	return chargebee.ListCouponCodesPage(site, "")
}

func ListCouponCodesPage(offset string) ([]chargebee.CouponCode, string, error) {
	return ListSiteCouponCodesPage(client.Site(), offset)
}

func ListSiteCouponCodesPage(site string, offset string) ([]chargebee.CouponCode, string, error) {
	return chargebee.ListCouponCodesPage(site, offset)
}

func ListCouponCodesPageSortBy(offset string, field string) ([]chargebee.CouponCode, string, error) {
	return ListSiteCouponCodesPageSortBy(client.Site(), offset, field)
}

func ListSiteCouponCodesPageSortBy(site string, offset string, field string) ([]chargebee.CouponCode, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListCouponCodesPageSortBy(site, offset, sortBy)
}

// item_coupon_sets.go

func ListAllCouponSets() ([]chargebee.CouponSet, error) {
	return ListAllSiteCouponSets(client.Site())
}

func ListAllSiteCouponSets(site string) ([]chargebee.CouponSet, error) {
	return chargebee.ListCouponSets(site)
}

func ListCouponSetsCreatedSince(createdSince *time.Time) ([]chargebee.CouponSet, error) {
	return ListSiteCouponSetsCreatedSince(client.Site(), createdSince)
}

func ListSiteCouponSetsCreatedSince(site string, createdSince *time.Time) ([]chargebee.CouponSet, error) {
	return chargebee.ListCouponSetsCreatedSince(site, createdSince)
}

func ListCouponSetsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.CouponSet, error) {
	return ListSiteCouponSetsUpdatedSince(client.Site(), updatedSince)
}

func ListSiteCouponSetsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.CouponSet, error) {
	return chargebee.ListCouponSetsUpdatedSince(site, updatedSince)
}

func ListCouponSets() ([]chargebee.CouponSet, string, error) {
	return ListSiteCouponSets(client.Site())
}

func ListSiteCouponSets(site string) ([]chargebee.CouponSet, string, error) {
	return chargebee.ListCouponSetsPage(site, "")
}

func ListCouponSetsPage(offset string) ([]chargebee.CouponSet, string, error) {
	return ListSiteCouponSetsPage(client.Site(), offset)
}

func ListSiteCouponSetsPage(site string, offset string) ([]chargebee.CouponSet, string, error) {
	return chargebee.ListCouponSetsPage(site, offset)
}

func ListCouponSetsPageSortBy(offset string, field string) ([]chargebee.CouponSet, string, error) {
	return ListSiteCouponSetsPageSortBy(client.Site(), offset, field)
}

func ListSiteCouponSetsPageSortBy(site string, offset string, field string) ([]chargebee.CouponSet, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListCouponSetsPageSortBy(site, offset, sortBy)
}

func RetrieveCouponSetById(id string) (*chargebee.CouponSet, error) {
	return RetrieveSiteCouponSetById(client.Site(), id)
}

func RetrieveSiteCouponSetById(site string, id string) (*chargebee.CouponSet, error) {
	return chargebee.RetrieveCouponSetById(site, id)
}

// item_coupons.go

func ListAllCoupons() ([]chargebee.Coupon, error) {
	return ListAllSiteCoupons(client.Site())
}

func ListAllSiteCoupons(site string) ([]chargebee.Coupon, error) {
	return chargebee.ListCoupons(site)
}

func ListCouponsCreatedSince(createdSince *time.Time) ([]chargebee.Coupon, error) {
	return ListSiteCouponsCreatedSince(client.Site(), createdSince)
}

func ListSiteCouponsCreatedSince(site string, createdSince *time.Time) ([]chargebee.Coupon, error) {
	return chargebee.ListCouponsCreatedSince(site, createdSince)
}

func ListCouponsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Coupon, error) {
	return ListSiteCouponsUpdatedSince(client.Site(), updatedSince)
}

func ListSiteCouponsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Coupon, error) {
	return chargebee.ListCouponsUpdatedSince(site, updatedSince)
}

func ListCoupons() ([]chargebee.Coupon, string, error) {
	return ListSiteCoupons(client.Site())
}

func ListSiteCoupons(site string) ([]chargebee.Coupon, string, error) {
	return chargebee.ListCouponsPage(site, "")
}

func ListCouponsPage(offset string) ([]chargebee.Coupon, string, error) {
	return ListSiteCouponsPage(client.Site(), offset)
}

func ListSiteCouponsPage(site string, offset string) ([]chargebee.Coupon, string, error) {
	return chargebee.ListCouponsPage(site, offset)
}

func ListCouponsPageSortBy(offset string, field string) ([]chargebee.Coupon, string, error) {
	return ListSiteCouponsPageSortBy(client.Site(), offset, field)
}

func ListSiteCouponsPageSortBy(site string, offset string, field string) ([]chargebee.Coupon, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListCouponsPageSortBy(site, offset, sortBy)
}

func RetrieveCouponById(id string) (*chargebee.Coupon, error) {
	return RetrieveSiteCouponById(client.Site(), id)
}

func RetrieveSiteCouponById(site string, id string) (*chargebee.Coupon, error) {
	return chargebee.RetrieveCouponById(site, id)
}

// item_differential_prices.go

func ListAllDifferentialPrices() ([]chargebee.DifferentialPrice, error) {
	return ListAllSiteDifferentialPrices(client.Site())
}

func ListAllSiteDifferentialPrices(site string) ([]chargebee.DifferentialPrice, error) {
	return chargebee.ListDifferentialPrices(site)
}

func ListDifferentialPricesCreatedSince(createdSince *time.Time) ([]chargebee.DifferentialPrice, error) {
	return ListSiteDifferentialPricesCreatedSince(client.Site(), createdSince)
}

func ListSiteDifferentialPricesCreatedSince(site string, createdSince *time.Time) ([]chargebee.DifferentialPrice, error) {
	return chargebee.ListDifferentialPricesCreatedSince(site, createdSince)
}

func ListDifferentialPricesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.DifferentialPrice, error) {
	return ListSiteDifferentialPricesUpdatedSince(client.Site(), updatedSince)
}

func ListSiteDifferentialPricesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.DifferentialPrice, error) {
	return chargebee.ListDifferentialPricesUpdatedSince(site, updatedSince)
}

func ListDifferentialPrices() ([]chargebee.DifferentialPrice, string, error) {
	return ListSiteDifferentialPrices(client.Site())
}

func ListSiteDifferentialPrices(site string) ([]chargebee.DifferentialPrice, string, error) {
	return chargebee.ListDifferentialPricesPage(site, "")
}

func ListDifferentialPricesPage(offset string) ([]chargebee.DifferentialPrice, string, error) {
	return ListSiteDifferentialPricesPage(client.Site(), offset)
}

func ListSiteDifferentialPricesPage(site string, offset string) ([]chargebee.DifferentialPrice, string, error) {
	return chargebee.ListDifferentialPricesPage(site, offset)
}

func ListDifferentialPricesPageSortBy(offset string, field string) ([]chargebee.DifferentialPrice, string, error) {
	return ListSiteDifferentialPricesPageSortBy(client.Site(), offset, field)
}

func ListSiteDifferentialPricesPageSortBy(site string, offset string, field string) ([]chargebee.DifferentialPrice, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListDifferentialPricesPageSortBy(site, offset, sortBy)
}

func RetrieveDifferentialPriceById(id string) (*chargebee.DifferentialPrice, error) {
	return RetrieveSiteDifferentialPriceById(client.Site(), id)
}

func RetrieveSiteDifferentialPriceById(site string, id string) (*chargebee.DifferentialPrice, error) {
	return chargebee.RetrieveDifferentialPriceById(site, id)
}

// item_families.go

func ListAllItemFamilies() ([]chargebee.ItemFamily, error) {
	return ListAllSiteItemFamilies(client.Site())
}

func ListAllSiteItemFamilies(site string) ([]chargebee.ItemFamily, error) {
	return chargebee.ListItemFamilies(site)
}

func ListItemFamiliesCreatedSince(createdSince *time.Time) ([]chargebee.ItemFamily, error) {
	return ListSiteItemFamiliesCreatedSince(client.Site(), createdSince)
}

func ListSiteItemFamiliesCreatedSince(site string, createdSince *time.Time) ([]chargebee.ItemFamily, error) {
	return chargebee.ListItemFamiliesCreatedSince(site, createdSince)
}

func ListItemFamiliesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.ItemFamily, error) {
	return ListSiteItemFamiliesUpdatedSince(client.Site(), updatedSince)
}

func ListSiteItemFamiliesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.ItemFamily, error) {
	return chargebee.ListItemFamiliesUpdatedSince(site, updatedSince)
}

func ListItemFamilies() ([]chargebee.ItemFamily, string, error) {
	return ListSiteItemFamilies(client.Site())
}

func ListSiteItemFamilies(site string) ([]chargebee.ItemFamily, string, error) {
	return chargebee.ListItemFamiliesPage(site, "")
}

func ListItemFamiliesPage(offset string) ([]chargebee.ItemFamily, string, error) {
	return ListSiteItemFamiliesPage(client.Site(), offset)
}

func ListSiteItemFamiliesPage(site string, offset string) ([]chargebee.ItemFamily, string, error) {
	return chargebee.ListItemFamiliesPage(site, offset)
}

func ListItemFamiliesPageSortBy(offset string, field string) ([]chargebee.ItemFamily, string, error) {
	return ListSiteItemFamiliesPageSortBy(client.Site(), offset, field)
}

func ListSiteItemFamiliesPageSortBy(site string, offset string, field string) ([]chargebee.ItemFamily, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListItemFamiliesPageSortBy(site, offset, sortBy)
}

func RetrieveItemFamilyById(id string) (*chargebee.ItemFamily, error) {
	return RetrieveSiteItemFamilyById(client.Site(), id)
}

func RetrieveSiteItemFamilyById(site string, id string) (*chargebee.ItemFamily, error) {
	return chargebee.RetrieveItemFamilyById(site, id)
}

// item_prices.go

func ListAllItemPrices() ([]chargebee.ItemPrice, error) {
	return ListAllSiteItemPrices(client.Site())
}

func ListAllSiteItemPrices(site string) ([]chargebee.ItemPrice, error) {
	return chargebee.ListItemPrices(site)
}

func ListItemPricesCreatedSince(createdSince *time.Time) ([]chargebee.ItemPrice, error) {
	return ListSiteItemPricesCreatedSince(client.Site(), createdSince)
}

func ListSiteItemPricesCreatedSince(site string, createdSince *time.Time) ([]chargebee.ItemPrice, error) {
	return chargebee.ListItemPricesCreatedSince(site, createdSince)
}

func ListItemPricesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.ItemPrice, error) {
	return ListSiteItemPricesUpdatedSince(client.Site(), updatedSince)
}

func ListSiteItemPricesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.ItemPrice, error) {
	return chargebee.ListItemPricesUpdatedSince(site, updatedSince)
}

func ListItemPrices() ([]chargebee.ItemPrice, string, error) {
	return ListSiteItemPrices(client.Site())
}

func ListSiteItemPrices(site string) ([]chargebee.ItemPrice, string, error) {
	return chargebee.ListItemPricesPage(site, "")
}

func ListItemPricesPage(offset string) ([]chargebee.ItemPrice, string, error) {
	return ListSiteItemPricesPage(client.Site(), offset)
}

func ListSiteItemPricesPage(site string, offset string) ([]chargebee.ItemPrice, string, error) {
	return chargebee.ListItemPricesPage(site, offset)
}

func ListItemPricesPageSortBy(offset string, field string) ([]chargebee.ItemPrice, string, error) {
	return ListSiteItemPricesPageSortBy(client.Site(), offset, field)
}

func ListSiteItemPricesPageSortBy(site string, offset string, field string) ([]chargebee.ItemPrice, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListItemPricesPageSortBy(site, offset, sortBy)
}

func RetrieveItemPriceById(id string) (*chargebee.ItemPrice, error) {
	return RetrieveSiteItemPriceById(client.Site(), id)
}

func RetrieveSiteItemPriceById(site string, id string) (*chargebee.ItemPrice, error) {
	return chargebee.RetrieveItemPriceById(site, id)
}

// items.go

func ListAllItems() ([]chargebee.Item, error) {
	return ListAllSiteItems(client.Site())
}

func ListAllSiteItems(site string) ([]chargebee.Item, error) {
	return chargebee.ListItems(site)
}

func ListItemsCreatedSince(createdSince *time.Time) ([]chargebee.Item, error) {
	return ListSiteItemsCreatedSince(client.Site(), createdSince)
}

func ListSiteItemsCreatedSince(site string, createdSince *time.Time) ([]chargebee.Item, error) {
	return chargebee.ListItemsCreatedSince(site, createdSince)
}

func ListItemsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Item, error) {
	return ListSiteItemsUpdatedSince(client.Site(), updatedSince)
}

func ListSiteItemsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Item, error) {
	return chargebee.ListItemsUpdatedSince(site, updatedSince)
}

func ListItems() ([]chargebee.Item, string, error) {
	return ListSiteItems(client.Site())
}

func ListSiteItems(site string) ([]chargebee.Item, string, error) {
	return chargebee.ListItemsPage(site, "")
}

func ListItemsPage(offset string) ([]chargebee.Item, string, error) {
	return ListSiteItemsPage(client.Site(), offset)
}

func ListSiteItemsPage(site string, offset string) ([]chargebee.Item, string, error) {
	return chargebee.ListItemsPage(site, offset)
}

func ListItemsPageSortBy(offset string, field string) ([]chargebee.Item, string, error) {
	return ListSiteItemsPageSortBy(client.Site(), offset, field)
}

func ListSiteItemsPageSortBy(site string, offset string, field string) ([]chargebee.Item, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListItemsPageSortBy(site, offset, sortBy)
}

func RetrieveItemById(id string) (*chargebee.Item, error) {
	return RetrieveSiteItemById(client.Site(), id)
}

func RetrieveSiteItemById(site string, id string) (*chargebee.Item, error) {
	return chargebee.RetrieveItemById(site, id)
}

// orders.go

func ListAllOrders() ([]chargebee.Order, error) {
	return ListAllSiteOrders(client.Site())
}

func ListAllSiteOrders(site string) ([]chargebee.Order, error) {
	return chargebee.ListOrders(site)
}

func ListOrdersCreatedSince(createdSince *time.Time) ([]chargebee.Order, error) {
	return ListSiteOrdersCreatedSince(client.Site(), createdSince)
}

func ListSiteOrdersCreatedSince(site string, createdSince *time.Time) ([]chargebee.Order, error) {
	return chargebee.ListOrdersCreatedSince(site, createdSince)
}

func ListOrdersUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Order, error) {
	return ListSiteOrdersUpdatedSince(client.Site(), updatedSince)
}

func ListSiteOrdersUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Order, error) {
	return chargebee.ListOrdersUpdatedSince(site, updatedSince)
}

func ListOrders() ([]chargebee.Order, string, error) {
	return ListSiteOrders(client.Site())
}

func ListSiteOrders(site string) ([]chargebee.Order, string, error) {
	return chargebee.ListOrdersPage(site, "")
}

func ListOrdersPage(offset string) ([]chargebee.Order, string, error) {
	return ListSiteOrdersPage(client.Site(), offset)
}

func ListSiteOrdersPage(site string, offset string) ([]chargebee.Order, string, error) {
	return chargebee.ListOrdersPage(site, offset)
}

func ListOrdersPageSortBy(offset string, field string) ([]chargebee.Order, string, error) {
	return ListSiteOrdersPageSortBy(client.Site(), offset, field)
}

func ListSiteOrdersPageSortBy(site string, offset string, field string) ([]chargebee.Order, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListOrdersPageSortBy(site, offset, sortBy)
}

func RetrieveOrderById(id string) (*chargebee.Order, error) {
	return RetrieveSiteOrderById(client.Site(), id)
}

func RetrieveSiteOrderById(site string, id string) (*chargebee.Order, error) {
	return chargebee.RetrieveOrderById(site, id)
}

// payment_sources.go

func ListAllPaymentSources() ([]chargebee.PaymentSource, error) {
	return ListAllSitePaymentSources(client.Site())
}

func ListAllSitePaymentSources(site string) ([]chargebee.PaymentSource, error) {
	return chargebee.ListPaymentSources(site)
}

func ListPaymentSourcesCreatedSince(createdSince *time.Time) ([]chargebee.PaymentSource, error) {
	return ListSitePaymentSourcesCreatedSince(client.Site(), createdSince)
}

func ListSitePaymentSourcesCreatedSince(site string, createdSince *time.Time) ([]chargebee.PaymentSource, error) {
	return chargebee.ListPaymentSourcesCreatedSince(site, createdSince)
}

func ListPaymentSourcesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.PaymentSource, error) {
	return ListSitePaymentSourcesUpdatedSince(client.Site(), updatedSince)
}

func ListSitePaymentSourcesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.PaymentSource, error) {
	return chargebee.ListPaymentSourcesUpdatedSince(site, updatedSince)
}

func ListPaymentSources() ([]chargebee.PaymentSource, string, error) {
	return ListSitePaymentSources(client.Site())
}

func ListSitePaymentSources(site string) ([]chargebee.PaymentSource, string, error) {
	return chargebee.ListPaymentSourcesPage(site, "")
}

func ListPaymentSourcesPage(offset string) ([]chargebee.PaymentSource, string, error) {
	return ListSitePaymentSourcesPage(client.Site(), offset)
}

func ListSitePaymentSourcesPage(site string, offset string) ([]chargebee.PaymentSource, string, error) {
	return chargebee.ListPaymentSourcesPage(site, offset)
}

func ListPaymentSourcesPageSortBy(offset string, field string) ([]chargebee.PaymentSource, string, error) {
	return ListSitePaymentSourcesPageSortBy(client.Site(), offset, field)
}

func ListSitePaymentSourcesPageSortBy(site string, offset string, field string) ([]chargebee.PaymentSource, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListPaymentSourcesPageSortBy(site, offset, sortBy)
}

func RetrievePaymentSourceById(id string) (*chargebee.PaymentSource, error) {
	return RetrieveSitePaymentSourceById(client.Site(), id)
}

func RetrieveSitePaymentSourceById(site string, id string) (*chargebee.PaymentSource, error) {
	return chargebee.RetrievePaymentSourceById(site, id)
}

// payment_virtual_bank_accounts.go

func ListAllVirtualBankAccounts() ([]chargebee.VirtualBankAccount, error) {
	return ListAllSiteVirtualBankAccounts(client.Site())
}

func ListAllSiteVirtualBankAccounts(site string) ([]chargebee.VirtualBankAccount, error) {
	return chargebee.ListVirtualBankAccounts(site)
}

func ListVirtualBankAccountsCreatedSince(createdSince *time.Time) ([]chargebee.VirtualBankAccount, error) {
	return ListSiteVirtualBankAccountsCreatedSince(client.Site(), createdSince)
}

func ListSiteVirtualBankAccountsCreatedSince(site string, createdSince *time.Time) ([]chargebee.VirtualBankAccount, error) {
	return chargebee.ListVirtualBankAccountsCreatedSince(site, createdSince)
}

func ListVirtualBankAccountsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.VirtualBankAccount, error) {
	return ListSiteVirtualBankAccountsUpdatedSince(client.Site(), updatedSince)
}

func ListSiteVirtualBankAccountsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.VirtualBankAccount, error) {
	return chargebee.ListVirtualBankAccountsUpdatedSince(site, updatedSince)
}

func ListVirtualBankAccounts() ([]chargebee.VirtualBankAccount, string, error) {
	return ListSiteVirtualBankAccounts(client.Site())
}

func ListSiteVirtualBankAccounts(site string) ([]chargebee.VirtualBankAccount, string, error) {
	return chargebee.ListVirtualBankAccountsPage(site, "")
}

func ListVirtualBankAccountsPage(offset string) ([]chargebee.VirtualBankAccount, string, error) {
	return ListSiteVirtualBankAccountsPage(client.Site(), offset)
}

func ListSiteVirtualBankAccountsPage(site string, offset string) ([]chargebee.VirtualBankAccount, string, error) {
	return chargebee.ListVirtualBankAccountsPage(site, offset)
}

func ListVirtualBankAccountsPageSortBy(offset string, field string) ([]chargebee.VirtualBankAccount, string, error) {
	return ListSiteVirtualBankAccountsPageSortBy(client.Site(), offset, field)
}

func ListSiteVirtualBankAccountsPageSortBy(site string, offset string, field string) ([]chargebee.VirtualBankAccount, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListVirtualBankAccountsPageSortBy(site, offset, sortBy)
}

func RetrieveVirtualBankAccountById(id string) (*chargebee.VirtualBankAccount, error) {
	return RetrieveSiteVirtualBankAccountById(client.Site(), id)
}

func RetrieveSiteVirtualBankAccountById(site string, id string) (*chargebee.VirtualBankAccount, error) {
	return chargebee.RetrieveVirtualBankAccountById(site, id)
}

// payments.go

func ListAllTransactions() ([]chargebee.Transaction, error) {
	return ListAllSiteTransactions(client.Site())
}

func ListAllSiteTransactions(site string) ([]chargebee.Transaction, error) {
	return chargebee.ListTransactions(site)
}

func ListTransactionsCreatedSince(createdSince *time.Time) ([]chargebee.Transaction, error) {
	return ListSiteTransactionsCreatedSince(client.Site(), createdSince)
}

func ListSiteTransactionsCreatedSince(site string, createdSince *time.Time) ([]chargebee.Transaction, error) {
	return chargebee.ListTransactionsCreatedSince(site, createdSince)
}

func ListTransactionsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Transaction, error) {
	return ListSiteTransactionsUpdatedSince(client.Site(), updatedSince)
}

func ListSiteTransactionsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Transaction, error) {
	return chargebee.ListTransactionsUpdatedSince(site, updatedSince)
}

func ListTransactions() ([]chargebee.Transaction, string, error) {
	return ListSiteTransactions(client.Site())
}

func ListSiteTransactions(site string) ([]chargebee.Transaction, string, error) {
	return chargebee.ListTransactionsPage(site, "")
}

func ListTransactionsPage(offset string) ([]chargebee.Transaction, string, error) {
	return ListSiteTransactionsPage(client.Site(), offset)
}

func ListSiteTransactionsPage(site string, offset string) ([]chargebee.Transaction, string, error) {
	return chargebee.ListTransactionsPage(site, offset)
}

func ListTransactionsPageSortBy(offset string, field string) ([]chargebee.Transaction, string, error) {
	return ListSiteTransactionsPageSortBy(client.Site(), offset, field)
}

func ListSiteTransactionsPageSortBy(site string, offset string, field string) ([]chargebee.Transaction, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListTransactionsPageSortBy(site, offset, sortBy)
}

func RetrieveTransactionById(id string) (*chargebee.Transaction, error) {
	return RetrieveSiteTransactionById(client.Site(), id)
}

func RetrieveSiteTransactionById(site string, id string) (*chargebee.Transaction, error) {
	return chargebee.RetrieveTransactionById(site, id)
}

// quote_line_groups.go

func ListAllQuoteLineGroups(id string) ([]chargebee.QuoteLineGroup, error) {
	return ListAllSiteQuoteLineGroups(client.Site(), id)
}

func ListAllSiteQuoteLineGroups(site, id string) ([]chargebee.QuoteLineGroup, error) {
	return chargebee.ListQuoteLineGroups(site, id)
}

func ListQuoteLineGroupsCreatedSince(id string, createdSince *time.Time) ([]chargebee.QuoteLineGroup, error) {
	return ListSiteQuoteLineGroupsCreatedSince(client.Site(), id, createdSince)
}

func ListSiteQuoteLineGroupsCreatedSince(site, id string, createdSince *time.Time) ([]chargebee.QuoteLineGroup, error) {
	return chargebee.ListQuoteLineGroupsCreatedSince(site, id, createdSince)
}

func ListQuoteLineGroupsUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.QuoteLineGroup, error) {
	return ListSiteQuoteLineGroupsUpdatedSince(client.Site(), id, updatedSince)
}

func ListSiteQuoteLineGroupsUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.QuoteLineGroup, error) {
	return chargebee.ListQuoteLineGroupsUpdatedSince(site, id, updatedSince)
}

func ListQuoteLineGroups(id string) ([]chargebee.QuoteLineGroup, string, error) {
	return ListSiteQuoteLineGroups(client.Site(), id)
}

func ListSiteQuoteLineGroups(site, id string) ([]chargebee.QuoteLineGroup, string, error) {
	return chargebee.ListQuoteLineGroupsPage(site, id, "")
}

func ListQuoteLineGroupsPage(id, offset string) ([]chargebee.QuoteLineGroup, string, error) {
	return ListSiteQuoteLineGroupsPage(client.Site(), id, offset)
}

func ListSiteQuoteLineGroupsPage(site, id string, offset string) ([]chargebee.QuoteLineGroup, string, error) {
	return chargebee.ListQuoteLineGroupsPage(site, id, offset)
}

func ListQuoteLineGroupsPageSortBy(id, offset string, field string) ([]chargebee.QuoteLineGroup, string, error) {
	return ListSiteQuoteLineGroupsPageSortBy(client.Site(), id, offset, field)
}

func ListSiteQuoteLineGroupsPageSortBy(site, id string, offset string, field string) ([]chargebee.QuoteLineGroup, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListQuoteLineGroupsPageSortBy(site, id, offset, sortBy)
}

// quotes.go

func ListAllQuotes() ([]chargebee.Quote, error) {
	return ListAllSiteQuotes(client.Site())
}

func ListAllSiteQuotes(site string) ([]chargebee.Quote, error) {
	return chargebee.ListQuotes(site)
}

func ListQuotesCreatedSince(createdSince *time.Time) ([]chargebee.Quote, error) {
	return ListSiteQuotesCreatedSince(client.Site(), createdSince)
}

func ListSiteQuotesCreatedSince(site string, createdSince *time.Time) ([]chargebee.Quote, error) {
	return chargebee.ListQuotesCreatedSince(site, createdSince)
}

func ListQuotesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Quote, error) {
	return ListSiteQuotesUpdatedSince(client.Site(), updatedSince)
}

func ListSiteQuotesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Quote, error) {
	return chargebee.ListQuotesUpdatedSince(site, updatedSince)
}

func ListQuotes() ([]chargebee.Quote, string, error) {
	return ListSiteQuotes(client.Site())
}

func ListSiteQuotes(site string) ([]chargebee.Quote, string, error) {
	return chargebee.ListQuotesPage(site, "")
}

func ListQuotesPage(offset string) ([]chargebee.Quote, string, error) {
	return ListSiteQuotesPage(client.Site(), offset)
}

func ListSiteQuotesPage(site string, offset string) ([]chargebee.Quote, string, error) {
	return chargebee.ListQuotesPage(site, offset)
}

func ListQuotesPageSortBy(offset string, field string) ([]chargebee.Quote, string, error) {
	return ListSiteQuotesPageSortBy(client.Site(), offset, field)
}

func ListSiteQuotesPageSortBy(site string, offset string, field string) ([]chargebee.Quote, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListQuotesPageSortBy(site, offset, sortBy)
}

func RetrieveQuoteById(id string) (*chargebee.Quote, error) {
	return RetrieveSiteQuoteById(client.Site(), id)
}

func RetrieveSiteQuoteById(site string, id string) (*chargebee.Quote, error) {
	return chargebee.RetrieveQuoteById(site, id)
}

// subscription_contract_terms.go

func ListAllContractTerms(id string) ([]chargebee.ContractTerm, error) {
	return ListAllSiteContractTerms(client.Site(), id)
}

func ListAllSiteContractTerms(site, id string) ([]chargebee.ContractTerm, error) {
	return chargebee.ListContractTerms(site, id)
}

func ListContractTermsCreatedSince(id string, createdSince *time.Time) ([]chargebee.ContractTerm, error) {
	return ListSiteContractTermsCreatedSince(client.Site(), id, createdSince)
}

func ListSiteContractTermsCreatedSince(site, id string, createdSince *time.Time) ([]chargebee.ContractTerm, error) {
	return chargebee.ListContractTermsCreatedSince(site, id, createdSince)
}

func ListContractTermsUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.ContractTerm, error) {
	return ListSiteContractTermsUpdatedSince(client.Site(), id, updatedSince)
}

func ListSiteContractTermsUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.ContractTerm, error) {
	return chargebee.ListContractTermsUpdatedSince(site, id, updatedSince)
}

func ListContractTerms(id string) ([]chargebee.ContractTerm, string, error) {
	return ListSiteContractTerms(client.Site(), id)
}

func ListSiteContractTerms(site, id string) ([]chargebee.ContractTerm, string, error) {
	return chargebee.ListContractTermsPage(site, id, "")
}

func ListContractTermsPage(id, offset string) ([]chargebee.ContractTerm, string, error) {
	return ListSiteContractTermsPage(client.Site(), id, offset)
}

func ListSiteContractTermsPage(site, id string, offset string) ([]chargebee.ContractTerm, string, error) {
	return chargebee.ListContractTermsPage(site, id, offset)
}

func ListContractTermsPageSortBy(id, offset string, field string) ([]chargebee.ContractTerm, string, error) {
	return ListSiteContractTermsPageSortBy(client.Site(), id, offset, field)
}

func ListSiteContractTermsPageSortBy(site, id string, offset string, field string) ([]chargebee.ContractTerm, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListContractTermsPageSortBy(site, id, offset, sortBy)
}

// subscription_discounts.go

func ListAllDiscounts(id string) ([]chargebee.Discount, error) {
	return ListAllSiteDiscounts(client.Site(), id)
}

func ListAllSiteDiscounts(site, id string) ([]chargebee.Discount, error) {
	return chargebee.ListDiscounts(site, id)
}

func ListDiscountsCreatedSince(id string, createdSince *time.Time) ([]chargebee.Discount, error) {
	return ListSiteDiscountsCreatedSince(client.Site(), id, createdSince)
}

func ListSiteDiscountsCreatedSince(site, id string, createdSince *time.Time) ([]chargebee.Discount, error) {
	return chargebee.ListDiscountsCreatedSince(site, id, createdSince)
}

func ListDiscountsUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.Discount, error) {
	return ListSiteDiscountsUpdatedSince(client.Site(), id, updatedSince)
}

func ListSiteDiscountsUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.Discount, error) {
	return chargebee.ListDiscountsUpdatedSince(site, id, updatedSince)
}

func ListDiscounts(id string) ([]chargebee.Discount, string, error) {
	return ListSiteDiscounts(client.Site(), id)
}

func ListSiteDiscounts(site, id string) ([]chargebee.Discount, string, error) {
	return chargebee.ListDiscountsPage(site, id, "")
}

func ListDiscountsPage(id, offset string) ([]chargebee.Discount, string, error) {
	return ListSiteDiscountsPage(client.Site(), id, offset)
}

func ListSiteDiscountsPage(site, id string, offset string) ([]chargebee.Discount, string, error) {
	return chargebee.ListDiscountsPage(site, id, offset)
}

func ListDiscountsPageSortBy(id, offset string, field string) ([]chargebee.Discount, string, error) {
	return ListSiteDiscountsPageSortBy(client.Site(), id, offset, field)
}

func ListSiteDiscountsPageSortBy(site, id string, offset string, field string) ([]chargebee.Discount, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListDiscountsPageSortBy(site, id, offset, sortBy)
}

// subscription_entitlements.go

func ListAllSubscriptionEntitlements(id string) ([]chargebee.SubscriptionEntitlement, error) {
	return ListAllSiteSubscriptionEntitlements(client.Site(), id)
}

func ListAllSiteSubscriptionEntitlements(site, id string) ([]chargebee.SubscriptionEntitlement, error) {
	return chargebee.ListSubscriptionEntitlements(site, id)
}

func ListSubscriptionEntitlementsCreatedSince(id string, createdSince *time.Time) ([]chargebee.SubscriptionEntitlement, error) {
	return ListSiteSubscriptionEntitlementsCreatedSince(client.Site(), id, createdSince)
}

func ListSiteSubscriptionEntitlementsCreatedSince(site, id string, createdSince *time.Time) ([]chargebee.SubscriptionEntitlement, error) {
	return chargebee.ListSubscriptionEntitlementsCreatedSince(site, id, createdSince)
}

func ListSubscriptionEntitlementsUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.SubscriptionEntitlement, error) {
	return ListSiteSubscriptionEntitlementsUpdatedSince(client.Site(), id, updatedSince)
}

func ListSiteSubscriptionEntitlementsUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.SubscriptionEntitlement, error) {
	return chargebee.ListSubscriptionEntitlementsUpdatedSince(site, id, updatedSince)
}

func ListSubscriptionEntitlements(id string) ([]chargebee.SubscriptionEntitlement, string, error) {
	return ListSiteSubscriptionEntitlements(client.Site(), id)
}

func ListSiteSubscriptionEntitlements(site, id string) ([]chargebee.SubscriptionEntitlement, string, error) {
	return chargebee.ListSubscriptionEntitlementsPage(site, id, "")
}

func ListSubscriptionEntitlementsPage(id, offset string) ([]chargebee.SubscriptionEntitlement, string, error) {
	return ListSiteSubscriptionEntitlementsPage(client.Site(), id, offset)
}

func ListSiteSubscriptionEntitlementsPage(site, id string, offset string) ([]chargebee.SubscriptionEntitlement, string, error) {
	return chargebee.ListSubscriptionEntitlementsPage(site, id, offset)
}

func ListSubscriptionEntitlementsPageSortBy(id, offset string, field string) ([]chargebee.SubscriptionEntitlement, string, error) {
	return ListSiteSubscriptionEntitlementsPageSortBy(client.Site(), id, offset, field)
}

func ListSiteSubscriptionEntitlementsPageSortBy(site, id string, offset string, field string) ([]chargebee.SubscriptionEntitlement, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListSubscriptionEntitlementsPageSortBy(site, id, offset, sortBy)
}

// subscription_gifts.go

func ListAllGifts() ([]chargebee.Gift, error) {
	return ListAllSiteGifts(client.Site())
}

func ListAllSiteGifts(site string) ([]chargebee.Gift, error) {
	return chargebee.ListGifts(site)
}

func ListGiftsCreatedSince(createdSince *time.Time) ([]chargebee.Gift, error) {
	return ListSiteGiftsCreatedSince(client.Site(), createdSince)
}

func ListSiteGiftsCreatedSince(site string, createdSince *time.Time) ([]chargebee.Gift, error) {
	return chargebee.ListGiftsCreatedSince(site, createdSince)
}

func ListGiftsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Gift, error) {
	return ListSiteGiftsUpdatedSince(client.Site(), updatedSince)
}

func ListSiteGiftsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Gift, error) {
	return chargebee.ListGiftsUpdatedSince(site, updatedSince)
}

func ListGifts() ([]chargebee.Gift, string, error) {
	return ListSiteGifts(client.Site())
}

func ListSiteGifts(site string) ([]chargebee.Gift, string, error) {
	return chargebee.ListGiftsPage(site, "")
}

func ListGiftsPage(offset string) ([]chargebee.Gift, string, error) {
	return ListSiteGiftsPage(client.Site(), offset)
}

func ListSiteGiftsPage(site string, offset string) ([]chargebee.Gift, string, error) {
	return chargebee.ListGiftsPage(site, offset)
}

func ListGiftsPageSortBy(offset string, field string) ([]chargebee.Gift, string, error) {
	return ListSiteGiftsPageSortBy(client.Site(), offset, field)
}

func ListSiteGiftsPageSortBy(site string, offset string, field string) ([]chargebee.Gift, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListGiftsPageSortBy(site, offset, sortBy)
}

func RetrieveGiftById(id string) (*chargebee.Gift, error) {
	return RetrieveSiteGiftById(client.Site(), id)
}

func RetrieveSiteGiftById(site string, id string) (*chargebee.Gift, error) {
	return chargebee.RetrieveGiftById(site, id)
}

// subscription_invoice_schedules.go

func ListAllAdvancedInvoiceSchedules(id string) ([]chargebee.AdvancedInvoiceSchedule, error) {
	return ListAllSiteAdvancedInvoiceSchedules(client.Site(), id)
}

func ListAllSiteAdvancedInvoiceSchedules(site, id string) ([]chargebee.AdvancedInvoiceSchedule, error) {
	return chargebee.ListAdvancedInvoiceSchedules(site, id)
}

func ListAdvancedInvoiceSchedulesCreatedSince(id string, createdSince *time.Time) ([]chargebee.AdvancedInvoiceSchedule, error) {
	return ListSiteAdvancedInvoiceSchedulesCreatedSince(client.Site(), id, createdSince)
}

func ListSiteAdvancedInvoiceSchedulesCreatedSince(site, id string, createdSince *time.Time) ([]chargebee.AdvancedInvoiceSchedule, error) {
	return chargebee.ListAdvancedInvoiceSchedulesCreatedSince(site, id, createdSince)
}

func ListAdvancedInvoiceSchedulesUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.AdvancedInvoiceSchedule, error) {
	return ListSiteAdvancedInvoiceSchedulesUpdatedSince(client.Site(), id, updatedSince)
}

func ListSiteAdvancedInvoiceSchedulesUpdatedSince(site, id string, updatedSince *time.Time) ([]chargebee.AdvancedInvoiceSchedule, error) {
	return chargebee.ListAdvancedInvoiceSchedulesUpdatedSince(site, id, updatedSince)
}

func ListAdvancedInvoiceSchedules(id string) ([]chargebee.AdvancedInvoiceSchedule, string, error) {
	return ListSiteAdvancedInvoiceSchedules(client.Site(), id)
}

func ListSiteAdvancedInvoiceSchedules(site, id string) ([]chargebee.AdvancedInvoiceSchedule, string, error) {
	return chargebee.ListAdvancedInvoiceSchedulesPage(site, id, "")
}

func ListAdvancedInvoiceSchedulesPage(id, offset string) ([]chargebee.AdvancedInvoiceSchedule, string, error) {
	return ListSiteAdvancedInvoiceSchedulesPage(client.Site(), id, offset)
}

func ListSiteAdvancedInvoiceSchedulesPage(site, id string, offset string) ([]chargebee.AdvancedInvoiceSchedule, string, error) {
	return chargebee.ListAdvancedInvoiceSchedulesPage(site, id, offset)
}

func ListAdvancedInvoiceSchedulesPageSortBy(id, offset string, field string) ([]chargebee.AdvancedInvoiceSchedule, string, error) {
	return ListSiteAdvancedInvoiceSchedulesPageSortBy(client.Site(), id, offset, field)
}

func ListSiteAdvancedInvoiceSchedulesPageSortBy(site, id string, offset string, field string) ([]chargebee.AdvancedInvoiceSchedule, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListAdvancedInvoiceSchedulesPageSortBy(site, id, offset, sortBy)
}

// subscription_ramps.go

func ListAllRamps() ([]chargebee.Ramp, error) {
	return ListAllSiteRamps(client.Site())
}

func ListAllSiteRamps(site string) ([]chargebee.Ramp, error) {
	return chargebee.ListRamps(site)
}

func ListRampsCreatedSince(createdSince *time.Time) ([]chargebee.Ramp, error) {
	return ListSiteRampsCreatedSince(client.Site(), createdSince)
}

func ListSiteRampsCreatedSince(site string, createdSince *time.Time) ([]chargebee.Ramp, error) {
	return chargebee.ListRampsCreatedSince(site, createdSince)
}

func ListRampsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Ramp, error) {
	return ListSiteRampsUpdatedSince(client.Site(), updatedSince)
}

func ListSiteRampsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Ramp, error) {
	return chargebee.ListRampsUpdatedSince(site, updatedSince)
}

func ListRamps() ([]chargebee.Ramp, string, error) {
	return ListSiteRamps(client.Site())
}

func ListSiteRamps(site string) ([]chargebee.Ramp, string, error) {
	return chargebee.ListRampsPage(site, "")
}

func ListRampsPage(offset string) ([]chargebee.Ramp, string, error) {
	return ListSiteRampsPage(client.Site(), offset)
}

func ListSiteRampsPage(site string, offset string) ([]chargebee.Ramp, string, error) {
	return chargebee.ListRampsPage(site, offset)
}

func ListRampsPageSortBy(offset string, field string) ([]chargebee.Ramp, string, error) {
	return ListSiteRampsPageSortBy(client.Site(), offset, field)
}

func ListSiteRampsPageSortBy(site string, offset string, field string) ([]chargebee.Ramp, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListRampsPageSortBy(site, offset, sortBy)
}

// subscription_usages.go

func ListAllUsages() ([]chargebee.Usage, error) {
	return ListAllSiteUsages(client.Site())
}

func ListAllSiteUsages(site string) ([]chargebee.Usage, error) {
	return chargebee.ListUsages(site)
}

func ListUsagesCreatedSince(createdSince *time.Time) ([]chargebee.Usage, error) {
	return ListSiteUsagesCreatedSince(client.Site(), createdSince)
}

func ListSiteUsagesCreatedSince(site string, createdSince *time.Time) ([]chargebee.Usage, error) {
	return chargebee.ListUsagesCreatedSince(site, createdSince)
}

func ListUsagesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Usage, error) {
	return ListSiteUsagesUpdatedSince(client.Site(), updatedSince)
}

func ListSiteUsagesUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Usage, error) {
	return chargebee.ListUsagesUpdatedSince(site, updatedSince)
}

func ListUsages() ([]chargebee.Usage, string, error) {
	return ListSiteUsages(client.Site())
}

func ListSiteUsages(site string) ([]chargebee.Usage, string, error) {
	return chargebee.ListUsagesPage(site, "")
}

func ListUsagesPage(offset string) ([]chargebee.Usage, string, error) {
	return ListSiteUsagesPage(client.Site(), offset)
}

func ListSiteUsagesPage(site string, offset string) ([]chargebee.Usage, string, error) {
	return chargebee.ListUsagesPage(site, offset)
}

func ListUsagesPageSortBy(offset string, field string) ([]chargebee.Usage, string, error) {
	return ListSiteUsagesPageSortBy(client.Site(), offset, field)
}

func ListSiteUsagesPageSortBy(site string, offset string, field string) ([]chargebee.Usage, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListUsagesPageSortBy(site, offset, sortBy)
}

// subscriptions.go

func ListAllSubscriptions() ([]chargebee.Subscription, error) {
	return ListAllSiteSubscriptions(client.Site())
}

func ListAllSiteSubscriptions(site string) ([]chargebee.Subscription, error) {
	return chargebee.ListSubscriptions(site)
}

func ListSubscriptionsCreatedSince(createdSince *time.Time) ([]chargebee.Subscription, error) {
	return ListSiteSubscriptionsCreatedSince(client.Site(), createdSince)
}

func ListSiteSubscriptionsCreatedSince(site string, createdSince *time.Time) ([]chargebee.Subscription, error) {
	return chargebee.ListSubscriptionsCreatedSince(site, createdSince)
}

func ListSubscriptionsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Subscription, error) {
	return ListSiteSubscriptionsUpdatedSince(client.Site(), updatedSince)
}

func ListSiteSubscriptionsUpdatedSince(site string, updatedSince *time.Time) ([]chargebee.Subscription, error) {
	return chargebee.ListSubscriptionsUpdatedSince(site, updatedSince)
}

func ListSubscriptions() ([]chargebee.Subscription, string, error) {
	return ListSiteSubscriptions(client.Site())
}

func ListSiteSubscriptions(site string) ([]chargebee.Subscription, string, error) {
	return chargebee.ListSubscriptionsPage(site, "")
}

func ListSubscriptionsPage(offset string) ([]chargebee.Subscription, string, error) {
	return ListSiteSubscriptionsPage(client.Site(), offset)
}

func ListSiteSubscriptionsPage(site string, offset string) ([]chargebee.Subscription, string, error) {
	return chargebee.ListSubscriptionsPage(site, offset)
}

func ListSubscriptionsPageSortBy(offset string, field string) ([]chargebee.Subscription, string, error) {
	return ListSiteSubscriptionsPageSortBy(client.Site(), offset, field)
}

func ListSiteSubscriptionsPageSortBy(site string, offset string, field string) ([]chargebee.Subscription, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.ListSubscriptionsPageSortBy(site, offset, sortBy)
}
