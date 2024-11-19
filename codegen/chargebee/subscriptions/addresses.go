package subscriptions

//import (
//	"github.com/bitsnap/chargebee-api-client/codegen/client"
//	"github.com/bitsnap/chargebee-api-client/codegen/scraper"
//)

// TODO: subscription addresses are a bit weird, we're not going to codegen this one

// GenerateRetrieveSubscriptionAddress generates chargebee client code to retrieve specific Subscription Address
//
// API: https://{site}.chargebee.com/api/v2/addresses/{address-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/addresses?lang=curl#retrieve_an_address
//func GenerateRetrieveSubscriptionAddress() string {
//	return client.GenerateListChildrenWithId("SubscriptionAddresses", "Address", "chargebee.com/api/v2/addresses", "", false, "subscription_id")
//}

// GenerateSubscriptionAddressesModel generates chargebee Subscription Address domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/addresses?lang=curl
//func GenerateSubscriptionAddressesModel() string {
//	attributes, err := scraper.ScrapeAttributes("Address attributes", "https://apidocs.chargebee.com/docs/api/addresses")
//	if err != nil {
//		panic(err)
//	}
//
//	return client.GenerateModel("Address", attributes)
//}
