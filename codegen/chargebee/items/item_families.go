package items

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateListItemFamilies generates chargebee client code to fetch all item families
//
// API: https://{site}.chargebee.com/api/v2/item_families
//
// Documentation: https://apidocs.chargebee.com/docs/api/item_families?lang=curl#list_item_families
func GenerateListItemFamilies() string {
	return client.GenerateList("ItemFamilies", "ItemFamily", "chargebee.com/api/v2/item_families")
}

// GenerateRetrieveItemFamily generates chargebee client code to retrieve specific item family
//
// API: https://{site}.chargebee.com/api/v2/item_families/{item-family-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/item_families?lang=curl#retrieve_an_item_family
func GenerateRetrieveItemFamily() string {
	return client.GenerateRetrieve("ItemFamily", "chargebee.com/api/v2/item_families")
}

// GenerateItemFamilyModel generates chargebee Item domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/item_families?lang=curl#item_family_attributes
func GenerateItemFamilyModel() string {
	return client.GenerateModel("ItemFamily", "Item family attributes", "https://apidocs.chargebee.com/docs/api/item_families")
}
