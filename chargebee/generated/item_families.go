package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/bitsnap/chargebee-api-client/chargebee/enums"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

type ItemFamily struct {
    Id string `json:"id" validate:"required"`
    Name string `json:"name" validate:"required"`
    Description string `json:"description"`
    Status enums.StatusEnum `json:"status"`
    ResourceVersion int64 `json:"resource_version"`
    UpdatedAt uint64 `json:"updated_at"`
    Channel enums.ChannelEnum `json:"channel"`
    BusinessEntityId string `json:"business_entity_id"`
}

func ListItemFamiliesPageSortBy(site string, offset string, sortBy *SortBy) ([]ItemFamily, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/item_families")
	if err != nil {
		return nil, "", err
	}
	
    content, err := GetQuery(client, parsedUrl, offset, sortBy)
    if err != nil {
        return nil, "", err
    }
    		
	type ItemFamilyListItem struct {
		ItemFamily ItemFamily `json:"ItemFamily"`
	}

    type ItemFamilyPage struct {
        List       []ItemFamilyListItem `json:"list"`
        NextOffset string `json:"next_offset,omitempty"`
    }

	entries := ItemFamilyPage{
		List:       make([]ItemFamilyListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}
	
	if len(entries.List) == 0 {
        return []ItemFamily{}, "", nil
    }
	
	result := make([]ItemFamily, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.ItemFamily)
	}

    if len(entries.NextOffset) > 0 {
        return ResultWithOffset(result, offset, entries.NextOffset)
    }
	
	return result, "", nil
}

func ListItemFamiliesPage(site string, offset string) ([]ItemFamily, string, error) {
	return ListItemFamiliesPageSortBy(site, offset, nil)
}

func RetrieveItemFamilyById(site string, id string) (*ItemFamily, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/item_families/" + id)
	if err != nil {
		return nil, err
	}
	
    content, err := GetQuery(client, parsedUrl, "", nil)
    if err != nil {
        return nil, err
    }
    	
	type ItemFamilyItem struct {
		ItemFamily ItemFamily `json:"ItemFamily"`
	}

    var item ItemFamilyItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.ItemFamily, nil
}