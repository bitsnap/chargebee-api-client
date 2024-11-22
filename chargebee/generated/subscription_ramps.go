package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/bitsnap/chargebee-api-client/chargebee/generated/enums"
	"github.com/bitsnap/chargebee-api-client/chargebee/generated/models"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Ramp struct {
	Id                     string                        `json:"id" validate:"required"`
	Description            string                        `json:"description"`
	SubscriptionId         string                        `json:"subscription_id" validate:"required"`
	EffectiveFrom          uint64                        `json:"effective_from" validate:"required"`
	Status                 enums.StatusEnum              `json:"status" validate:"required"`
	CreatedAt              uint64                        `json:"created_at" validate:"required"`
	ResourceVersion        int64                         `json:"resource_version"`
	UpdatedAt              uint64                        `json:"updated_at"`
	ItemsToRemove          []string                      `json:"items_to_remove"`
	CouponsToRemove        []string                      `json:"coupons_to_remove"`
	DiscountsToRemove      []string                      `json:"discounts_to_remove"`
	Deleted                bool                          `json:"deleted" validate:"required"`
	ItemsToAdd             []models.ItemsToAdd           `json:"items_to_add"`
	ItemsToUpdate          []models.ItemsToUpdate        `json:"items_to_update"`
	CouponsToAdd           []models.CouponsToAdd         `json:"coupons_to_add"`
	DiscountsToAdd         []models.DiscountsToAdd       `json:"discounts_to_add"`
	ItemTiers              []models.ItemTier             `json:"item_tiers"`
	StatusTransitionReason models.StatusTransitionReason `json:"status_transition_reason"`
}

func ListRampsPageSortBy(site string, offset string, sortBy *SortBy) ([]Ramp, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/ramps")
	if err != nil {
		return nil, "", err
	}

	content, err := GetQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type RampListItem struct {
		Ramp Ramp `json:"Ramp"`
	}

	type RampPage struct {
		List       []RampListItem `json:"list"`
		NextOffset string         `json:"next_offset,omitempty"`
	}

	entries := RampPage{
		List: make([]RampListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []Ramp{}, "", nil
	}

	result := make([]Ramp, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Ramp)
	}

	if len(entries.NextOffset) > 0 {
		return ResultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListRampsPage(site string, offset string) ([]Ramp, string, error) {
	return ListRampsPageSortBy(site, offset, nil)
}
