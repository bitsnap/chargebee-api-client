package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	"github.com/bitsnap/chargebee-api-client/chargebee/enums"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Usage struct {
	Id              string           `json:"id"`
	UsageDate       uint64           `json:"usage_date" validate:"required"`
	SubscriptionId  string           `json:"subscription_id" validate:"required"`
	ItemPriceId     string           `json:"item_price_id" validate:"required"`
	InvoiceId       string           `json:"invoice_id"`
	LineItemId      string           `json:"line_item_id"`
	Quantity        string           `json:"quantity" validate:"required"`
	Source          enums.SourceEnum `json:"source"`
	Note            string           `json:"note"`
	ResourceVersion int64            `json:"resource_version"`
	UpdatedAt       uint64           `json:"updated_at"`
	CreatedAt       uint64           `json:"created_at" validate:"required"`
}

func ListUsagesPageSortBy(site string, offset string, sortBy *SortBy) ([]Usage, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/usages")
	if err != nil {
		return nil, "", err
	}

	content, err := getQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type UsageListItem struct {
		Usage Usage `json:"Usage"`
	}

	type UsagePage struct {
		List       []UsageListItem `json:"list"`
		NextOffset string          `json:"next_offset,omitempty"`
	}

	entries := UsagePage{
		List: make([]UsageListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []Usage{}, "", nil
	}

	result := make([]Usage, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Usage)
	}

	if len(entries.NextOffset) > 0 {
		return resultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListUsagesPage(site string, offset string) ([]Usage, string, error) {
	return ListUsagesPageSortBy(site, offset, nil)
}

func ListSubscriptionUsagesPageSortBy(site string, id string, offset string, sortBy *SortBy) ([]Usage, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/subscription/" + id + "/usages")
	if err != nil {
		return nil, "", err
	}

	content, err := getQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type UsageListItem struct {
		Usage Usage `json:"Usage"`
	}

	type UsagePage struct {
		List       []UsageListItem `json:"list"`
		NextOffset string          `json:"next_offset,omitempty"`
	}

	entries := UsagePage{
		List: make([]UsageListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []Usage{}, "", nil
	}

	result := make([]Usage, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Usage)
	}

	if len(entries.NextOffset) > 0 {
		return resultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListSubscriptionUsagesPage(site string, id string, offset string) ([]Usage, string, error) {
	return ListSubscriptionUsagesPageSortBy(site, id, offset, nil)
}
