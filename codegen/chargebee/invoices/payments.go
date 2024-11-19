package invoices

import "github.com/bitsnap/chargebee-api-client/codegen/client"

// GenerateListInvoicePayments generates chargebee client code to fetch all invoice payments
//
// API: https://{site}.chargebee.com/api/v2/invoices/{invoice-id}/payments
//
// Documentation: https://apidocs.chargebee.com/docs/api/transactions?lang=curl#list_payments_for_an_invoice
func GenerateListInvoicePayments() string {
	return client.GenerateListChildren("InvoicePayments", "Transaction", "chargebee.com/api/v2/invoices", "payments")
}
