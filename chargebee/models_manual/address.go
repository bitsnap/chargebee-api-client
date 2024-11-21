package models_manual

import (
	"github.com/bitsnap/chargebee-api-client/chargebee/enums"
	"github.com/bitsnap/chargebee-api-client/chargebee/enums_manual"
)

type Address struct {
	FirstName        string                            `json:"first_name"`
	LastName         string                            `json:"last_name"`
	Email            string                            `json:"email"`
	Company          string                            `json:"company"`
	Phone            string                            `json:"phone"`
	Line1            string                            `json:"line1"`
	Line2            string                            `json:"line2"`
	Line3            string                            `json:"line3"`
	City             string                            `json:"city"`
	StateCode        string                            `json:"state_code"`
	State            string                            `json:"state"`
	Country          string                            `json:"country"`
	Zip              string                            `json:"zip"`
	ValidationStatus enums_manual.ValidationStatusEnum `json:"validation_status"`
}
