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

type Gift struct {
	Id              string                `json:"id" validate:"required"`
	Status          enums.StatusEnum      `json:"status" validate:"required"`
	ScheduledAt     uint64                `json:"scheduled_at"`
	AutoClaim       bool                  `json:"auto_claim" validate:"required"`
	NoExpiry        bool                  `json:"no_expiry" validate:"required"`
	ClaimExpiryDate uint64                `json:"claim_expiry_date"`
	ResourceVersion int64                 `json:"resource_version"`
	UpdatedAt       uint64                `json:"updated_at"`
	Gifter          models.Gifter         `json:"gifter" validate:"required"`
	GiftReceiver    models.GiftReceiver   `json:"gift_receiver" validate:"required"`
	GiftTimelines   []models.GiftTimeline `json:"gift_timelines"`
}

func ListGiftsPageSortBy(site string, offset string, sortBy *SortBy) ([]Gift, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/gifts")
	if err != nil {
		return nil, "", err
	}

	content, err := GetQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type GiftListItem struct {
		Gift Gift `json:"Gift"`
	}

	type GiftPage struct {
		List       []GiftListItem `json:"list"`
		NextOffset string         `json:"next_offset,omitempty"`
	}

	entries := GiftPage{
		List: make([]GiftListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []Gift{}, "", nil
	}

	result := make([]Gift, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Gift)
	}

	if len(entries.NextOffset) > 0 {
		return ResultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListGiftsPage(site string, offset string) ([]Gift, string, error) {
	return ListGiftsPageSortBy(site, offset, nil)
}

func RetrieveGiftById(site string, id string) (*Gift, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/gifts/" + id)
	if err != nil {
		return nil, err
	}

	content, err := GetQuery(client, parsedUrl, "", nil)
	if err != nil {
		return nil, err
	}

	type GiftItem struct {
		Gift Gift `json:"Gift"`
	}

	var item GiftItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.Gift, nil
}
