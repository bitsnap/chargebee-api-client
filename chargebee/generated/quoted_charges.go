package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"github.com/bitsnap/chargebee-api-client/chargebee/generated/models"
)

type QuotedCharge struct {
	Charges      []models.Charge    `json:"charges"`
	InvoiceItems []models.ItemPrice `json:"invoice_items"`
	ItemTiers    []models.ItemTier  `json:"item_tiers"`
	Coupons      []models.Coupon    `json:"coupons"`
	Discounts    []models.Discount  `json:"discounts"`
}
