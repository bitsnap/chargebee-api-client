package payments

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// GenerateListPaymentSources generates chargebee client code to fetch all payment sources
//
// API: https://{site}.chargebee.com/api/v2/payment_sources
//
// Documentation: https://apidocs.chargebee.com/docs/api/payment_sources?lang=curl#list_payment_sources
func GenerateListPaymentSources() string {
	return templates.GenerateList("PaymentSources", "PaymentSource", "chargebee.com/api/v2/payment_sources")
}

// GenerateRetrievePaymentSource generates chargebee client code to retrieve specific payment source
//
// API: https://{site}.chargebee.com/api/v2/payment_sources/{cust-payment-source-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/payment_sources?lang=curl#retrieve_a_payment_source
func GenerateRetrievePaymentSource() string {
	return templates.GenerateRetrieve("PaymentSource", "chargebee.com/api/v2/payment_sources")
}

// GeneratePaymentSourcesModel generates chargebee payment source domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/payment_sources?lang=curl#payment_source_attributes
func GeneratePaymentSourcesModel() string {
	return templates.GenerateModel("PaymentSource", "Payment source attributes", "https://apidocs.chargebee.com/docs/api/payment_sources")
}
