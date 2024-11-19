package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	"github.com/bitsnap/chargebee-api-client/chargebee/enums"
	"github.com/bitsnap/chargebee-api-client/chargebee/models"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type AdvancedInvoiceSchedule struct {
	Id                    string                       `json:"id" validate:"required"`
	ScheduleType          enums.ScheduleTypeEnum       `json:"schedule_type"`
	FixedIntervalSchedule models.FixedIntervalSchedule `json:"fixed_interval_schedule"`
	SpecificDatesSchedule models.SpecificDatesSchedule `json:"specific_dates_schedule"`
}

func ListAdvancedInvoiceSchedulesPageSortBy(site string, id string, offset string, sortBy *SortBy) ([]AdvancedInvoiceSchedule, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/subscriptions/" + id + "/retrieve_advance_invoice_schedule")
	if err != nil {
		return nil, "", err
	}

	content, err := getQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type AdvancedInvoiceScheduleListItem struct {
		AdvancedInvoiceSchedule AdvancedInvoiceSchedule `json:"AdvancedInvoiceSchedule"`
	}

	type AdvancedInvoiceSchedulePage struct {
		List       []AdvancedInvoiceScheduleListItem `json:"list"`
		NextOffset string                            `json:"next_offset,omitempty"`
	}

	entries := AdvancedInvoiceSchedulePage{
		List: make([]AdvancedInvoiceScheduleListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []AdvancedInvoiceSchedule{}, "", nil
	}

	result := make([]AdvancedInvoiceSchedule, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.AdvancedInvoiceSchedule)
	}

	if len(entries.NextOffset) > 0 {
		return resultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListAdvancedInvoiceSchedulesPage(site string, id string, offset string) ([]AdvancedInvoiceSchedule, string, error) {
	return ListAdvancedInvoiceSchedulesPageSortBy(site, id, offset, nil)
}
