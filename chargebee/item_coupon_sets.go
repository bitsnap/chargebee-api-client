package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type CouponSet struct {
	Id            string `json:"id" validate:"required"`
	CouponId      string `json:"coupon_id" validate:"required"`
	Name          string `json:"name" validate:"required"`
	TotalCount    int    `json:"total_count"`
	RedeemedCount int    `json:"redeemed_count"`
	ArchivedCount int    `json:"archived_count"`
}

func ListCouponSetsPageSortBy(site string, offset string, sortBy *SortBy) ([]CouponSet, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/coupon_sets")
	if err != nil {
		return nil, "", err
	}

	content, err := getQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type CouponSetListItem struct {
		CouponSet CouponSet `json:"CouponSet"`
	}

	type CouponSetPage struct {
		List       []CouponSetListItem `json:"list"`
		NextOffset string              `json:"next_offset,omitempty"`
	}

	entries := CouponSetPage{
		List: make([]CouponSetListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []CouponSet{}, "", nil
	}

	result := make([]CouponSet, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.CouponSet)
	}

	if len(entries.NextOffset) > 0 {
		return resultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListCouponSetsPage(site string, offset string) ([]CouponSet, string, error) {
	return ListCouponSetsPageSortBy(site, offset, nil)
}

func RetrieveCouponSetById(site string, id string) (*CouponSet, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/coupon_sets/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQuery(client, parsedUrl, "", nil)
	if err != nil {
		return nil, err
	}

	type CouponSetItem struct {
		CouponSet CouponSet `json:"CouponSet"`
	}

	var item CouponSetItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.CouponSet, nil
}
