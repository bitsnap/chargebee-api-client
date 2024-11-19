package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	"github.com/bitsnap/chargebee-api-client/chargebee/enums"
	"github.com/bitsnap/chargebee-api-client/chargebee/models"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type PaymentSource struct {
	Id               string                   `json:"id" validate:"required"`
	ResourceVersion  int64                    `json:"resource_version"`
	UpdatedAt        uint64                   `json:"updated_at"`
	CreatedAt        uint64                   `json:"created_at" validate:"required"`
	CustomerId       string                   `json:"customer_id" validate:"required"`
	Type             enums.TypeEnum           `json:"type" validate:"required"`
	ReferenceId      string                   `json:"reference_id" validate:"required"`
	Status           enums.StatusEnum         `json:"status" validate:"required"`
	Gateway          enums.GatewayEnum        `json:"gateway" validate:"required"`
	GatewayAccountId string                   `json:"gateway_account_id"`
	IpAddress        string                   `json:"ip_address"`
	IssuingCountry   string                   `json:"issuing_country"`
	Deleted          bool                     `json:"deleted" validate:"required"`
	BusinessEntityId string                   `json:"business_entity_id"`
	Card             models.Card              `json:"card"`
	BankAccount      models.BankAccount       `json:"bank_account"`
	Boleto           models.CustVoucherSource `json:"boleto"`
	BillingAddress   models.BillingAddress    `json:"billing_address"`
	AmazonPayment    models.AmazonPayment     `json:"amazon_payment"`
	Upi              models.Upi               `json:"upi"`
	Paypal           models.Paypal            `json:"paypal"`
	Venmo            models.Venmo             `json:"venmo"`
	KlarnaPayNow     models.KlarnaPayNow      `json:"klarna_pay_now"`
	Mandates         []models.Mandate         `json:"mandates"`
}

func ListPaymentSourcesPageSortBy(site string, offset string, sortBy *SortBy) ([]PaymentSource, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/payment_sources")
	if err != nil {
		return nil, "", err
	}

	content, err := getQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type PaymentSourceListItem struct {
		PaymentSource PaymentSource `json:"PaymentSource"`
	}

	type PaymentSourcePage struct {
		List       []PaymentSourceListItem `json:"list"`
		NextOffset string                  `json:"next_offset,omitempty"`
	}

	entries := PaymentSourcePage{
		List: make([]PaymentSourceListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []PaymentSource{}, "", nil
	}

	result := make([]PaymentSource, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.PaymentSource)
	}

	if len(entries.NextOffset) > 0 {
		return resultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListPaymentSourcesPage(site string, offset string) ([]PaymentSource, string, error) {
	return ListPaymentSourcesPageSortBy(site, offset, nil)
}

func RetrievePaymentSourceById(site string, id string) (*PaymentSource, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/payment_sources/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQuery(client, parsedUrl, "", nil)
	if err != nil {
		return nil, err
	}

	type PaymentSourceItem struct {
		PaymentSource PaymentSource `json:"PaymentSource"`
	}

	var item PaymentSourceItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.PaymentSource, nil
}
