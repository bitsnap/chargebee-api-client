package items

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateListItemAttached generates chargebee client code to fetch all item attached
//
// API: https://{site}.chargebee.com/api/v2/attached_items
//
// Documentation: https://apidocs.chargebee.com/docs/api/attached_items?lang=curl
func GenerateListItemAttached() string {
	return client.GenerateListChildren("ItemsAttached", "ItemAttached", "chargebee.com/api/v2/items", "attached_items")
}

// GenerateRetrieveItemAttached generates chargebee client code to retrieve specific item attached
//
// API:  https://{site}.chargebee.com/api/v2/item_prices/{item-price-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/item_prices?lang=curl#retrieve_an_item_price
func GenerateRetrieveItemAttached() string {
	return client.GenerateRetrieve("ItemAttached", "chargebee.com/api/v2/attached_items")
}

// GenerateItemAttachedModel generates chargebee Attached Item domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/attached_items?lang=curl#attached_item_attributes
func GenerateItemAttachedModel() string {
	return client.GenerateModel("ItemAttached", "Attached item attributes", "https://apidocs.chargebee.com/docs/api/attached_items")
}
