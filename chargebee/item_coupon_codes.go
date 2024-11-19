package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	"github.com/bitsnap/chargebee-api-client/chargebee/enums"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type CouponCode struct {
	Code          string           `json:"code" validate:"required"`
	Status        enums.StatusEnum `json:"status" validate:"required"`
	CouponId      string           `json:"coupon_id" validate:"required"`
	CouponSetId   string           `json:"coupon_set_id" validate:"required"`
	CouponSetName string           `json:"coupon_set_name" validate:"required"`
}

func ListCouponCodesPageSortBy(site string, offset string, sortBy *SortBy) ([]CouponCode, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/coupon_codes")
	if err != nil {
		return nil, "", err
	}

	content, err := getQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type CouponCodeListItem struct {
		CouponCode CouponCode `json:"CouponCode"`
	}

	type CouponCodePage struct {
		List       []CouponCodeListItem `json:"list"`
		NextOffset string               `json:"next_offset,omitempty"`
	}

	entries := CouponCodePage{
		List: make([]CouponCodeListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []CouponCode{}, "", nil
	}

	result := make([]CouponCode, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.CouponCode)
	}

	if len(entries.NextOffset) > 0 {
		return resultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListCouponCodesPage(site string, offset string) ([]CouponCode, string, error) {
	return ListCouponCodesPageSortBy(site, offset, nil)
}

func RetrieveCouponCodeById(site string, id string) (*CouponCode, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/coupon_codes/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQuery(client, parsedUrl, "", nil)
	if err != nil {
		return nil, err
	}

	type CouponCodeItem struct {
		CouponCode CouponCode `json:"CouponCode"`
	}

	var item CouponCodeItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.CouponCode, nil
}
