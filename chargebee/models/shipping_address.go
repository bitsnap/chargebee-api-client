package models

type ShippingAddress struct {
	Index int `json:"index,omitempty"`
	// Label            string                     `json:"label" validate:"required"`
	// SubscriptionId   string                     `json:"subscription_id" validate:"required"`
	Address
}
