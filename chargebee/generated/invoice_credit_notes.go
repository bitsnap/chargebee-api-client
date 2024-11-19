package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/bitsnap/chargebee-api-client/chargebee/generated/enums"
	"github.com/bitsnap/chargebee-api-client/chargebee/generated/models"
	"github.com/shopspring/decimal"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

type CreditNote struct {
    Id string `json:"id" validate:"required"`
    CustomerId string `json:"customer_id" validate:"required"`
    SubscriptionId string `json:"subscription_id"`
    ReferenceInvoiceId string `json:"reference_invoice_id"`
    Type enums.TypeEnum `json:"type" validate:"required"`
    ReasonCode enums.ReasonCodeEnum `json:"reason_code"`
    Status enums.StatusEnum `json:"status" validate:"required"`
    VatNumber string `json:"vat_number"`
    Date uint64 `json:"date"`
    PriceType enums.PriceTypeEnum `json:"price_type" validate:"required"`
    CurrencyCode string `json:"currency_code" validate:"required"`
    Total uint64 `json:"total"`
    AmountAllocated uint64 `json:"amount_allocated"`
    AmountRefunded uint64 `json:"amount_refunded"`
    AmountAvailable uint64 `json:"amount_available"`
    RefundedAt uint64 `json:"refunded_at"`
    VoidedAt uint64 `json:"voided_at"`
    GeneratedAt uint64 `json:"generated_at"`
    ResourceVersion int64 `json:"resource_version"`
    UpdatedAt uint64 `json:"updated_at"`
    Channel enums.ChannelEnum `json:"channel"`
    SubTotal uint64 `json:"sub_total" validate:"required"`
    SubTotalInLocalCurrency uint64 `json:"sub_total_in_local_currency"`
    TotalInLocalCurrency uint64 `json:"total_in_local_currency"`
    LocalCurrencyCode string `json:"local_currency_code"`
    RoundOffAmount uint64 `json:"round_off_amount"`
    FractionalCorrection uint64 `json:"fractional_correction"`
    Deleted bool `json:"deleted" validate:"required"`
    TaxCategory string `json:"tax_category"`
    LocalCurrencyExchangeRate *decimal.Decimal `json:"local_currency_exchange_rate"`
    CreateReasonCode string `json:"create_reason_code"`
    VatNumberPrefix string `json:"vat_number_prefix"`
    BusinessEntityId string `json:"business_entity_id"`
    Einvoice models.Einvoice `json:"einvoice"`
    LineItems []models.LineItem `json:"line_items"`
    Discounts []models.Discount `json:"discounts"`
    LineItemDiscounts []models.LineItemDiscount `json:"line_item_discounts"`
    LineItemTiers []models.LineItemTier `json:"line_item_tiers"`
    Taxes []models.Tax `json:"taxes"`
    LineItemTaxes []models.LineItemTax `json:"line_item_taxes"`
    LinkedRefunds []models.CreditNoteTransaction `json:"linked_refunds"`
    LinkedTaxWithheldRefunds []models.LinkedTaxWithheldRefund `json:"linked_tax_withheld_refunds"`
    Allocations []models.AppliedCredit `json:"allocations"`
    ShippingAddress models.ShippingAddress `json:"shipping_address"`
    BillingAddress models.BillingAddress `json:"billing_address"`
    SiteDetailsAtCreation models.SiteDetailsAtCreation `json:"site_details_at_creation"`
    TaxOrigin models.TaxOrigin `json:"tax_origin"`
}

func ListCreditNotesPageSortBy(site string, offset string, sortBy *SortBy) ([]CreditNote, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/credit_notes")
	if err != nil {
		return nil, "", err
	}
	
    content, err := GetQuery(client, parsedUrl, offset, sortBy)
    if err != nil {
        return nil, "", err
    }
    		
	type CreditNoteListItem struct {
		CreditNote CreditNote `json:"CreditNote"`
	}

    type CreditNotePage struct {
        List       []CreditNoteListItem `json:"list"`
        NextOffset string `json:"next_offset,omitempty"`
    }

	entries := CreditNotePage{
		List:       make([]CreditNoteListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}
	
	if len(entries.List) == 0 {
        return []CreditNote{}, "", nil
    }
	
	result := make([]CreditNote, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.CreditNote)
	}

    if len(entries.NextOffset) > 0 {
        return ResultWithOffset(result, offset, entries.NextOffset)
    }
	
	return result, "", nil
}

func ListCreditNotesPage(site string, offset string) ([]CreditNote, string, error) {
	return ListCreditNotesPageSortBy(site, offset, nil)
}

func RetrieveCreditNoteById(site string, id string) (*CreditNote, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/credit_notes/" + id)
	if err != nil {
		return nil, err
	}
	
    content, err := GetQuery(client, parsedUrl, "", nil)
    if err != nil {
        return nil, err
    }
    	
	type CreditNoteItem struct {
		CreditNote CreditNote `json:"CreditNote"`
	}

    var item CreditNoteItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.CreditNote, nil
}