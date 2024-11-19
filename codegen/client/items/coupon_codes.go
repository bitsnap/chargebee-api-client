package items

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// GenerateListItemCouponCodes generates chargebee client code to fetch all item coupon codes
//
// API:  https://{site}.chargebee.com/api/v2/coupon_codes
//
// Documentation: https://apidocs.chargebee.com/docs/api/coupons?lang=curl#list_coupons
func GenerateListItemCouponCodes() string {
	return templates.GenerateList("CouponCodes", "CouponCode", "chargebee.com/api/v2/coupon_codes")
}

// GenerateRetrieveItemCouponCode generates chargebee client code to retrieve specific item coupon code
//
// API: https://{site}.chargebee.com/api/v2/coupon_codes/{coupon-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/coupons?lang=curl#retrieve_a_coupon
func GenerateRetrieveItemCouponCode() string {
	return templates.GenerateRetrieve("CouponCode", "chargebee.com/api/v2/coupon_codes")
}

// GenerateItemCouponCodeModel generates chargebee Item domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/coupon_codes?lang=curl#coupon_code_attributes
func GenerateItemCouponCodeModel() string {
	return templates.GenerateModel("CouponCode", "Coupon code attributes", "https://apidocs.chargebee.com/docs/api/coupon_codes")
}
