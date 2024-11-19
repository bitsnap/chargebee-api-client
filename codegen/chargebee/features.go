package chargebee

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateListFeatures generates chargebee client code to fetch all features
//
// API: https://{site}.chargebee.com/api/v2/features
//
// Documentation: https://apidocs.chargebee.com/docs/api/features?lang=curl#list_features
func GenerateListFeatures() string {
	return client.GenerateList("Features", "Feature", "chargebee.com/api/v2/features")
}

// GenerateRetrieveFeature generates chargebee client code to retrieve specific feature
//
// API:   https://{site}.chargebee.com/api/v2/features/{feature-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/features?lang=curl#retrieve_a_feature
func GenerateRetrieveFeature() string {
	return client.GenerateRetrieve("Feature", "chargebee.com/api/v2/features")
}

// GenerateFeaturesModel generates chargebee feature domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/features?lang=curl#feature_attributes
func GenerateFeaturesModel() string {
	return client.GenerateModel("Feature", "Feature attributes", "https://apidocs.chargebee.com/docs/api/features")
}
