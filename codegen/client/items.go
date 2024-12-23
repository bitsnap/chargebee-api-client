package client

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// GenerateListItems generates chargebee client code to fetch all items
//
// API: https://{site}.chargebee.com/api/v2/items
//
// Documentation: https://apidocs.chargebee.com/docs/api/items?lang=curl#list_items
func GenerateListItems() string {
	return templates.GenerateList("Items", "Item", "chargebee.com/api/v2/items")
}

// GenerateRetrieveItem generates chargebee client code to retrieve specific item
//
// API:  https://{site}.chargebee.com/api/v2/items/{item-id}
//
// Documentation: hhttps://apidocs.chargebee.com/docs/api/items?lang=curl#retrieve_an_item
func GenerateRetrieveItem() string {
	return templates.GenerateRetrieve("Item", "chargebee.com/api/v2/items")
}

// GenerateItemsModel generates chargebee Item domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/items?lang=curl#item_attributes
func GenerateItemsModel() string {
	return templates.GenerateModel("Item", "Item attributes", "https://apidocs.chargebee.com/docs/api/items")
}
