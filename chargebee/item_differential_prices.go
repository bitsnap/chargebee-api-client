package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	"github.com/bitsnap/chargebee-api-client/chargebee/enums"
	"github.com/bitsnap/chargebee-api-client/chargebee/models"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type DifferentialPrice struct {
	Id              string                `json:"id" validate:"required"`
	ItemPriceId     string                `json:"item_price_id" validate:"required"`
	ParentItemId    string                `json:"parent_item_id" validate:"required"`
	Price           uint64                `json:"price"`
	PriceInDecimal  string                `json:"price_in_decimal"`
	Status          enums.StatusEnum      `json:"status"`
	ResourceVersion int64                 `json:"resource_version"`
	UpdatedAt       uint64                `json:"updated_at"`
	CreatedAt       uint64                `json:"created_at" validate:"required"`
	ModifiedAt      uint64                `json:"modified_at" validate:"required"`
	CurrencyCode    string                `json:"currency_code" validate:"required"`
	Tiers           []models.Tier         `json:"tiers"`
	ParentPeriods   []models.ParentPeriod `json:"parent_periods"`
}

func ListDifferentialPricesPageSortBy(site string, offset string, sortBy *SortBy) ([]DifferentialPrice, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/item_prices")
	if err != nil {
		return nil, "", err
	}

	content, err := getQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type DifferentialPriceListItem struct {
		DifferentialPrice DifferentialPrice `json:"DifferentialPrice"`
	}

	type DifferentialPricePage struct {
		List       []DifferentialPriceListItem `json:"list"`
		NextOffset string                      `json:"next_offset,omitempty"`
	}

	entries := DifferentialPricePage{
		List: make([]DifferentialPriceListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []DifferentialPrice{}, "", nil
	}

	result := make([]DifferentialPrice, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.DifferentialPrice)
	}

	if len(entries.NextOffset) > 0 {
		return resultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListDifferentialPricesPage(site string, offset string) ([]DifferentialPrice, string, error) {
	return ListDifferentialPricesPageSortBy(site, offset, nil)
}

func RetrieveDifferentialPriceById(site string, id string) (*DifferentialPrice, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/item_prices/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQuery(client, parsedUrl, "", nil)
	if err != nil {
		return nil, err
	}

	type DifferentialPriceItem struct {
		DifferentialPrice DifferentialPrice `json:"DifferentialPrice"`
	}

	var item DifferentialPriceItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.DifferentialPrice, nil
}
