package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/bitsnap/chargebee-api-client/chargebee/enums"
	"github.com/bitsnap/chargebee-api-client/chargebee/models"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type InvoiceUnbilledCharge struct {
	Id                  string                 `json:"id"`
	CustomerId          string                 `json:"customer_id"`
	SubscriptionId      string                 `json:"subscription_id"`
	DateFrom            uint64                 `json:"date_from"`
	DateTo              uint64                 `json:"date_to"`
	UnitAmount          uint64                 `json:"unit_amount"`
	PricingModel        enums.PricingModelEnum `json:"pricing_model"`
	Quantity            int                    `json:"quantity"`
	Amount              uint64                 `json:"amount"`
	CurrencyCode        string                 `json:"currency_code" validate:"required"`
	DiscountAmount      uint64                 `json:"discount_amount"`
	Description         string                 `json:"description"`
	EntityType          enums.EntityTypeEnum   `json:"entity_type" validate:"required"`
	EntityId            string                 `json:"entity_id"`
	IsVoided            bool                   `json:"is_voided" validate:"required"`
	VoidedAt            uint64                 `json:"voided_at"`
	UnitAmountInDecimal string                 `json:"unit_amount_in_decimal"`
	QuantityInDecimal   string                 `json:"quantity_in_decimal"`
	AmountInDecimal     string                 `json:"amount_in_decimal"`
	UpdatedAt           uint64                 `json:"updated_at" validate:"required"`
	IsAdvanceCharge     bool                   `json:"is_advance_charge"`
	BusinessEntityId    string                 `json:"business_entity_id"`
	Deleted             bool                   `json:"deleted" validate:"required"`
	Tiers               []models.LineItemTier  `json:"tiers"`
}

func ListInvoiceUnbilledChargesPageSortBy(site string, offset string, sortBy *SortBy) ([]InvoiceUnbilledCharge, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/unbilled_charges")
	if err != nil {
		return nil, "", err
	}

	content, err := GetQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type InvoiceUnbilledChargeListItem struct {
		InvoiceUnbilledCharge InvoiceUnbilledCharge `json:"InvoiceUnbilledCharge"`
	}

	type InvoiceUnbilledChargePage struct {
		List       []InvoiceUnbilledChargeListItem `json:"list"`
		NextOffset string                          `json:"next_offset,omitempty"`
	}

	entries := InvoiceUnbilledChargePage{
		List: make([]InvoiceUnbilledChargeListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []InvoiceUnbilledCharge{}, "", nil
	}

	result := make([]InvoiceUnbilledCharge, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.InvoiceUnbilledCharge)
	}

	if len(entries.NextOffset) > 0 {
		return ResultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListInvoiceUnbilledChargesPage(site string, offset string) ([]InvoiceUnbilledCharge, string, error) {
	return ListInvoiceUnbilledChargesPageSortBy(site, offset, nil)
}
