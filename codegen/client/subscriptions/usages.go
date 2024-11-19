package subscriptions

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// GenerateListSubscriptionUsages generates chargebee client code to fetch all subscription usages
//
// API: https://{site}.chargebee.com/api/v2/usages
//
// Documentation: https://apidocs.chargebee.com/docs/api/usages?lang=curl#list_usages
func GenerateListSubscriptionUsages() string {
	return templates.GenerateList("Usages", "Usage", "chargebee.com/api/v2/usages")
}

// GenerateRetrieveSubscriptionUsage generates chargebee client code to retrieve specific Subscription Usage
//
// API: https://{site}.chargebee.com/api/v2/subscription/{subscription-id}/usages
//
// Documentation: https://apidocs.chargebee.com/docs/api/usages?lang=curl#retrieve_a_usage
func GenerateRetrieveSubscriptionUsage() string {
	return templates.GenerateListChildren("SubscriptionUsages", "Usage", "chargebee.com/api/v2/subscription", "usages")
}

// GenerateSubscriptionUsagesModel generates chargebee Subscription Usages domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/usages?lang=curl#usage_attributes
func GenerateSubscriptionUsagesModel() string {
	return templates.GenerateModel("Usage", "Usage attributes", "https://apidocs.chargebee.com/docs/api/usages")
}
