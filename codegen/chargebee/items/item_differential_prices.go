package items

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateListItemDifferentialPrices generates chargebee client code to fetch all item differential prices
//
// API: https://{site}.chargebee.com/api/v2/item_prices
//
// Documentation: https://apidocs.chargebee.com/docs/api/differential_prices?lang=curl#list_differential_prices
func GenerateListItemDifferentialPrices() string {
	return client.GenerateList("DifferentialPrices", "DifferentialPrice", "chargebee.com/api/v2/item_prices")
}

// GenerateRetrieveItemDifferentialPrice generates chargebee client code to retrieve specific item differential price
//
// API:  https://{site}.chargebee.com/api/v2/item_prices/{item-price-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/differential_prices?lang=curl#retrieve_a_differential_price
func GenerateRetrieveItemDifferentialPrice() string {
	return client.GenerateRetrieve("DifferentialPrice", "chargebee.com/api/v2/item_prices")
}

// GenerateItemDifferentialPriceModel generates chargebee Item domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/differential_prices?lang=curl#differential_price_attributes
func GenerateItemDifferentialPriceModel() string {
	return client.GenerateModel("DifferentialPrice", "Differential price attributes", "https://apidocs.chargebee.com/docs/api/differential_prices")
}
