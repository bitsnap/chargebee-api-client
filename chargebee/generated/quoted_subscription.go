package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"github.com/bitsnap/chargebee-api-client/chargebee/generated/enums"
	"github.com/bitsnap/chargebee-api-client/chargebee/generated/models"

)

type QuotedSubscription struct {
    Id string `json:"id" validate:"required"`
    StartDate uint64 `json:"start_date"`
    TrialEnd uint64 `json:"trial_end"`
    RemainingBillingCycles int `json:"remaining_billing_cycles"`
    PoNumber string `json:"po_number"`
    PlanQuantityInDecimal string `json:"plan_quantity_in_decimal"`
    PlanUnitPriceInDecimal string `json:"plan_unit_price_in_decimal"`
    ChangesScheduledAt uint64 `json:"changes_scheduled_at"`
    ChangeOption enums.ChangeOptionEnum `json:"change_option"`
    ContractTermBillingCycleOnRenewal int `json:"contract_term_billing_cycle_on_renewal"`
    Coupons []models.Coupon `json:"coupons"`
    Discounts []models.Discount `json:"discounts"`
    SubscriptionItems []models.SubscriptionItem `json:"subscription_items"`
    ItemTiers []models.ItemTier `json:"item_tiers"`
    QuotedContractTerm models.QuotedContractTerm `json:"quoted_contract_term"`
}