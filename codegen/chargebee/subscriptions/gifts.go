package subscriptions

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateRetrieveSubscriptionGift generates chargebee client code to retrieve specific gift
//
// API: https://{site}.chargebee.com/api/v2/gifts/{gift-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/gifts?lang=curl#retrieve_a_gift
func GenerateRetrieveSubscriptionGift() string {
	return client.GenerateRetrieve("Gift", "chargebee.com/api/v2/gifts")
}

// GenerateListSubscriptionGifts generates chargebee client code to fetch all gifts
//
// API: https://{site}.chargebee.com/api/v2/gifts
//
// Documentation: https://apidocs.chargebee.com/docs/api/gifts?lang=curl#list_gifts
func GenerateListSubscriptionGifts() string {
	return client.GenerateList("Gifts", "Gift", "chargebee.com/api/v2/gifts")
}

// GenerateSubscriptionGiftsModel generates chargebee Gifts domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/gifts?lang=curl#gift_attributes
func GenerateSubscriptionGiftsModel() string {
	return client.GenerateModel("Gift", "Gift attributes", "https://apidocs.chargebee.com/docs/api/gifts")
}
