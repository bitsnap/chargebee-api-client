package items

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateListItemPrices generates chargebee client code to fetch all item prices
//
// API: https://{site}.chargebee.com/api/v2/item_prices
//
// Documentation: https://apidocs.chargebee.com/docs/api/item_prices?lang=curl
func GenerateListItemPrices() string {
	return client.GenerateList("ItemPrices", "ItemPrice", "chargebee.com/api/v2/item_prices")
}

// GenerateRetrieveItemPrice generates chargebee client code to retrieve specific item price
//
// API:  https://{site}.chargebee.com/api/v2/item_prices/{item-price-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/item_prices?lang=curl#retrieve_an_item_price
func GenerateRetrieveItemPrice() string {
	return client.GenerateRetrieve("ItemPrice", "chargebee.com/api/v2/item_prices")
}

// GenerateItemPriceModel generates chargebee Item domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/item_prices?lang=curl#item_price_attributes
func GenerateItemPriceModel() string {
	return client.GenerateModel("ItemPrice", "Item price attributes", "https://apidocs.chargebee.com/docs/api/item_prices")
}
