package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/bitsnap/chargebee-api-client/chargebee/enums"
	"github.com/bitsnap/chargebee-api-client/chargebee/models"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

type Item struct {
    Id string `json:"id" validate:"required"`
    Name string `json:"name" validate:"required"`
    ExternalName string `json:"external_name"`
    Description string `json:"description"`
    Status enums.StatusEnum `json:"status"`
    ResourceVersion int64 `json:"resource_version"`
    UpdatedAt uint64 `json:"updated_at"`
    ItemFamilyId string `json:"item_family_id"`
    Type enums.TypeEnum `json:"type" validate:"required"`
    IsShippable bool `json:"is_shippable"`
    IsGiftable bool `json:"is_giftable" validate:"required"`
    RedirectUrl string `json:"redirect_url"`
    EnabledForCheckout bool `json:"enabled_for_checkout" validate:"required"`
    EnabledInPortal bool `json:"enabled_in_portal" validate:"required"`
    IncludedInMrr bool `json:"included_in_mrr"`
    ItemApplicability enums.ItemApplicabilityEnum `json:"item_applicability" validate:"required"`
    GiftClaimRedirectUrl string `json:"gift_claim_redirect_url"`
    Unit string `json:"unit"`
    Metered bool `json:"metered" validate:"required"`
    UsageCalculation enums.UsageCalculationEnum `json:"usage_calculation"`
    ArchivedAt uint64 `json:"archived_at"`
    Channel enums.ChannelEnum `json:"channel"`
    Metadata map[string]interface{} `json:"metadata"`
    BusinessEntityId string `json:"business_entity_id"`
    ApplicableItems []models.ApplicableItem `json:"applicable_items"`
    BundleItems []models.BundleItem `json:"bundle_items"`
}

func ListItemsPageSortBy(site string, offset string, sortBy *SortBy) ([]Item, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/items")
	if err != nil {
		return nil, "", err
	}
	
    content, err := GetQuery(client, parsedUrl, offset, sortBy)
    if err != nil {
        return nil, "", err
    }
    		
	type ItemListItem struct {
		Item Item `json:"Item"`
	}

    type ItemPage struct {
        List       []ItemListItem `json:"list"`
        NextOffset string `json:"next_offset,omitempty"`
    }

	entries := ItemPage{
		List:       make([]ItemListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}
	
	if len(entries.List) == 0 {
        return []Item{}, "", nil
    }
	
	result := make([]Item, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Item)
	}

    if len(entries.NextOffset) > 0 {
        return ResultWithOffset(result, offset, entries.NextOffset)
    }
	
	return result, "", nil
}

func ListItemsPage(site string, offset string) ([]Item, string, error) {
	return ListItemsPageSortBy(site, offset, nil)
}

func RetrieveItemById(site string, id string) (*Item, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/items/" + id)
	if err != nil {
		return nil, err
	}
	
    content, err := GetQuery(client, parsedUrl, "", nil)
    if err != nil {
        return nil, err
    }
    	
	type ItemItem struct {
		Item Item `json:"Item"`
	}

    var item ItemItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.Item, nil
}