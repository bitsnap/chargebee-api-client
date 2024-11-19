package customer

import "github.com/bitsnap/chargebee-api-client/codegen/templates"

// GenerateListCustomerVouchers generates chargebee client code to fetch all customer vouchers
//
// API: https://{site}.chargebee.com/api/v2/customers/{invoice-id}/payment_vouchers
//
// Documentation: https://apidocs.chargebee.com/docs/api/payment_vouchers?lang=curl#list_vouchers_for_an_invoice
func GenerateListCustomerVouchers() string {
	return templates.GenerateListChildren("CustomerPaymentVouchers", "PaymentVoucher", "chargebee.com/api/v2/customers", "payment_vouchers")
}
