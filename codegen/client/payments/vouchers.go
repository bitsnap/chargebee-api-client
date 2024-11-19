package payments

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// GenerateVouchersModel generates chargebee Vouchers domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/payment_vouchers?lang=curl#payment_voucher_attributes
func GenerateVouchersModel() string {
	return templates.GenerateModel("PaymentVoucher", "Payment voucher attributes", "https://apidocs.chargebee.com/docs/api/payment_vouchers")
}
