package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/bitsnap/chargebee-api-client/chargebee/generated/models"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type SubscriptionEntitlement struct {
	SubscriptionId string           `json:"subscription_id" validate:"required"`
	FeatureId      string           `json:"feature_id"`
	FeatureName    string           `json:"feature_name"`
	FeatureUnit    string           `json:"feature_unit"`
	FeatureType    string           `json:"feature_type"`
	Value          string           `json:"value"`
	Name           string           `json:"name"`
	IsOverridden   bool             `json:"is_overridden" validate:"required"`
	IsEnabled      bool             `json:"is_enabled" validate:"required"`
	ExpiresAt      uint64           `json:"expires_at"`
	Components     models.Component `json:"components"`
}

func ListSubscriptionEntitlementsPageSortBy(site string, id string, offset string, sortBy *SortBy) ([]SubscriptionEntitlement, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/subscriptions/" + id + "/subscription_entitlements")
	if err != nil {
		return nil, "", err
	}

	content, err := GetQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type SubscriptionEntitlementListItem struct {
		SubscriptionEntitlement SubscriptionEntitlement `json:"SubscriptionEntitlement"`
	}

	type SubscriptionEntitlementPage struct {
		List       []SubscriptionEntitlementListItem `json:"list"`
		NextOffset string                            `json:"next_offset,omitempty"`
	}

	entries := SubscriptionEntitlementPage{
		List: make([]SubscriptionEntitlementListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []SubscriptionEntitlement{}, "", nil
	}

	result := make([]SubscriptionEntitlement, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.SubscriptionEntitlement)
	}

	if len(entries.NextOffset) > 0 {
		return ResultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListSubscriptionEntitlementsPage(site string, id string, offset string) ([]SubscriptionEntitlement, string, error) {
	return ListSubscriptionEntitlementsPageSortBy(site, id, offset, nil)
}
