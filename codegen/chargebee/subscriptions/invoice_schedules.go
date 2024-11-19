package subscriptions

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateListSubscriptionAdvancedInvoiceSchedule generates chargebee client code to list subscription advanced invoice schedules
//
// API: https://{site}.chargebee.com/api/v2/subscriptions/{subscription-id}/retrieve_advance_invoice_schedules
//
// Documentation: https://apidocs.chargebee.com/docs/api/subscriptions?lang=curl#retrieve_advance_invoice
func GenerateListSubscriptionAdvancedInvoiceSchedule() string {
	return client.GenerateListChildren("AdvancedInvoiceSchedules", "AdvancedInvoiceSchedule", "chargebee.com/api/v2/subscriptions", "retrieve_advance_invoice_schedule")
}

// GenerateSubscriptionAdvancedInvoiceScheduleModel generates chargebee subscription advanced invoice schedule domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/advance_invoice_schedules?prod_cat_ver=2
func GenerateSubscriptionAdvancedInvoiceScheduleModel() string {
	return client.GenerateModel("AdvancedInvoiceSchedule", "Advance invoice schedule attributes", "https://apidocs.chargebee.com/docs/api/advance_invoice_schedules")
}
