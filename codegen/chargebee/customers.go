package chargebee

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateListCustomers generates chargebee client code to fetch all customers
//
// API: https://{site}.chargebee.com/api/v2/customers
//
// Documentation: https://apidocs.chargebee.com/docs/api/customers?lang=curl#list_customers
func GenerateListCustomers() string {
	return client.GenerateList("Customers", "Customer", "chargebee.com/api/v2/customers")
}

// GenerateRetrieveCustomer generates chargebee client code to retrieve specific customer
//
// API: https://{site}.chargebee.com/api/v2/customers/{customer_id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/customers?lang=curl#retrieve_a_customer
func GenerateRetrieveCustomer() string {
	return client.GenerateRetrieve("Customer", "chargebee.com/api/v2/customers")
}

// GenerateCustomersModel generates chargebee Customers domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/customers?lang=curl#customer_attributes
func GenerateCustomersModel() string {
	return client.GenerateModel("Customer", "Customer attributes", "https://apidocs.chargebee.com/docs/api/customers")
}
