package items

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateListItemCouponSets generates chargebee client code to fetch all coupon sets
//
// API: https://{site}.chargebee.com/api/v2/coupon_sets
//
// Documentation: https://apidocs.chargebee.com/docs/api/coupon_sets?lang=curl#list_coupon_sets
func GenerateListItemCouponSets() string {
	return client.GenerateList("CouponSets", "CouponSet", "chargebee.com/api/v2/coupon_sets")
}

// GenerateRetrieveItemCouponSet generates chargebee client code to retrieve specific coupon set
//
// API: https://{site}.chargebee.com/api/v2/coupon_sets/{coupon-set-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/coupon_sets?lang=curl#retrieve_a_coupon_set
func GenerateRetrieveItemCouponSet() string {
	return client.GenerateRetrieve("CouponSet", "chargebee.com/api/v2/coupon_sets")
}

// GenerateItemCouponSetModel generates chargebee coupon set domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/coupon_sets?lang=curl#coupon_set_attributes
func GenerateItemCouponSetModel() string {
	return client.GenerateModel("CouponSet", "Coupon set attributes", "https://apidocs.chargebee.com/docs/api/coupon_sets")
}
