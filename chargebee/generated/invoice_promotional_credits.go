package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/bitsnap/chargebee-api-client/chargebee/generated/enums"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

type PromotionalCredit struct {
    Id string `json:"id" validate:"required"`
    CustomerId string `json:"customer_id" validate:"required"`
    Type enums.TypeEnum `json:"type" validate:"required"`
    AmountInDecimal string `json:"amount_in_decimal"`
    Amount uint64 `json:"amount" validate:"required"`
    CurrencyCode string `json:"currency_code" validate:"required"`
    Description string `json:"description" validate:"required"`
    CreditType enums.CreditTypeEnum `json:"credit_type" validate:"required"`
    Reference string `json:"reference"`
    ClosingBalance uint64 `json:"closing_balance" validate:"required"`
    DoneBy string `json:"done_by"`
    CreatedAt uint64 `json:"created_at" validate:"required"`
}

func ListPromotionalCreditsPageSortBy(site string, offset string, sortBy *SortBy) ([]PromotionalCredit, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/promotional_credits")
	if err != nil {
		return nil, "", err
	}
	
    content, err := GetQuery(client, parsedUrl, offset, sortBy)
    if err != nil {
        return nil, "", err
    }
    		
	type PromotionalCreditListItem struct {
		PromotionalCredit PromotionalCredit `json:"PromotionalCredit"`
	}

    type PromotionalCreditPage struct {
        List       []PromotionalCreditListItem `json:"list"`
        NextOffset string `json:"next_offset,omitempty"`
    }

	entries := PromotionalCreditPage{
		List:       make([]PromotionalCreditListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}
	
	if len(entries.List) == 0 {
        return []PromotionalCredit{}, "", nil
    }
	
	result := make([]PromotionalCredit, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.PromotionalCredit)
	}

    if len(entries.NextOffset) > 0 {
        return ResultWithOffset(result, offset, entries.NextOffset)
    }
	
	return result, "", nil
}

func ListPromotionalCreditsPage(site string, offset string) ([]PromotionalCredit, string, error) {
	return ListPromotionalCreditsPageSortBy(site, offset, nil)
}

func RetrievePromotionalCreditById(site string, id string) (*PromotionalCredit, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/promotional_credits/" + id)
	if err != nil {
		return nil, err
	}
	
    content, err := GetQuery(client, parsedUrl, "", nil)
    if err != nil {
        return nil, err
    }
    	
	type PromotionalCreditItem struct {
		PromotionalCredit PromotionalCredit `json:"PromotionalCredit"`
	}

    var item PromotionalCreditItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.PromotionalCredit, nil
}