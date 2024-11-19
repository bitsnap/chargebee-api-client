package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func ListCustomerHierarchiesPageSortBy(site string, id string, offset string, sortBy *SortBy) ([]Hierarchy, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/customers/" + id + "/hierarchy")
	if err != nil {
		return nil, "", err
	}

	content, err := getQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type HierarchyListItem struct {
		Hierarchy Hierarchy `json:"Hierarchy"`
	}

	type HierarchyPage struct {
		List       []HierarchyListItem `json:"list"`
		NextOffset string              `json:"next_offset,omitempty"`
	}

	entries := HierarchyPage{
		List: make([]HierarchyListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []Hierarchy{}, "", nil
	}

	result := make([]Hierarchy, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Hierarchy)
	}

	if len(entries.NextOffset) > 0 {
		return resultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListCustomerHierarchiesPage(site string, id string, offset string) ([]Hierarchy, string, error) {
	return ListCustomerHierarchiesPageSortBy(site, id, offset, nil)
}

type Hierarchy struct {
	CustomerId     string   `json:"customer_id" validate:"required"`
	ParentId       string   `json:"parent_id"`
	PaymentOwnerId string   `json:"payment_owner_id" validate:"required"`
	InvoiceOwnerId string   `json:"invoice_owner_id" validate:"required"`
	ChildrenIds    []string `json:"children_ids"`
}
