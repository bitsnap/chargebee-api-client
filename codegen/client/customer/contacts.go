package customer

import "github.com/bitsnap/chargebee-api-client/codegen/templates"

// GenerateListCustomerContacts generates chargebee client code to list customer contacts
//
// API: https://{site}.chargebee.com/api/v2/customers/{customer_id}/contacts
//
// Documentation: https://apidocs.chargebee.com/docs/api/customers?lang=curl#list_of_contacts_for_a_customer
func GenerateListCustomerContacts() string {
	return templates.GenerateListChildren("CustomerContacts", "Contact", "chargebee.com/api/v2/customers", "contacts")
}
