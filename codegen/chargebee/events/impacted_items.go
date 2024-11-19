package events

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// NOTE: not available for API fetch, webhooks only

// GenerateImpactedItemsModel generates chargebee entitlement domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/impacted_items?lang=curl#impacted_item_attributes
func GenerateImpactedItemsModel() string {
	return client.GenerateModel("ImpactedItem", "Impacted item attributes", "https://apidocs.chargebee.com/docs/api/impacted_items")
}
