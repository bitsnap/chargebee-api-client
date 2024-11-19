package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	"github.com/bitsnap/chargebee-api-client/chargebee/enums"
	"github.com/bitsnap/chargebee-api-client/chargebee/models"
	"github.com/shopspring/decimal"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Transaction struct {
	Id                       string                         `json:"id" validate:"required"`
	CustomerId               string                         `json:"customer_id"`
	SubscriptionId           string                         `json:"subscription_id"`
	GatewayAccountId         string                         `json:"gateway_account_id"`
	PaymentSourceId          string                         `json:"payment_source_id"`
	PaymentMethod            enums.PaymentMethodEnum        `json:"payment_method" validate:"required"`
	ReferenceNumber          string                         `json:"reference_number"`
	Gateway                  enums.GatewayEnum              `json:"gateway" validate:"required"`
	Type                     enums.TypeEnum                 `json:"type" validate:"required"`
	Date                     uint64                         `json:"date"`
	SettledAt                uint64                         `json:"settled_at"`
	ExchangeRate             *decimal.Decimal               `json:"exchange_rate"`
	CurrencyCode             string                         `json:"currency_code" validate:"required"`
	Amount                   uint64                         `json:"amount"`
	IdAtGateway              string                         `json:"id_at_gateway"`
	Status                   enums.StatusEnum               `json:"status"`
	FraudFlag                enums.FraudFlagEnum            `json:"fraud_flag"`
	InitiatorType            enums.InitiatorTypeEnum        `json:"initiator_type"`
	ThreeDSecure             bool                           `json:"three_d_secure"`
	AuthorizationReason      enums.AuthorizationReasonEnum  `json:"authorization_reason"`
	ErrorCode                string                         `json:"error_code"`
	ErrorText                string                         `json:"error_text"`
	VoidedAt                 uint64                         `json:"voided_at"`
	ResourceVersion          int64                          `json:"resource_version"`
	UpdatedAt                uint64                         `json:"updated_at"`
	FraudReason              string                         `json:"fraud_reason"`
	CustomPaymentMethodId    string                         `json:"custom_payment_method_id"`
	AmountUnused             uint64                         `json:"amount_unused"`
	MaskedCardNumber         string                         `json:"masked_card_number"`
	ReferenceTransactionId   string                         `json:"reference_transaction_id"`
	RefundedTxnId            string                         `json:"refunded_txn_id"`
	ReferenceAuthorizationId string                         `json:"reference_authorization_id"`
	AmountCapturable         uint64                         `json:"amount_capturable"`
	ReversalTransactionId    string                         `json:"reversal_transaction_id"`
	Deleted                  bool                           `json:"deleted" validate:"required"`
	Iin                      string                         `json:"iin"`
	Last4                    string                         `json:"last4"`
	MerchantReferenceId      string                         `json:"merchant_reference_id"`
	BusinessEntityId         string                         `json:"business_entity_id"`
	PaymentMethodDetails     string                         `json:"payment_method_details"`
	CustomPaymentMethodName  string                         `json:"custom_payment_method_name"`
	LinkedInvoices           []models.InvoiceTransaction    `json:"linked_invoices"`
	LinkedCreditNotes        []models.CreditNoteTransaction `json:"linked_credit_notes"`
	LinkedRefunds            []models.TxnRefundsAndReversal `json:"linked_refunds"`
	LinkedPayments           []models.LinkedPayment         `json:"linked_payments"`
	ErrorDetail              models.GatewayErrorDetail      `json:"error_detail"`
}

func ListTransactionsPageSortBy(site string, offset string, sortBy *SortBy) ([]Transaction, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/transactions")
	if err != nil {
		return nil, "", err
	}

	content, err := getQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type TransactionListItem struct {
		Transaction Transaction `json:"Transaction"`
	}

	type TransactionPage struct {
		List       []TransactionListItem `json:"list"`
		NextOffset string                `json:"next_offset,omitempty"`
	}

	entries := TransactionPage{
		List: make([]TransactionListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []Transaction{}, "", nil
	}

	result := make([]Transaction, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Transaction)
	}

	if len(entries.NextOffset) > 0 {
		return resultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListTransactionsPage(site string, offset string) ([]Transaction, string, error) {
	return ListTransactionsPageSortBy(site, offset, nil)
}

func RetrieveTransactionById(site string, id string) (*Transaction, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/transactions/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQuery(client, parsedUrl, "", nil)
	if err != nil {
		return nil, err
	}

	type TransactionItem struct {
		Transaction Transaction `json:"Transaction"`
	}

	var item TransactionItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.Transaction, nil
}
