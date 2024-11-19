package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

func ListInvoicePaymentVouchersPageSortBy(site string, id string, offset string, sortBy *SortBy) ([]PaymentVoucher, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/invoices/" + id + "/payment_vouchers")
	if err != nil {
		return nil, "", err
	}
		
    content, err := GetQuery(client, parsedUrl, offset, sortBy)
    if err != nil {
        return nil, "", err
    }
    	
	type PaymentVoucherListItem struct {
		PaymentVoucher PaymentVoucher `json:"PaymentVoucher"`
	}

    type PaymentVoucherPage struct {
        List       []PaymentVoucherListItem `json:"list"`
        NextOffset string `json:"next_offset,omitempty"`
    }

	entries := PaymentVoucherPage{
		List:       make([]PaymentVoucherListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}
	
	if len(entries.List) == 0 {
        return []PaymentVoucher{}, "", nil
    }
	
	result := make([]PaymentVoucher, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.PaymentVoucher)
	}

    if len(entries.NextOffset) > 0 {
        return ResultWithOffset(result, offset, entries.NextOffset)
    }
	
	return result, "", nil
}

func ListInvoicePaymentVouchersPage(site string, id string, offset string) ([]PaymentVoucher, string, error) {
	return ListInvoicePaymentVouchersPageSortBy(site, id, offset, nil)
}