package client

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// GenerateListOrders generates chargebee client code to fetch all orders
//
// API: https://{site}.chargebee.com/api/v2/orders
//
// Documentation: https://apidocs.chargebee.com/docs/api/orders?lang=curl#list_orders
func GenerateListOrders() string {
	return templates.GenerateList("Orders", "Order", "chargebee.com/api/v2/orders")
}

// GenerateRetrieveOrder generates chargebee client code to retrieve specific order
//
// API: https://{site}.chargebee.com/api/v2/orders/{order-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/orders?lang=curl#retrieve_an_order
func GenerateRetrieveOrder() string {
	return templates.GenerateRetrieve("Order", "chargebee.com/api/v2/orders")
}

// GenerateOrdersModel generates chargebee Order domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/orders?lang=curl#order_attributes
func GenerateOrdersModel() string {
	return templates.GenerateModel("Order", "Order attributes", "https://apidocs.chargebee.com/docs/api/orders")
}
