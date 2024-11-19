package subscriptions

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// GenerateListSubscriptionEntitlements generates chargebee client code to fetch all subscription entitlements
//
// API:  https://{site}.chargebee.com/api/v2/subscriptions/{subscription-id}/subscription_entitlements
//
// Documentation: https://apidocs.chargebee.com/docs/api/subscription_entitlements?lang=curl#list_subscription_entitlements
func GenerateListSubscriptionEntitlements() string {
	return templates.GenerateListChildren("SubscriptionEntitlements", "SubscriptionEntitlement", "chargebee.com/api/v2/subscriptions", "subscription_entitlements")
}

// GenerateSubscriptionEntitlementsModel generates chargebee entitlements domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/subscription_entitlements?lang=curl#subscription_entitlement_attributes
func GenerateSubscriptionEntitlementsModel() string {
	return templates.GenerateModel("SubscriptionEntitlement", "Subscription entitlement attributes", "https://apidocs.chargebee.com/docs/api/subscription_entitlements")
}
