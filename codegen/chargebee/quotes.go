package chargebee

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateListQuotes generates chargebee client code to fetch all quotes
//
// API: https://{site}.chargebee.com/api/v2/quotes
//
// Documentation: https://apidocs.chargebee.com/docs/api/quotes?lang=curl#list_quotes
func GenerateListQuotes() string {
	return client.GenerateList("Quotes", "Quote", "chargebee.com/api/v2/quotes")
}

// GenerateRetrieveQuote generates chargebee client code to retrieve specific quote
//
// API: https://{site}.chargebee.com/api/v2/quotes/{quote-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/quotes?lang=curl#retrieve_a_quote
func GenerateRetrieveQuote() string {
	return client.GenerateRetrieve("Quote", "chargebee.com/api/v2/quotes")
}

// GenerateQuotesModel generates chargebee Quote domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/quotes?lang=curl#quote_attributes
func GenerateQuotesModel() string {
	return client.GenerateModel("Quote", "Quote attributes", "https://apidocs.chargebee.com/docs/api/quotes")
}
