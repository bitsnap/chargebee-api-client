package invoices

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// GenerateListInvoiceUnBilledCharges generates chargebee client code to fetch all invoice Un-billed Charges
//
// API: https://{site}.chargebee.com/api/v2/unbilled_charges
//
// Documentation: https://apidocs.chargebee.com/docs/api/unbilled_charges?lang=curl#list_unbilled_charges
func GenerateListInvoiceUnBilledCharges() string {
	return templates.GenerateList("InvoiceUnbilledCharges", "InvoiceUnbilledCharge", "chargebee.com/api/v2/unbilled_charges")
}

// GenerateInvoiceUnBilledChargeModel generates chargebee Invoice Un-billed Charge domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/unbilled_charges?lang=curl#unbilled_charge_attributes
func GenerateInvoiceUnBilledChargeModel() string {
	return templates.GenerateModel("InvoiceUnbilledCharge", "Unbilled charge attributes", "https://apidocs.chargebee.com/docs/api/unbilled_charges")
}
