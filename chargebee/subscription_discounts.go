package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	"github.com/bitsnap/chargebee-api-client/chargebee/enums"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Discount struct {
	Id            string                 `json:"id" validate:"required"`
	InvoiceName   string                 `json:"invoice_name"`
	Type          enums.TypeEnum         `json:"type" validate:"required"`
	Percentage    float64                `json:"percentage"`
	Amount        uint64                 `json:"amount"`
	CurrencyCode  string                 `json:"currency_code"`
	DurationType  enums.DurationTypeEnum `json:"duration_type" validate:"required"`
	Period        int                    `json:"period"`
	PeriodUnit    enums.PeriodUnitEnum   `json:"period_unit"`
	IncludedInMrr bool                   `json:"included_in_mrr" validate:"required"`
	ApplyOn       enums.ApplyOnEnum      `json:"apply_on" validate:"required"`
	ItemPriceId   string                 `json:"item_price_id"`
	CreatedAt     uint64                 `json:"created_at" validate:"required"`
	ApplyTill     uint64                 `json:"apply_till"`
	AppliedCount  int                    `json:"applied_count"`
	CouponId      string                 `json:"coupon_id" validate:"required"`
	Index         int                    `json:"index" validate:"required"`
}

func ListDiscountsPageSortBy(site string, id string, offset string, sortBy *SortBy) ([]Discount, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/subscriptions/" + id + "/discounts")
	if err != nil {
		return nil, "", err
	}

	content, err := getQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type DiscountListItem struct {
		Discount Discount `json:"Discount"`
	}

	type DiscountPage struct {
		List       []DiscountListItem `json:"list"`
		NextOffset string             `json:"next_offset,omitempty"`
	}

	entries := DiscountPage{
		List: make([]DiscountListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []Discount{}, "", nil
	}

	result := make([]Discount, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Discount)
	}

	if len(entries.NextOffset) > 0 {
		return resultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListDiscountsPage(site string, id string, offset string) ([]Discount, string, error) {
	return ListDiscountsPageSortBy(site, id, offset, nil)
}
