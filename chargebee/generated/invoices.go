package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/bitsnap/chargebee-api-client/chargebee/enums"
	"github.com/bitsnap/chargebee-api-client/chargebee/models"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/shopspring/decimal"
)

type Invoice struct {
	Id                        string                       `json:"id" validate:"required"`
	PoNumber                  string                       `json:"po_number"`
	CustomerId                string                       `json:"customer_id" validate:"required"`
	SubscriptionId            string                       `json:"subscription_id"`
	Recurring                 bool                         `json:"recurring" validate:"required"`
	Status                    enums.StatusEnum             `json:"status" validate:"required"`
	VatNumber                 string                       `json:"vat_number"`
	PriceType                 enums.PriceTypeEnum          `json:"price_type" validate:"required"`
	Date                      uint64                       `json:"date"`
	DueDate                   uint64                       `json:"due_date"`
	NetTermDays               int                          `json:"net_term_days"`
	ExchangeRate              *decimal.Decimal             `json:"exchange_rate"`
	CurrencyCode              string                       `json:"currency_code" validate:"required"`
	Total                     uint64                       `json:"total"`
	AmountPaid                uint64                       `json:"amount_paid"`
	AmountAdjusted            uint64                       `json:"amount_adjusted"`
	WriteOffAmount            uint64                       `json:"write_off_amount"`
	CreditsApplied            uint64                       `json:"credits_applied"`
	AmountDue                 uint64                       `json:"amount_due"`
	PaidAt                    uint64                       `json:"paid_at"`
	DunningStatus             enums.DunningStatusEnum      `json:"dunning_status"`
	NextRetryAt               uint64                       `json:"next_retry_at"`
	VoidedAt                  uint64                       `json:"voided_at"`
	ResourceVersion           int64                        `json:"resource_version"`
	UpdatedAt                 uint64                       `json:"updated_at"`
	SubTotal                  uint64                       `json:"sub_total" validate:"required"`
	SubTotalInLocalCurrency   uint64                       `json:"sub_total_in_local_currency"`
	TotalInLocalCurrency      uint64                       `json:"total_in_local_currency"`
	LocalCurrencyCode         string                       `json:"local_currency_code"`
	Tax                       uint64                       `json:"tax" validate:"required"`
	LocalCurrencyExchangeRate *decimal.Decimal             `json:"local_currency_exchange_rate"`
	FirstInvoice              bool                         `json:"first_invoice"`
	NewSalesAmount            uint64                       `json:"new_sales_amount"`
	HasAdvanceCharges         bool                         `json:"has_advance_charges"`
	TermFinalized             bool                         `json:"term_finalized" validate:"required"`
	IsGifted                  bool                         `json:"is_gifted" validate:"required"`
	GeneratedAt               uint64                       `json:"generated_at"`
	ExpectedPaymentDate       uint64                       `json:"expected_payment_date"`
	AmountToCollect           uint64                       `json:"amount_to_collect"`
	RoundOffAmount            uint64                       `json:"round_off_amount"`
	PaymentOwner              string                       `json:"payment_owner"`
	VoidReasonCode            string                       `json:"void_reason_code"`
	Deleted                   bool                         `json:"deleted" validate:"required"`
	TaxCategory               string                       `json:"tax_category"`
	VatNumberPrefix           string                       `json:"vat_number_prefix"`
	Channel                   enums.ChannelEnum            `json:"channel"`
	BusinessEntityId          string                       `json:"business_entity_id"`
	LineItems                 []models.LineItem            `json:"line_items"`
	Discounts                 []models.Discount            `json:"discounts"`
	LineItemDiscounts         []models.LineItemDiscount    `json:"line_item_discounts"`
	Taxes                     []models.Tax                 `json:"taxes"`
	LineItemTaxes             []models.LineItemTax         `json:"line_item_taxes"`
	LineItemTiers             []models.LineItemTier        `json:"line_item_tiers"`
	LinkedPayments            []models.InvoiceTransaction  `json:"linked_payments"`
	DunningAttempts           []models.DunningAttempt      `json:"dunning_attempts"`
	AppliedCredits            []models.AppliedCredit       `json:"applied_credits"`
	AdjustmentCreditNotes     []models.CreatedCreditNote   `json:"adjustment_credit_notes"`
	IssuedCreditNotes         []models.CreatedCreditNote   `json:"issued_credit_notes"`
	LinkedOrders              []models.LinkedOrder         `json:"linked_orders"`
	Notes                     []models.Note                `json:"notes"`
	ShippingAddress           models.ShippingAddress       `json:"shipping_address"`
	StatementDescriptor       models.StatementDescriptor   `json:"statement_descriptor"`
	BillingAddress            models.BillingAddress        `json:"billing_address"`
	Einvoice                  models.Einvoice              `json:"einvoice"`
	LinkedTaxesWithheld       []models.LinkedTaxWithheld   `json:"linked_taxes_withheld"`
	SiteDetailsAtCreation     models.SiteDetailsAtCreation `json:"site_details_at_creation"`
	TaxOrigin                 models.TaxOrigin             `json:"tax_origin"`
}

func ListInvoicesPageSortBy(site string, offset string, sortBy *SortBy) ([]Invoice, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/invoices")
	if err != nil {
		return nil, "", err
	}

	content, err := GetQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type InvoiceListItem struct {
		Invoice Invoice `json:"Invoice"`
	}

	type InvoicePage struct {
		List       []InvoiceListItem `json:"list"`
		NextOffset string            `json:"next_offset,omitempty"`
	}

	entries := InvoicePage{
		List: make([]InvoiceListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []Invoice{}, "", nil
	}

	result := make([]Invoice, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Invoice)
	}

	if len(entries.NextOffset) > 0 {
		return ResultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListInvoicesPage(site string, offset string) ([]Invoice, string, error) {
	return ListInvoicesPageSortBy(site, offset, nil)
}

func RetrieveInvoiceById(site string, id string) (*Invoice, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/invoices/" + id)
	if err != nil {
		return nil, err
	}

	content, err := GetQuery(client, parsedUrl, "", nil)
	if err != nil {
		return nil, err
	}

	type InvoiceItem struct {
		Invoice Invoice `json:"Invoice"`
	}

	var item InvoiceItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.Invoice, nil
}
