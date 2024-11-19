package chargebee

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateListEntitlements generates chargebee client code to fetch all items
//
// API: https://{site}.chargebee.com/api/v2/entitlements
//
// Documentation: https://apidocs.chargebee.com/docs/api/entitlements?lang=curl#list_all_entitlements
func GenerateListEntitlements() string {
	return client.GenerateList("Entitlements", "Entitlement", "chargebee.com/api/v2/entitlements")
}

// GenerateEntitlementsModel generates chargebee entitlement domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/entitlements?lang=curl#entitlement_attributes
func GenerateEntitlementsModel() string {
	return client.GenerateModel("Entitlement", "Entitlement attributes", "https://apidocs.chargebee.com/docs/api/entitlements")
}
