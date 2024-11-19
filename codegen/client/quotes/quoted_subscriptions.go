package quotes

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// GenerateQuotedSubscriptionModel generates chargebee Quoted Subscription domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/quoted_subscriptions?lang=curl
func GenerateQuotedSubscriptionModel() string {
	return templates.GenerateModel("QuotedSubscription", "Quoted subscription attributes", "https://apidocs.chargebee.com/docs/api/quoted_subscriptions")
}
