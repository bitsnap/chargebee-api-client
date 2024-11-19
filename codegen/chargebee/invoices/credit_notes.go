package invoices

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateListInvoiceCreditNotes generates chargebee client code to fetch all invoice credit notes
//
// API: https://{site}.chargebee.com/api/v2/credit_notes
//
// Documentation: https://apidocs.chargebee.com/docs/api/credit_notes?lang=curl#list_credit_notes
func GenerateListInvoiceCreditNotes() string {
	return client.GenerateList("CreditNotes", "CreditNote", "chargebee.com/api/v2/credit_notes")
}

// GenerateRetrieveInvoiceCreditNote generates chargebee client code to retrieve specific invoice credit note
//
// API: https://{site}.chargebee.com/api/v2/credit_notes
//
// Documentation: https://apidocs.chargebee.com/docs/api/credit_notes?lang=curl#retrieve_a_credit_note
func GenerateRetrieveInvoiceCreditNote() string {
	return client.GenerateRetrieve("CreditNote", "chargebee.com/api/v2/credit_notes")
}

// GenerateInvoiceCreditNoteModel generates chargebee Invoice Credit Note domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/credit_notes?lang=curl#credit_note_attributes
func GenerateInvoiceCreditNoteModel() string {
	return client.GenerateModel("CreditNote", "Credit note attributes", "https://apidocs.chargebee.com/docs/api/credit_notes")
}
