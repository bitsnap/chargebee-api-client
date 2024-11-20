package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"github.com/bitsnap/chargebee-api-client/chargebee/models"

)

type ImpactedSubscription struct {
    Count int `json:"count"`
    SubscriptionIds []string `json:"subscription_ids"`
    Download models.Download `json:"download"`
}