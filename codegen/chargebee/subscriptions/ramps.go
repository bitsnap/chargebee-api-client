package subscriptions

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateListRamps generates chargebee client code to fetch all ramps
//
// API: https://{site}.chargebee.com/api/v2/ramps
//
// Documentation: https://apidocs.chargebee.com/docs/api/ramps?lang=curl#list_ramps
func GenerateListRamps() string {
	return client.GenerateList("Ramps", "Ramp", "chargebee.com/api/v2/ramps")
}

// GenerateSubscriptionRampsModel generates chargebee Ramps domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/ramps?lang=curl#list_ramps
func GenerateSubscriptionRampsModel() string {
	return client.GenerateModel("Ramp", "Ramp attributes", "https://apidocs.chargebee.com/docs/api/ramps")
}
