package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/bitsnap/chargebee-api-client/chargebee/enums"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

type PaymentReferenceNumber struct {
    Id string `json:"id" validate:"required"`
    Type enums.TypeEnum `json:"type" validate:"required"`
    Number string `json:"number" validate:"required"`
    InvoiceId string `json:"invoice_id"`
}

func ListPaymentReferenceNumbersPageSortBy(site string, id string, offset string, sortBy *SortBy) ([]PaymentReferenceNumber, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/invoices/payment_reference_numbers")
	if err != nil {
		return nil, "", err
	}
	
	q := parsedUrl.Query()
	q.Add("id[is]", id)
	parsedUrl.RawQuery = q.Encode()
		
    content, err := GetQuery(client, parsedUrl, offset, sortBy)
    if err != nil {
        return nil, "", err
    }
    	
	type PaymentReferenceNumberListItem struct {
		PaymentReferenceNumber PaymentReferenceNumber `json:"PaymentReferenceNumber"`
	}

    type PaymentReferenceNumberPage struct {
        List       []PaymentReferenceNumberListItem `json:"list"`
        NextOffset string `json:"next_offset,omitempty"`
    }

	entries := PaymentReferenceNumberPage{
		List:       make([]PaymentReferenceNumberListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}
	
	if len(entries.List) == 0 {
        return []PaymentReferenceNumber{}, "", nil
    }
	
	result := make([]PaymentReferenceNumber, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.PaymentReferenceNumber)
	}

    if len(entries.NextOffset) > 0 {
        return ResultWithOffset(result, offset, entries.NextOffset)
    }
	
	return result, "", nil
}

func ListPaymentReferenceNumbersPage(site string, id string, offset string) ([]PaymentReferenceNumber, string, error) {
	return ListPaymentReferenceNumbersPageSortBy(site, id, offset, nil)
}