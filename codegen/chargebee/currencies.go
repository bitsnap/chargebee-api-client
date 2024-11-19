package chargebee

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateListCurrencies generates chargebee client code to fetch all currencies
//
// API: https://{site}.chargebee.com/api/v2/currencies
//
// Documentation: https://apidocs.chargebee.com/docs/api/currencies?lang=curl#list_currencies
func GenerateListCurrencies() string {
	return client.GenerateList("Currencies", "Currency", "chargebee.com/api/v2/currencies")
}

// GenerateRetrieveCurrency generates chargebee client code to retrieve specific currency
//
// API: https://{site}.chargebee.com/api/v2/currencies/{currency-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/currencies?lang=curl#retrieve_an_currency
func GenerateRetrieveCurrency() string {
	return client.GenerateRetrieve("Currency", "chargebee.com/api/v2/currencies")
}

// GenerateCurrenciesModel generates chargebee currency domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/currencies?lang=curl#currency_attributes
func GenerateCurrenciesModel() string {
	return client.GenerateModel("Currency", "Currency attributes", "https://apidocs.chargebee.com/docs/api/currencies")
}
