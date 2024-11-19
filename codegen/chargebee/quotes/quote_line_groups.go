package quotes

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateListQuoteLineGroups generates chargebee client code to fetch all quote specific line groups
//
// API: https://{site}.chargebee.com/api/v2/quotes/{quote-id}/quote_line_groups
//
// Documentation: https://apidocs.chargebee.com/docs/api/quotes?lang=curl#list_quote_line_groups
func GenerateListQuoteLineGroups() string {
	return client.GenerateListChildren("QuoteLineGroups", "QuoteLineGroup", "chargebee.com/api/v2/quotes", "quote_line_groups")
}

// GenerateListAllQuoteLineGroups generates chargebee client code to fetch all quote line groups
//
// API: https://{site}.chargebee.com/api/v2/quote_line_groups
//
// Documentation: https://apidocs.chargebee.com/docs/api/quotes?lang=curl#list_quote_line_groups
func GenerateListAllQuoteLineGroups() string {
	return client.GenerateList("AllQuoteLineGroups", "QuoteLineGroup", "chargebee.com/api/v2/quote_line_groups")
}

// GenerateRetrieveQuoteLineGroup generates chargebee client code to retrieve specific invoice installment
//
// API: https://{site}.chargebee.com/api/v2/quote_line_groups/{installment-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/quote_line_groups?prod_cat_ver=2#quote_line_group_attributes
func GenerateRetrieveQuoteLineGroup() string {
	return client.GenerateModel("QuoteLineGroup", "Quote line group attributes", "https://apidocs.chargebee.com/docs/api/quote_line_groups")
}
