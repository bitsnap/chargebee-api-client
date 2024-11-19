package invoices

import "github.com/bitsnap/chargebee-api-client/codegen/client"

// GenerateListInvoiceVouchers generates chargebee client code to fetch all invoice vouchers
//
// API: https://{site}.chargebee.com/api/v2/invoices/{invoice-id}/payment_vouchers
//
// Documentation: https://apidocs.chargebee.com/docs/api/payment_vouchers?lang=curl#list_vouchers_for_an_invoice
func GenerateListInvoiceVouchers() string {
	return client.GenerateListChildren("InvoicePaymentVouchers", "PaymentVoucher", "chargebee.com/api/v2/invoices", "payment_vouchers")
}
