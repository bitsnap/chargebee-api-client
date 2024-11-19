package invoices

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// GenerateListInvoicePaymentReferenceNumbers generates chargebee client code to fetch all invoice payment reference numbers
//
// API: https://{site}.chargebee.com/api/v2/invoices/payment_reference_numbers
//
// Documentation: https://apidocs.chargebee.com/docs/api/invoices?lang=curl#list_payment_reference_numbers
func GenerateListInvoicePaymentReferenceNumbers() string {
	return templates.GenerateListChildrenWithId("PaymentReferenceNumbers", "PaymentReferenceNumber", "chargebee.com/api/v2/invoices", "payment_reference_numbers", false, "")
}

// GenerateInvoicePaymentReferenceNumberModel generates chargebee Invoice payment reference  number domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/payment_reference_numbers?lang=curl#payment_reference_number_attributes
func GenerateInvoicePaymentReferenceNumberModel() string {
	return templates.GenerateModel("PaymentReferenceNumber", "Payment reference number attributes", "https://apidocs.chargebee.com/docs/api/payment_reference_numbers")
}
