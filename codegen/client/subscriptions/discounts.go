package subscriptions

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// GenerateListSubscriptionDiscounts generates chargebee client code to fetch all subscription discrounts
//
// API: https://{site}.chargebee.com/api/v2/subscriptions/{subscription-id}/discounts
//
// Documentation: https://apidocs.chargebee.com/docs/api/subscriptions?lang=curl#list_discounts_for_a_subscription
func GenerateListSubscriptionDiscounts() string {
	return templates.GenerateListChildren("Discounts", "Discount", "chargebee.com/api/v2/subscriptions", "discounts")
}

// GenerateSubscriptionDiscountsModel generates chargebee Subscription Discounts domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/discounts?lang=curl#discount_attributes
func GenerateSubscriptionDiscountsModel() string {
	return templates.GenerateModel("Discount", "Discount attributes", "https://apidocs.chargebee.com/docs/api/discounts")
}
