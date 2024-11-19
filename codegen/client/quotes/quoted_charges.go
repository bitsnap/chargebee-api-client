package quotes

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// GenerateQuotedChargeModel generates chargebee Quoted Charge domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/quoted_charges?lang=curl
func GenerateQuotedChargeModel() string {
	return templates.GenerateModel("QuotedCharge", "Quoted charge attributes", "https://apidocs.chargebee.com/docs/api/quoted_charges")
}
