package items

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// GenerateListItemCoupons generates chargebee client code to fetch all item coupons
//
// API:  https://{site}.chargebee.com/api/v2/coupons
//
// Documentation: https://apidocs.chargebee.com/docs/api/coupons?lang=curl#list_coupons
func GenerateListItemCoupons() string {
	return templates.GenerateList("Coupons", "Coupon", "chargebee.com/api/v2/coupons")
}

// GenerateRetrieveItemCoupon generates chargebee client code to retrieve specific item coupon
//
// API: https://{site}.chargebee.com/api/v2/coupon_codes/{coupon-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/coupons?lang=curl#retrieve_a_coupon
func GenerateRetrieveItemCoupon() string {
	return templates.GenerateRetrieve("Coupon", "chargebee.com/api/v2/coupons")
}

// GenerateItemCouponModel generates chargebee Item domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/coupons?lang=curl#coupon_attributes
func GenerateItemCouponModel() string {
	return templates.GenerateModel("Coupon", "Coupon attributes", "https://apidocs.chargebee.com/docs/api/coupons")
}
