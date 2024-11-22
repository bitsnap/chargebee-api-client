package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func ListCustomerContactsPageSortBy(site string, id string, offset string, sortBy *SortBy) ([]Contact, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/customers/" + id + "/contacts")
	if err != nil {
		return nil, "", err
	}

	content, err := GetQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type ContactListItem struct {
		Contact Contact `json:"Contact"`
	}

	type ContactPage struct {
		List       []ContactListItem `json:"list"`
		NextOffset string            `json:"next_offset,omitempty"`
	}

	entries := ContactPage{
		List: make([]ContactListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []Contact{}, "", nil
	}

	result := make([]Contact, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Contact)
	}

	if len(entries.NextOffset) > 0 {
		return ResultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListCustomerContactsPage(site string, id string, offset string) ([]Contact, string, error) {
	return ListCustomerContactsPageSortBy(site, id, offset, nil)
}
