package invoices

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateListInvoicePromotionalCredits generates chargebee client code to fetch all invoice promotional credits
//
// API: https://{site}.chargebee.com/api/v2/promotional_credits
//
// Documentation: https://apidocs.chargebee.com/docs/api/promotional_credits?lang=curl#list_promotional_credits
func GenerateListInvoicePromotionalCredits() string {
	return client.GenerateList("PromotionalCredits", "PromotionalCredit", "chargebee.com/api/v2/promotional_credits")
}

// GenerateRetrieveInvoicePromotionalCredit generates chargebee client code to retrieve specific invoice promotional credit
//
// API: https://{site}.chargebee.com/api/v2/promotional_credits/{account-credit-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/promotional_credits?lang=curl#retrieve_a_promotional_credit
func GenerateRetrieveInvoicePromotionalCredit() string {
	return client.GenerateRetrieve("PromotionalCredit", "chargebee.com/api/v2/promotional_credits")
}

// GenerateInvoicePromotionalCreditModel generates chargebee Invoice promotional credit domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/promotional_credits?lang=curl#promotional_credit_attributes
func GenerateInvoicePromotionalCreditModel() string {
	return client.GenerateModel("PromotionalCredit", "Promotional credit attributes", "https://apidocs.chargebee.com/docs/api/promotional_credits")
}
