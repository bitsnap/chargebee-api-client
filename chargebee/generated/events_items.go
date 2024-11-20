package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"github.com/bitsnap/chargebee-api-client/chargebee/models"

)

type ImpactedItem struct {
    Count int `json:"count"`
    Items []string `json:"items"`
    Download models.Download `json:"download"`
}