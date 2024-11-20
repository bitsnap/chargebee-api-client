package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/bitsnap/chargebee-api-client/chargebee/enums"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

type Currency struct {
    Id string `json:"id"`
    Enabled bool `json:"enabled" validate:"required"`
    ForexType enums.ForexTypeEnum `json:"forex_type"`
    CurrencyCode string `json:"currency_code"`
    IsBaseCurrency bool `json:"is_base_currency"`
    ManualExchangeRate string `json:"manual_exchange_rate"`
}

func ListCurrenciesPageSortBy(site string, offset string, sortBy *SortBy) ([]Currency, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/currencies")
	if err != nil {
		return nil, "", err
	}
	
    content, err := GetQuery(client, parsedUrl, offset, sortBy)
    if err != nil {
        return nil, "", err
    }
    		
	type CurrencyListItem struct {
		Currency Currency `json:"Currency"`
	}

    type CurrencyPage struct {
        List       []CurrencyListItem `json:"list"`
        NextOffset string `json:"next_offset,omitempty"`
    }

	entries := CurrencyPage{
		List:       make([]CurrencyListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}
	
	if len(entries.List) == 0 {
        return []Currency{}, "", nil
    }
	
	result := make([]Currency, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Currency)
	}

    if len(entries.NextOffset) > 0 {
        return ResultWithOffset(result, offset, entries.NextOffset)
    }
	
	return result, "", nil
}

func ListCurrenciesPage(site string, offset string) ([]Currency, string, error) {
	return ListCurrenciesPageSortBy(site, offset, nil)
}

func RetrieveCurrencyById(site string, id string) (*Currency, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/currencies/" + id)
	if err != nil {
		return nil, err
	}
	
    content, err := GetQuery(client, parsedUrl, "", nil)
    if err != nil {
        return nil, err
    }
    	
	type CurrencyItem struct {
		Currency Currency `json:"Currency"`
	}

    var item CurrencyItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.Currency, nil
}