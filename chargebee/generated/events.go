package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/bitsnap/chargebee-api-client/chargebee/enums"
	"github.com/bitsnap/chargebee-api-client/chargebee/models"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Event struct {
	Id         string                 `json:"id" validate:"required"`
	OccurredAt uint64                 `json:"occurred_at" validate:"required"`
	Source     enums.SourceEnum       `json:"source" validate:"required"`
	User       string                 `json:"user"`
	EventType  enums.EventTypeEnum    `json:"event_type"`
	ApiVersion enums.ApiVersionEnum   `json:"api_version"`
	Content    map[string]interface{} `json:"content" validate:"required"`
	OriginUser string                 `json:"origin_user"`
	Webhooks   []models.Webhook       `json:"webhooks"`
}

func ListEventsPageSortBy(site string, offset string, sortBy *SortBy) ([]Event, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/events")
	if err != nil {
		return nil, "", err
	}

	content, err := GetQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type EventListItem struct {
		Event Event `json:"Event"`
	}

	type EventPage struct {
		List       []EventListItem `json:"list"`
		NextOffset string          `json:"next_offset,omitempty"`
	}

	entries := EventPage{
		List: make([]EventListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []Event{}, "", nil
	}

	result := make([]Event, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Event)
	}

	if len(entries.NextOffset) > 0 {
		return ResultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListEventsPage(site string, offset string) ([]Event, string, error) {
	return ListEventsPageSortBy(site, offset, nil)
}

func RetrieveEventById(site string, id string) (*Event, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/events/" + id)
	if err != nil {
		return nil, err
	}

	content, err := GetQuery(client, parsedUrl, "", nil)
	if err != nil {
		return nil, err
	}

	type EventItem struct {
		Event Event `json:"Event"`
	}

	var item EventItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.Event, nil
}
