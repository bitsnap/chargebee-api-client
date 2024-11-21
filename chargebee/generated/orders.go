package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/bitsnap/chargebee-api-client/chargebee/generated/enums"
	"github.com/bitsnap/chargebee-api-client/chargebee/generated/models"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

type Order struct {
    Id string `json:"id" validate:"required"`
    DocumentNumber string `json:"document_number"`
    InvoiceId string `json:"invoice_id"`
    SubscriptionId string `json:"subscription_id"`
    CustomerId string `json:"customer_id"`
    Status enums.StatusEnum `json:"status"`
    CancellationReason enums.CancellationReasonEnum `json:"cancellation_reason"`
    PaymentStatus enums.PaymentStatusEnum `json:"payment_status"`
    OrderType enums.OrderTypeEnum `json:"order_type"`
    PriceType enums.PriceTypeEnum `json:"price_type" validate:"required"`
    ReferenceId string `json:"reference_id"`
    FulfillmentStatus string `json:"fulfillment_status"`
    OrderDate uint64 `json:"order_date"`
    ShippingDate uint64 `json:"shipping_date"`
    Note string `json:"note"`
    TrackingId string `json:"tracking_id"`
    TrackingUrl string `json:"tracking_url"`
    BatchId string `json:"batch_id"`
    CreatedBy string `json:"created_by"`
    ShipmentCarrier string `json:"shipment_carrier"`
    InvoiceRoundOffAmount uint64 `json:"invoice_round_off_amount"`
    Tax uint64 `json:"tax"`
    AmountPaid uint64 `json:"amount_paid"`
    AmountAdjusted uint64 `json:"amount_adjusted"`
    RefundableCreditsIssued uint64 `json:"refundable_credits_issued"`
    RefundableCredits uint64 `json:"refundable_credits"`
    RoundingAdjustement uint64 `json:"rounding_adjustement"`
    PaidOn uint64 `json:"paid_on"`
    ShippingCutOffDate uint64 `json:"shipping_cut_off_date"`
    CreatedAt uint64 `json:"created_at" validate:"required"`
    StatusUpdateAt uint64 `json:"status_update_at"`
    DeliveredAt uint64 `json:"delivered_at"`
    ShippedAt uint64 `json:"shipped_at"`
    ResourceVersion int64 `json:"resource_version"`
    UpdatedAt uint64 `json:"updated_at"`
    CancelledAt uint64 `json:"cancelled_at"`
    ResentStatus enums.ResentStatusEnum `json:"resent_status"`
    IsResent bool `json:"is_resent" validate:"required"`
    OriginalOrderId string `json:"original_order_id"`
    Discount uint64 `json:"discount"`
    SubTotal uint64 `json:"sub_total"`
    Total uint64 `json:"total"`
    Deleted bool `json:"deleted" validate:"required"`
    CurrencyCode string `json:"currency_code"`
    IsGifted bool `json:"is_gifted"`
    GiftNote string `json:"gift_note"`
    GiftId string `json:"gift_id"`
    ResendReason string `json:"resend_reason"`
    BusinessEntityId string `json:"business_entity_id"`
    OrderLineItems []models.OrderLineItem `json:"order_line_items"`
    ShippingAddress models.ShippingAddress `json:"shipping_address"`
    BillingAddress models.BillingAddress `json:"billing_address"`
    LineItemTaxes []models.LineItemTax `json:"line_item_taxes"`
    LineItemDiscounts []models.LineItemDiscount `json:"line_item_discounts"`
    LinkedCreditNotes []models.OrderLineItemLinkedCredit `json:"linked_credit_notes"`
    ResentOrders []models.ResentOrder `json:"resent_orders"`
}

func ListOrdersPageSortBy(site string, offset string, sortBy *SortBy) ([]Order, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/orders")
	if err != nil {
		return nil, "", err
	}
	
    content, err := GetQuery(client, parsedUrl, offset, sortBy)
    if err != nil {
        return nil, "", err
    }
    		
	type OrderListItem struct {
		Order Order `json:"Order"`
	}

    type OrderPage struct {
        List       []OrderListItem `json:"list"`
        NextOffset string `json:"next_offset,omitempty"`
    }

	entries := OrderPage{
		List:       make([]OrderListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}
	
	if len(entries.List) == 0 {
        return []Order{}, "", nil
    }
	
	result := make([]Order, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Order)
	}

    if len(entries.NextOffset) > 0 {
        return ResultWithOffset(result, offset, entries.NextOffset)
    }
	
	return result, "", nil
}

func ListOrdersPage(site string, offset string) ([]Order, string, error) {
	return ListOrdersPageSortBy(site, offset, nil)
}

func RetrieveOrderById(site string, id string) (*Order, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/orders/" + id)
	if err != nil {
		return nil, err
	}
	
    content, err := GetQuery(client, parsedUrl, "", nil)
    if err != nil {
        return nil, err
    }
    	
	type OrderItem struct {
		Order Order `json:"Order"`
	}

    var item OrderItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.Order, nil
}