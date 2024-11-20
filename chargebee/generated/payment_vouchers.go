package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"github.com/bitsnap/chargebee-api-client/chargebee/enums"
	"github.com/bitsnap/chargebee-api-client/chargebee/models"

)

type PaymentVoucher struct {
    Id string `json:"id" validate:"required"`
    IdAtGateway string `json:"id_at_gateway"`
    PaymentVoucherType enums.PaymentVoucherTypeEnum `json:"payment_voucher_type" validate:"required"`
    ExpiresAt uint64 `json:"expires_at"`
    Status enums.StatusEnum `json:"status"`
    SubscriptionId string `json:"subscription_id"`
    CurrencyCode string `json:"currency_code" validate:"required"`
    Amount uint64 `json:"amount"`
    GatewayAccountId string `json:"gateway_account_id"`
    PaymentSourceId string `json:"payment_source_id"`
    Gateway enums.GatewayEnum `json:"gateway" validate:"required"`
    Payload string `json:"payload"`
    ErrorCode string `json:"error_code"`
    ErrorText string `json:"error_text"`
    Url string `json:"url"`
    Date uint64 `json:"date"`
    ResourceVersion int64 `json:"resource_version"`
    UpdatedAt uint64 `json:"updated_at"`
    CustomerId string `json:"customer_id" validate:"required"`
    LinkedInvoices []models.InvoicePaymentVoucher `json:"linked_invoices"`
}