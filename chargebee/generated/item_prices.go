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

type ItemPrice struct {
	Id                              string                                    `json:"id" validate:"required"`
	Name                            string                                    `json:"name" validate:"required"`
	ItemFamilyId                    string                                    `json:"item_family_id"`
	ItemId                          string                                    `json:"item_id"`
	Description                     string                                    `json:"description"`
	Status                          enums.StatusEnum                          `json:"status"`
	ExternalName                    string                                    `json:"external_name"`
	PriceVariantId                  string                                    `json:"price_variant_id"`
	ProrationType                   enums.ProrationTypeEnum                   `json:"proration_type"`
	PricingModel                    enums.PricingModelEnum                    `json:"pricing_model" validate:"required"`
	Price                           uint64                                    `json:"price"`
	PriceInDecimal                  string                                    `json:"price_in_decimal"`
	Period                          int                                       `json:"period"`
	CurrencyCode                    string                                    `json:"currency_code" validate:"required"`
	PeriodUnit                      enums.PeriodUnitEnum                      `json:"period_unit"`
	TrialPeriod                     int                                       `json:"trial_period"`
	TrialPeriodUnit                 enums.TrialPeriodUnitEnum                 `json:"trial_period_unit"`
	TrialEndAction                  enums.TrialEndActionEnum                  `json:"trial_end_action"`
	ShippingPeriod                  int                                       `json:"shipping_period"`
	ShippingPeriodUnit              enums.ShippingPeriodUnitEnum              `json:"shipping_period_unit"`
	BillingCycles                   int                                       `json:"billing_cycles"`
	FreeQuantity                    int                                       `json:"free_quantity" validate:"required"`
	FreeQuantityInDecimal           string                                    `json:"free_quantity_in_decimal"`
	Channel                         enums.ChannelEnum                         `json:"channel"`
	ResourceVersion                 int64                                     `json:"resource_version"`
	UpdatedAt                       uint64                                    `json:"updated_at"`
	CreatedAt                       uint64                                    `json:"created_at" validate:"required"`
	UsageAccumulationResetFrequency enums.UsageAccumulationResetFrequencyEnum `json:"usage_accumulation_reset_frequency"`
	ArchivedAt                      uint64                                    `json:"archived_at"`
	InvoiceNotes                    string                                    `json:"invoice_notes"`
	IsTaxable                       bool                                      `json:"is_taxable"`
	Metadata                        map[string]any                            `json:"metadata"`
	ItemType                        enums.ItemTypeEnum                        `json:"item_type"`
	ShowDescriptionInInvoices       bool                                      `json:"show_description_in_invoices"`
	ShowDescriptionInQuotes         bool                                      `json:"show_description_in_quotes"`
	BusinessEntityId                string                                    `json:"business_entity_id"`
	Tiers                           []models.Tier                             `json:"tiers"`
	TaxDetail                       models.TaxDetail                          `json:"tax_detail"`
	TaxProvidersFields              []models.TaxProvidersField                `json:"tax_providers_fields"`
	AccountingDetail                models.AccountingDetail                   `json:"accounting_detail"`
}

func ListItemPricesPageSortBy(site string, offset string, sortBy *SortBy) ([]ItemPrice, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/item_prices")
	if err != nil {
		return nil, "", err
	}

	content, err := GetQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type ItemPriceListItem struct {
		ItemPrice ItemPrice `json:"ItemPrice"`
	}

	type ItemPricePage struct {
		List       []ItemPriceListItem `json:"list"`
		NextOffset string              `json:"next_offset,omitempty"`
	}

	entries := ItemPricePage{
		List: make([]ItemPriceListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []ItemPrice{}, "", nil
	}

	result := make([]ItemPrice, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.ItemPrice)
	}

	if len(entries.NextOffset) > 0 {
		return ResultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListItemPricesPage(site string, offset string) ([]ItemPrice, string, error) {
	return ListItemPricesPageSortBy(site, offset, nil)
}

func RetrieveItemPriceById(site string, id string) (*ItemPrice, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/item_prices/" + id)
	if err != nil {
		return nil, err
	}

	content, err := GetQuery(client, parsedUrl, "", nil)
	if err != nil {
		return nil, err
	}

	type ItemPriceItem struct {
		ItemPrice ItemPrice `json:"ItemPrice"`
	}

	var item ItemPriceItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.ItemPrice, nil
}
