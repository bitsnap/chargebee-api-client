package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

type Contact struct {
	Id               string `json:"id" validate:"required"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email" validate:"required"`
	Phone            string `json:"phone"`
	Label            string `json:"label"`
	Enabled          bool   `json:"enabled" validate:"required"`
	SendAccountEmail bool   `json:"send_account_email" validate:"required"`
	SendBillingEmail bool   `json:"send_billing_email" validate:"required"`
}
