package payments

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateVouchersModel generates chargebee Vouchers domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/payment_vouchers?lang=curl#payment_voucher_attributes
func GenerateVouchersModel() string {
	return client.GenerateModel("PaymentVoucher", "Payment voucher attributes", "https://apidocs.chargebee.com/docs/api/payment_vouchers")
}
