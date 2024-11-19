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

type Coupon struct {
	Id                     string                          `json:"id" validate:"required"`
	Name                   string                          `json:"name" validate:"required"`
	InvoiceName            string                          `json:"invoice_name"`
	DiscountType           enums.DiscountTypeEnum          `json:"discount_type" validate:"required"`
	DiscountPercentage     float64                         `json:"discount_percentage"`
	DiscountAmount         uint64                          `json:"discount_amount"`
	CurrencyCode           string                          `json:"currency_code"`
	DurationType           enums.DurationTypeEnum          `json:"duration_type" validate:"required"`
	ValidFrom              uint64                          `json:"valid_from"`
	ValidTill              uint64                          `json:"valid_till"`
	MaxRedemptions         int                             `json:"max_redemptions"`
	Status                 enums.StatusEnum                `json:"status"`
	ApplyOn                enums.ApplyOnEnum               `json:"apply_on" validate:"required"`
	CreatedAt              uint64                          `json:"created_at" validate:"required"`
	ArchivedAt             uint64                          `json:"archived_at"`
	ResourceVersion        int64                           `json:"resource_version"`
	UpdatedAt              uint64                          `json:"updated_at"`
	Period                 int                             `json:"period"`
	PeriodUnit             enums.PeriodUnitEnum            `json:"period_unit"`
	Redemptions            int                             `json:"redemptions"`
	InvoiceNotes           string                          `json:"invoice_notes"`
	ItemConstraints        []models.ItemConstraint         `json:"item_constraints"`
	ItemConstraintCriteria []models.ItemConstraintCriteria `json:"item_constraint_criteria"`
	CouponConstraints      []models.CouponConstraint       `json:"coupon_constraints"`
}

func ListCouponsPageSortBy(site string, offset string, sortBy *SortBy) ([]Coupon, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/coupons")
	if err != nil {
		return nil, "", err
	}

	content, err := GetQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type CouponListItem struct {
		Coupon Coupon `json:"Coupon"`
	}

	type CouponPage struct {
		List       []CouponListItem `json:"list"`
		NextOffset string           `json:"next_offset,omitempty"`
	}

	entries := CouponPage{
		List: make([]CouponListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []Coupon{}, "", nil
	}

	result := make([]Coupon, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Coupon)
	}

	if len(entries.NextOffset) > 0 {
		return ResultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListCouponsPage(site string, offset string) ([]Coupon, string, error) {
	return ListCouponsPageSortBy(site, offset, nil)
}

func RetrieveCouponById(site string, id string) (*Coupon, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/coupons/" + id)
	if err != nil {
		return nil, err
	}

	content, err := GetQuery(client, parsedUrl, "", nil)
	if err != nil {
		return nil, err
	}

	type CouponItem struct {
		Coupon Coupon `json:"Coupon"`
	}

	var item CouponItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.Coupon, nil
}
