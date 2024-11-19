package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	"github.com/bitsnap/chargebee-api-client/chargebee/enums"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Entitlement struct {
	Id          string               `json:"id" validate:"required"`
	EntityId    string               `json:"entity_id"`
	EntityType  enums.EntityTypeEnum `json:"entity_type"`
	FeatureId   string               `json:"feature_id"`
	FeatureName string               `json:"feature_name"`
	Value       string               `json:"value"`
	Name        string               `json:"name"`
}

func ListEntitlementsPageSortBy(site string, offset string, sortBy *SortBy) ([]Entitlement, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/entitlements")
	if err != nil {
		return nil, "", err
	}

	content, err := getQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type EntitlementListItem struct {
		Entitlement Entitlement `json:"Entitlement"`
	}

	type EntitlementPage struct {
		List       []EntitlementListItem `json:"list"`
		NextOffset string                `json:"next_offset,omitempty"`
	}

	entries := EntitlementPage{
		List: make([]EntitlementListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []Entitlement{}, "", nil
	}

	result := make([]Entitlement, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Entitlement)
	}

	if len(entries.NextOffset) > 0 {
		return resultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListEntitlementsPage(site string, offset string) ([]Entitlement, string, error) {
	return ListEntitlementsPageSortBy(site, offset, nil)
}
