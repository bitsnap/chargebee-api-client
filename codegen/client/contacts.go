package client

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// GenerateListContacts generates chargebee client code to fetch all currencies
//
// API: https://{site}.chargebee.com/api/v2/contacts
//
// Documentation: https://apidocs.chargebee.com/docs/api/currencies?lang=curl#list_currencies
func GenerateListContacts() string {
	return templates.GenerateList("Contacts", "Contact", "chargebee.com/api/v2/contacts")
}

// GenerateContactsModel generates chargebee contacts domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/contacts?prod_cat_ver=2#contact_attributes
func GenerateContactsModel() string {
	return templates.GenerateModel("Contact", "Contact attributes", "https://apidocs.chargebee.com/docs/api/contacts")
}
