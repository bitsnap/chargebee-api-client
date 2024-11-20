package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/bitsnap/chargebee-api-client/chargebee/enums"
	"github.com/bitsnap/chargebee-api-client/chargebee/models"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

type Customer struct {
    Id string `json:"id" validate:"required"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    Email string `json:"email"`
    Phone string `json:"phone"`
    Company string `json:"company"`
    VatNumber string `json:"vat_number"`
    AutoCollection enums.AutoCollectionEnum `json:"auto_collection" validate:"required"`
    OfflinePaymentMethod enums.OfflinePaymentMethodEnum `json:"offline_payment_method"`
    NetTermDays int `json:"net_term_days" validate:"required"`
    VatNumberValidatedTime uint64 `json:"vat_number_validated_time"`
    VatNumberStatus enums.VatNumberStatusEnum `json:"vat_number_status"`
    AllowDirectDebit bool `json:"allow_direct_debit" validate:"required"`
    IsLocationValid bool `json:"is_location_valid"`
    CreatedAt uint64 `json:"created_at" validate:"required"`
    CreatedFromIp string `json:"created_from_ip"`
    ExemptionDetails []string `json:"exemption_details"`
    Taxability enums.TaxabilityEnum `json:"taxability"`
    EntityCode enums.EntityCodeEnum `json:"entity_code"`
    ExemptNumber string `json:"exempt_number"`
    ResourceVersion int64 `json:"resource_version"`
    UpdatedAt uint64 `json:"updated_at"`
    Locale string `json:"locale"`
    BillingDate int `json:"billing_date"`
    BillingMonth int `json:"billing_month"`
    BillingDateMode enums.BillingDateModeEnum `json:"billing_date_mode"`
    BillingDayOfWeek enums.BillingDayOfWeekEnum `json:"billing_day_of_week"`
    BillingDayOfWeekMode enums.BillingDayOfWeekModeEnum `json:"billing_day_of_week_mode"`
    PiiCleared enums.PiiClearedEnum `json:"pii_cleared"`
    AutoCloseInvoices bool `json:"auto_close_invoices"`
    Channel enums.ChannelEnum `json:"channel"`
    ActiveId string `json:"active_id"`
    FraudFlag enums.FraudFlagEnum `json:"fraud_flag"`
    PrimaryPaymentSourceId string `json:"primary_payment_source_id"`
    BackupPaymentSourceId string `json:"backup_payment_source_id"`
    InvoiceNotes string `json:"invoice_notes"`
    BusinessEntityId string `json:"business_entity_id"`
    PreferredCurrencyCode string `json:"preferred_currency_code"`
    PromotionalCredits uint64 `json:"promotional_credits" validate:"required"`
    UnbilledCharges uint64 `json:"unbilled_charges" validate:"required"`
    RefundableCredits uint64 `json:"refundable_credits" validate:"required"`
    ExcessPayments uint64 `json:"excess_payments" validate:"required"`
    IsEinvoiceEnabled bool `json:"is_einvoice_enabled"`
    EinvoicingMethod enums.EinvoicingMethodEnum `json:"einvoicing_method"`
    Deleted bool `json:"deleted" validate:"required"`
    RegisteredForGst bool `json:"registered_for_gst"`
    ConsolidatedInvoicing bool `json:"consolidated_invoicing"`
    CustomerType enums.CustomerTypeEnum `json:"customer_type"`
    BusinessCustomerWithoutVatNumber bool `json:"business_customer_without_vat_number"`
    ClientProfileId string `json:"client_profile_id"`
    UseDefaultHierarchySettings bool `json:"use_default_hierarchy_settings"`
    VatNumberPrefix string `json:"vat_number_prefix"`
    EntityIdentifierScheme string `json:"entity_identifier_scheme"`
    EntityIdentifierStandard string `json:"entity_identifier_standard"`
    BillingAddress models.BillingAddress `json:"billing_address"`
    ReferralUrls []models.ReferralUrl `json:"referral_urls"`
    Contacts []models.Contact `json:"contacts"`
    PaymentMethod models.PaymentMethod `json:"payment_method"`
    Balances []models.CustomerBalance `json:"balances"`
    EntityIdentifiers []models.EntityIdentifier `json:"entity_identifiers"`
    TaxProvidersFields []models.TaxProvidersField `json:"tax_providers_fields"`
    Relationship models.Relationship `json:"relationship"`
    ParentAccountAccess models.ParentAccountAccess `json:"parent_account_access"`
    ChildAccountAccess models.ChildAccountAccess `json:"child_account_access"`
}

func ListCustomersPageSortBy(site string, offset string, sortBy *SortBy) ([]Customer, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/customers")
	if err != nil {
		return nil, "", err
	}
	
    content, err := GetQuery(client, parsedUrl, offset, sortBy)
    if err != nil {
        return nil, "", err
    }
    		
	type CustomerListItem struct {
		Customer Customer `json:"Customer"`
	}

    type CustomerPage struct {
        List       []CustomerListItem `json:"list"`
        NextOffset string `json:"next_offset,omitempty"`
    }

	entries := CustomerPage{
		List:       make([]CustomerListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}
	
	if len(entries.List) == 0 {
        return []Customer{}, "", nil
    }
	
	result := make([]Customer, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Customer)
	}

    if len(entries.NextOffset) > 0 {
        return ResultWithOffset(result, offset, entries.NextOffset)
    }
	
	return result, "", nil
}

func ListCustomersPage(site string, offset string) ([]Customer, string, error) {
	return ListCustomersPageSortBy(site, offset, nil)
}

func RetrieveCustomerById(site string, id string) (*Customer, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/customers/" + id)
	if err != nil {
		return nil, err
	}
	
    content, err := GetQuery(client, parsedUrl, "", nil)
    if err != nil {
        return nil, err
    }
    	
	type CustomerItem struct {
		Customer Customer `json:"Customer"`
	}

    var item CustomerItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.Customer, nil
}