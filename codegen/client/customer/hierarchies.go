package customer

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// GenerateRetrieveCustomerAccountHierarchy generates chargebee client code to list customer hierarchies
//
// API: https://{site}.chargebee.com/api/v2/customers/{customer_id}/hierarchy
//
// Documentation: https://apidocs.chargebee.com/docs/api/customers?lang=curl#get_hierarchy
//
//	https://apidocs.chargebee.com/docs/api/hierarchies?lang=curl
func GenerateRetrieveCustomerAccountHierarchy() string {
	return templates.GenerateListChildren("CustomerHierarchies", "Hierarchy", "chargebee.com/api/v2/customers", "hierarchy")
}

// GenerateHierarchyModel generates chargebee contacts domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/hierarchies?lang=curl
func GenerateHierarchyModel() string {
	return templates.GenerateModel("Hierarchy", "Hierarchy attributes", "https://apidocs.chargebee.com/docs/api/hierarchies")
}
