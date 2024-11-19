package quotes

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateQuotedChargeModel generates chargebee Quoted Charge domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/quoted_charges?lang=curl
func GenerateQuotedChargeModel() string {
	return client.GenerateModel("QuotedCharge", "Quoted charge attributes", "https://apidocs.chargebee.com/docs/api/quoted_charges")
}
