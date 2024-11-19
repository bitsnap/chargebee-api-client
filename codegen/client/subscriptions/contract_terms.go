package subscriptions

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// GenerateListSubscriptionContractTerms generates chargebee client code to fetch all subscription contract terms
//
// API: https://{site}.chargebee.com/api/v2/subscriptions/{subscription-id}/contract_terms
//
// Documentation: https://apidocs.chargebee.com/docs/api/subscriptions?lang=curl#list_contract_terms_for_a_subscription
func GenerateListSubscriptionContractTerms() string {
	return templates.GenerateListChildren("ContractTerms", "ContractTerm", "chargebee.com/api/v2/subscriptions", "contract_terms")
}

// GenerateSubscriptionContractTermsModel generates chargebee Subscription Contract Terms domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/contract_terms?lang=curl
func GenerateSubscriptionContractTermsModel() string {
	return templates.GenerateModel("ContractTerm", "Address attributes", "https://apidocs.chargebee.com/docs/api/contract_terms")
}
