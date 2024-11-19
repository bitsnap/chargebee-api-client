package chargebee

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateListInvoices generates chargebee client code to fetch all invoices
//
// API: https://{site}.
//
// Documentation: https://apidocs.chargebee.com/docs/api/invoices?lang=curl#list_invoices
func GenerateListInvoices() string {
	return client.GenerateList("Invoices", "Invoice", "chargebee.com/api/v2/invoices")
}

// GenerateRetrieveInvoice generates chargebee client code to retrieve specific invoice
//
// API: https://{site}.chargebee.com/api/v2/invoices/{invoice-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/invoices?lang=curl#retrieve_an_invoice
func GenerateRetrieveInvoice() string {
	return client.GenerateRetrieve("Invoice", "chargebee.com/api/v2/invoices")
}

// GenerateInvoicesModel generates chargebee Invoice domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/invoices?lang=curl#invoice_attributes
func GenerateInvoicesModel() string {
	return client.GenerateModel("Invoice", "Invoice attributes", "https://apidocs.chargebee.com/docs/api/invoices")
}
