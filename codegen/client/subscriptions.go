package client

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// GenerateListSubscriptions generates chargebee client code to fetch all customer subscriptions
//
// API: https://{site}.chargebee.com/api/v2/subscriptions
//
// Documentation: https://apidocs.chargebee.com/docs/api/subscriptions?lang=curl#list_subscriptions
func GenerateListSubscriptions() string {
	return templates.GenerateList("Subscriptions", "Subscription", "chargebee.com/api/v2/subscriptions")
}

// GenerateRetrieveSubscription generates chargebee client code to retrieve specific subscription
//
// API:https://{site}.chargebee.com/api/v2/subscriptions/{subscription-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/subscriptions?lang=curl#retrieve_a_subscription
func GenerateRetrieveSubscription() string {
	return ""
}

// GenerateRetrieveSubscriptionWithScheduledChanges generates chargebee client code to retrieve subscription with scheduled changes
//
// API: https://{site}.chargebee.com/api/v2/subscriptions/{subscription-id}/retrieve_with_scheduled_changes
//
// Documentation: https://apidocs.chargebee.com/docs/api/subscriptions?lang=curl#retrieve_with_scheduled_changes
func GenerateRetrieveSubscriptionWithScheduledChanges() string {
	return templates.GenerateListChildren("SubscriptionWithScheduledChanges", "Subscription", "chargebee.com/api/v2/subscriptions", "retrieve_with_scheduled_changes")
}

// GenerateSubscriptionsModel generates chargebee Subscription domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/subscriptions?lang=curl#subscription_attributes
func GenerateSubscriptionsModel() string {
	return templates.GenerateModel("Subscription", "Subscription attributes", "https://apidocs.chargebee.com/docs/api/subscriptions")
}
