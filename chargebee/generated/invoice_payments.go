package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

func ListInvoicePaymentsPageSortBy(site string, id string, offset string, sortBy *SortBy) ([]Transaction, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/invoices/" + id + "/payments")
	if err != nil {
		return nil, "", err
	}
		
    content, err := GetQuery(client, parsedUrl, offset, sortBy)
    if err != nil {
        return nil, "", err
    }
    	
	type TransactionListItem struct {
		Transaction Transaction `json:"Transaction"`
	}

    type TransactionPage struct {
        List       []TransactionListItem `json:"list"`
        NextOffset string `json:"next_offset,omitempty"`
    }

	entries := TransactionPage{
		List:       make([]TransactionListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}
	
	if len(entries.List) == 0 {
        return []Transaction{}, "", nil
    }
	
	result := make([]Transaction, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Transaction)
	}

    if len(entries.NextOffset) > 0 {
        return ResultWithOffset(result, offset, entries.NextOffset)
    }
	
	return result, "", nil
}

func ListInvoicePaymentsPage(site string, id string, offset string) ([]Transaction, string, error) {
	return ListInvoicePaymentsPageSortBy(site, id, offset, nil)
}