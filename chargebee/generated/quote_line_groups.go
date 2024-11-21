package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/bitsnap/chargebee-api-client/chargebee/generated/enums"
	"github.com/bitsnap/chargebee-api-client/chargebee/generated/models"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

func ListQuoteLineGroupsPageSortBy(site string, id string, offset string, sortBy *SortBy) ([]QuoteLineGroup, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/quotes/" + id + "/quote_line_groups")
	if err != nil {
		return nil, "", err
	}
		
    content, err := GetQuery(client, parsedUrl, offset, sortBy)
    if err != nil {
        return nil, "", err
    }
    	
	type QuoteLineGroupListItem struct {
		QuoteLineGroup QuoteLineGroup `json:"QuoteLineGroup"`
	}

    type QuoteLineGroupPage struct {
        List       []QuoteLineGroupListItem `json:"list"`
        NextOffset string `json:"next_offset,omitempty"`
    }

	entries := QuoteLineGroupPage{
		List:       make([]QuoteLineGroupListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}
	
	if len(entries.List) == 0 {
        return []QuoteLineGroup{}, "", nil
    }
	
	result := make([]QuoteLineGroup, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.QuoteLineGroup)
	}

    if len(entries.NextOffset) > 0 {
        return ResultWithOffset(result, offset, entries.NextOffset)
    }
	
	return result, "", nil
}

func ListQuoteLineGroupsPage(site string, id string, offset string) ([]QuoteLineGroup, string, error) {
	return ListQuoteLineGroupsPageSortBy(site, id, offset, nil)
}

func ListAllQuoteLineGroupsPageSortBy(site string, offset string, sortBy *SortBy) ([]QuoteLineGroup, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/quote_line_groups")
	if err != nil {
		return nil, "", err
	}
	
    content, err := GetQuery(client, parsedUrl, offset, sortBy)
    if err != nil {
        return nil, "", err
    }
    		
	type QuoteLineGroupListItem struct {
		QuoteLineGroup QuoteLineGroup `json:"QuoteLineGroup"`
	}

    type QuoteLineGroupPage struct {
        List       []QuoteLineGroupListItem `json:"list"`
        NextOffset string `json:"next_offset,omitempty"`
    }

	entries := QuoteLineGroupPage{
		List:       make([]QuoteLineGroupListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}
	
	if len(entries.List) == 0 {
        return []QuoteLineGroup{}, "", nil
    }
	
	result := make([]QuoteLineGroup, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.QuoteLineGroup)
	}

    if len(entries.NextOffset) > 0 {
        return ResultWithOffset(result, offset, entries.NextOffset)
    }
	
	return result, "", nil
}

func ListAllQuoteLineGroupsPage(site string, offset string) ([]QuoteLineGroup, string, error) {
	return ListAllQuoteLineGroupsPageSortBy(site, offset, nil)
}

type QuoteLineGroup struct {
    Version int `json:"version"`
    Id string `json:"id"`
    SubTotal uint64 `json:"sub_total" validate:"required"`
    Total uint64 `json:"total"`
    CreditsApplied uint64 `json:"credits_applied"`
    AmountPaid uint64 `json:"amount_paid"`
    AmountDue uint64 `json:"amount_due"`
    ChargeEvent enums.ChargeEventEnum `json:"charge_event"`
    BillingCycleNumber int `json:"billing_cycle_number"`
    LineItems []models.LineItem `json:"line_items"`
    Discounts []models.Discount `json:"discounts"`
    LineItemDiscounts []models.LineItemDiscount `json:"line_item_discounts"`
    Taxes []models.Tax `json:"taxes"`
    LineItemTaxes []models.LineItemTax `json:"line_item_taxes"`
}