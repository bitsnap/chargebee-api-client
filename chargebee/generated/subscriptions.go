package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/bitsnap/chargebee-api-client/chargebee/generated/enums"
	"github.com/bitsnap/chargebee-api-client/chargebee/generated/models"
	"github.com/shopspring/decimal"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

type Subscription struct {
    Id string `json:"id" validate:"required"`
    CurrencyCode string `json:"currency_code" validate:"required"`
    StartDate uint64 `json:"start_date"`
    TrialEnd uint64 `json:"trial_end"`
    RemainingBillingCycles int `json:"remaining_billing_cycles"`
    PoNumber string `json:"po_number"`
    PlanQuantityInDecimal string `json:"plan_quantity_in_decimal"`
    PlanUnitPriceInDecimal string `json:"plan_unit_price_in_decimal"`
    CustomerId string `json:"customer_id" validate:"required"`
    Status enums.StatusEnum `json:"status" validate:"required"`
    TrialStart uint64 `json:"trial_start"`
    TrialEndAction enums.TrialEndActionEnum `json:"trial_end_action"`
    CurrentTermStart uint64 `json:"current_term_start"`
    CurrentTermEnd uint64 `json:"current_term_end"`
    NextBillingAt uint64 `json:"next_billing_at"`
    CreatedAt uint64 `json:"created_at"`
    StartedAt uint64 `json:"started_at"`
    ActivatedAt uint64 `json:"activated_at"`
    ContractTermBillingCycleOnRenewal int `json:"contract_term_billing_cycle_on_renewal"`
    OverrideRelationship bool `json:"override_relationship"`
    PauseDate uint64 `json:"pause_date"`
    ResumeDate uint64 `json:"resume_date"`
    CancelledAt uint64 `json:"cancelled_at"`
    CancelReason enums.CancelReasonEnum `json:"cancel_reason"`
    CreatedFromIp string `json:"created_from_ip"`
    ResourceVersion int64 `json:"resource_version"`
    UpdatedAt uint64 `json:"updated_at"`
    HasScheduledAdvanceInvoices bool `json:"has_scheduled_advance_invoices" validate:"required"`
    HasScheduledChanges bool `json:"has_scheduled_changes" validate:"required"`
    PaymentSourceId string `json:"payment_source_id"`
    PlanFreeQuantityInDecimal string `json:"plan_free_quantity_in_decimal"`
    PlanAmountInDecimal string `json:"plan_amount_in_decimal"`
    CancelScheduleCreatedAt uint64 `json:"cancel_schedule_created_at"`
    Channel enums.ChannelEnum `json:"channel"`
    NetTermDays int `json:"net_term_days"`
    ActiveId string `json:"active_id"`
    DueInvoicesCount int `json:"due_invoices_count"`
    DueSince uint64 `json:"due_since"`
    TotalDues uint64 `json:"total_dues"`
    Mrr uint64 `json:"mrr"`
    ExchangeRate *decimal.Decimal `json:"exchange_rate"`
    BaseCurrencyCode string `json:"base_currency_code"`
    InvoiceNotes string `json:"invoice_notes"`
    Metadata map[string]interface{} `json:"metadata"`
    Deleted bool `json:"deleted" validate:"required"`
    ChangesScheduledAt uint64 `json:"changes_scheduled_at"`
    CancelReasonCode string `json:"cancel_reason_code"`
    FreePeriod int `json:"free_period"`
    FreePeriodUnit enums.FreePeriodUnitEnum `json:"free_period_unit"`
    CreatePendingInvoices bool `json:"create_pending_invoices"`
    AutoCloseInvoices bool `json:"auto_close_invoices"`
    BusinessEntityId string `json:"business_entity_id"`
    SubscriptionItems []models.SubscriptionItem `json:"subscription_items"`
    ItemTiers []models.ItemTier `json:"item_tiers"`
    ChargedItems []models.ChargedItem `json:"charged_items"`
    Coupons []models.Coupon `json:"coupons"`
    ShippingAddress models.ShippingAddress `json:"shipping_address"`
    ReferralInfo models.ReferralInfo `json:"referral_info"`
    ContractTerm models.ContractTerm `json:"contract_term"`
    Discounts []models.Discount `json:"discounts"`
}

func ListSubscriptionsPageSortBy(site string, offset string, sortBy *SortBy) ([]Subscription, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/subscriptions")
	if err != nil {
		return nil, "", err
	}
	
    content, err := GetQuery(client, parsedUrl, offset, sortBy)
    if err != nil {
        return nil, "", err
    }
    		
	type SubscriptionListItem struct {
		Subscription Subscription `json:"Subscription"`
	}

    type SubscriptionPage struct {
        List       []SubscriptionListItem `json:"list"`
        NextOffset string `json:"next_offset,omitempty"`
    }

	entries := SubscriptionPage{
		List:       make([]SubscriptionListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}
	
	if len(entries.List) == 0 {
        return []Subscription{}, "", nil
    }
	
	result := make([]Subscription, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Subscription)
	}

    if len(entries.NextOffset) > 0 {
        return ResultWithOffset(result, offset, entries.NextOffset)
    }
	
	return result, "", nil
}

func ListSubscriptionsPage(site string, offset string) ([]Subscription, string, error) {
	return ListSubscriptionsPageSortBy(site, offset, nil)
}

func ListSubscriptionWithScheduledChangesPageSortBy(site string, id string, offset string, sortBy *SortBy) ([]Subscription, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/subscriptions/" + id + "/retrieve_with_scheduled_changes")
	if err != nil {
		return nil, "", err
	}
		
    content, err := GetQuery(client, parsedUrl, offset, sortBy)
    if err != nil {
        return nil, "", err
    }
    	
	type SubscriptionListItem struct {
		Subscription Subscription `json:"Subscription"`
	}

    type SubscriptionPage struct {
        List       []SubscriptionListItem `json:"list"`
        NextOffset string `json:"next_offset,omitempty"`
    }

	entries := SubscriptionPage{
		List:       make([]SubscriptionListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}
	
	if len(entries.List) == 0 {
        return []Subscription{}, "", nil
    }
	
	result := make([]Subscription, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Subscription)
	}

    if len(entries.NextOffset) > 0 {
        return ResultWithOffset(result, offset, entries.NextOffset)
    }
	
	return result, "", nil
}

func ListSubscriptionWithScheduledChangesPage(site string, id string, offset string) ([]Subscription, string, error) {
	return ListSubscriptionWithScheduledChangesPageSortBy(site, id, offset, nil)
}