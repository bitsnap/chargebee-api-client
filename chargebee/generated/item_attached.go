package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/bitsnap/chargebee-api-client/chargebee/generated/enums"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

type ItemAttached struct {
    Id string `json:"id" validate:"required"`
    ParentItemId string `json:"parent_item_id" validate:"required"`
    ItemId string `json:"item_id" validate:"required"`
    Type enums.TypeEnum `json:"type" validate:"required"`
    Status enums.StatusEnum `json:"status"`
    Quantity int `json:"quantity"`
    QuantityInDecimal string `json:"quantity_in_decimal"`
    BillingCycles int `json:"billing_cycles"`
    ChargeOnEvent enums.ChargeOnEventEnum `json:"charge_on_event" validate:"required"`
    ChargeOnce bool `json:"charge_once" validate:"required"`
    CreatedAt uint64 `json:"created_at" validate:"required"`
    ResourceVersion int64 `json:"resource_version"`
    UpdatedAt uint64 `json:"updated_at"`
    Channel enums.ChannelEnum `json:"channel"`
    BusinessEntityId string `json:"business_entity_id"`
}

func ListItemsAttachedPageSortBy(site string, id string, offset string, sortBy *SortBy) ([]ItemAttached, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/items/" + id + "/attached_items")
	if err != nil {
		return nil, "", err
	}
		
    content, err := GetQuery(client, parsedUrl, offset, sortBy)
    if err != nil {
        return nil, "", err
    }
    	
	type ItemAttachedListItem struct {
		ItemAttached ItemAttached `json:"ItemAttached"`
	}

    type ItemAttachedPage struct {
        List       []ItemAttachedListItem `json:"list"`
        NextOffset string `json:"next_offset,omitempty"`
    }

	entries := ItemAttachedPage{
		List:       make([]ItemAttachedListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}
	
	if len(entries.List) == 0 {
        return []ItemAttached{}, "", nil
    }
	
	result := make([]ItemAttached, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.ItemAttached)
	}

    if len(entries.NextOffset) > 0 {
        return ResultWithOffset(result, offset, entries.NextOffset)
    }
	
	return result, "", nil
}

func ListItemsAttachedPage(site string, id string, offset string) ([]ItemAttached, string, error) {
	return ListItemsAttachedPageSortBy(site, id, offset, nil)
}

func RetrieveItemAttachedById(site string, id string) (*ItemAttached, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/attached_items/" + id)
	if err != nil {
		return nil, err
	}
	
    content, err := GetQuery(client, parsedUrl, "", nil)
    if err != nil {
        return nil, err
    }
    	
	type ItemAttachedItem struct {
		ItemAttached ItemAttached `json:"ItemAttached"`
	}

    var item ItemAttachedItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.ItemAttached, nil
}