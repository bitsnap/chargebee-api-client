package models

import "github.com/bitsnap/chargebee-api-client/chargebee/enums"

type PaymentMethod struct {
	Type    enums.TypeEnum    `json:"type"`
	Gateway enums.GatewayEnum `json:"gateway"`

	GatewayAccountId string `json:"gateway_account_id"`
	ReferenceId      string `json:"reference_id"`
}
