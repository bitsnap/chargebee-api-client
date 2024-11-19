package events

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// NOTE: not available for API fetch, webhooks only

// GenerateImpactedSubscriptionsModel generates chargebee entitlement domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/impacted_subscriptions?lang=curl#impacted_subscription_attributes
func GenerateImpactedSubscriptionsModel() string {
	return templates.GenerateModel("ImpactedSubscription", "Impacted subscription attributes", "https://apidocs.chargebee.com/docs/api/impacted_subscriptions")
}
