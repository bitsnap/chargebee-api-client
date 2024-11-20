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

type Quote struct {
    Id string `json:"id" validate:"required"`
    Name string `json:"name"`
    PoNumber string `json:"po_number"`
    CustomerId string `json:"customer_id" validate:"required"`
    SubscriptionId string `json:"subscription_id"`
    InvoiceId string `json:"invoice_id"`
    Status enums.StatusEnum `json:"status" validate:"required"`
    OperationType enums.OperationTypeEnum `json:"operation_type" validate:"required"`
    VatNumber string `json:"vat_number"`
    PriceType enums.PriceTypeEnum `json:"price_type" validate:"required"`
    ValidTill uint64 `json:"valid_till" validate:"required"`
    Date uint64 `json:"date" validate:"required"`
    TotalPayable uint64 `json:"total_payable"`
    ChargeOnAcceptance uint64 `json:"charge_on_acceptance"`
    SubTotal uint64 `json:"sub_total" validate:"required"`
    Total uint64 `json:"total"`
    CreditsApplied uint64 `json:"credits_applied"`
    AmountPaid uint64 `json:"amount_paid"`
    AmountDue uint64 `json:"amount_due"`
    Version int `json:"version"`
    ResourceVersion int64 `json:"resource_version"`
    UpdatedAt uint64 `json:"updated_at"`
    VatNumberPrefix string `json:"vat_number_prefix"`
    TaxCategory string `json:"tax_category"`
    CurrencyCode string `json:"currency_code" validate:"required"`
    Notes []string `json:"notes"`
    ContractTermStart uint64 `json:"contract_term_start"`
    ContractTermEnd uint64 `json:"contract_term_end"`
    ContractTermTerminationFee uint64 `json:"contract_term_termination_fee"`
    BusinessEntityId string `json:"business_entity_id"`
    LineItems []models.LineItem `json:"line_items"`
    Discounts []models.Discount `json:"discounts"`
    LineItemDiscounts []models.LineItemDiscount `json:"line_item_discounts"`
    Taxes []models.Tax `json:"taxes"`
    LineItemTaxes []models.LineItemTax `json:"line_item_taxes"`
    LineItemTiers []models.LineItemTier `json:"line_item_tiers"`
    ShippingAddress models.ShippingAddress `json:"shipping_address"`
    BillingAddress models.BillingAddress `json:"billing_address"`
}

func ListQuotesPageSortBy(site string, offset string, sortBy *SortBy) ([]Quote, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/quotes")
	if err != nil {
		return nil, "", err
	}
	
    content, err := GetQuery(client, parsedUrl, offset, sortBy)
    if err != nil {
        return nil, "", err
    }
    		
	type QuoteListItem struct {
		Quote Quote `json:"Quote"`
	}

    type QuotePage struct {
        List       []QuoteListItem `json:"list"`
        NextOffset string `json:"next_offset,omitempty"`
    }

	entries := QuotePage{
		List:       make([]QuoteListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}
	
	if len(entries.List) == 0 {
        return []Quote{}, "", nil
    }
	
	result := make([]Quote, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Quote)
	}

    if len(entries.NextOffset) > 0 {
        return ResultWithOffset(result, offset, entries.NextOffset)
    }
	
	return result, "", nil
}

func ListQuotesPage(site string, offset string) ([]Quote, string, error) {
	return ListQuotesPageSortBy(site, offset, nil)
}

func RetrieveQuoteById(site string, id string) (*Quote, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/quotes/" + id)
	if err != nil {
		return nil, err
	}
	
    content, err := GetQuery(client, parsedUrl, "", nil)
    if err != nil {
        return nil, err
    }
    	
	type QuoteItem struct {
		Quote Quote `json:"Quote"`
	}

    var item QuoteItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.Quote, nil
}