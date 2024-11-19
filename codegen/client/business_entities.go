package client

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// GenerateListBusinessEntityTransfers generates chargebee client code to fetch all business entity transfers
//
// API: https://{site}.chargebee.com/api/v2/business_entities/transfers
//
// Documentation: https://apidocs.chargebee.com/docs/api/business_entities?lang=curl#list_the_business_entity_transfers
func GenerateListBusinessEntityTransfers() string {
	return templates.GenerateList("BusinessEntityTransfers", "BusinessEntityTransfer", "chargebee.com/api/v2/business_entities/transfers")
}

// GenerateBusinessEntityTransferModel generates chargebee BusinessEntityTransfer domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/business_entities?lang=curl#business_entity_attributes
func GenerateBusinessEntityTransferModel() string {
	return templates.GenerateModel("BusinessEntityTransfer", "Business entity transfer attributes", "https://apidocs.chargebee.com/docs/api/business_entity_transfers")
}

// GenerateBusinessEntitiesModel generates chargebee BusinessEntity domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/business_entities?lang=curl#business_entity_attributes
func GenerateBusinessEntitiesModel() string {
	return templates.GenerateModel("BusinessEntity", "Business entity attributes", "https://apidocs.chargebee.com/docs/api/business_entities")
}
