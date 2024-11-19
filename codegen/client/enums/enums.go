package enums

type ScrappedEnum struct {
	Name     string
	Variants map[string]string
	IsInt    bool
}

func (en *ScrappedEnum) String() string {
	t := "string"
	if en.IsInt {
		t = "int"
	}

	return `package chargebee
	
	// THIS IS GENERATED CODE. DO NOT EDIT.
	
	type ` + en.Name + ` ` + t + `
	const (
	
	)
	`
}

type enumScraper func() *ScrappedEnum

var content = map[string]enumScraper{
	"api_version.go":                        ScrapeApiVersionEnum,
	"apply_on.go":                           ScrapeApiVersionEnum,
	"authorization_reason.go":               ScrapeApiVersionEnum,
	"autocollection.go":                     ScrapeApiVersionEnum,
	"billing.go":                            ScrapeApiVersionEnum,
	"cancellation_reason.go":                ScrapeApiVersionEnum,
	"cancel_reason.go":                      ScrapeApiVersionEnum,
	"change_option.go":                      ScrapeApiVersionEnum,
	"channel.go":                            ScrapeApiVersionEnum,
	"charge_event.go":                       ScrapeApiVersionEnum,
	"charge_on_event.go":                    ScrapeApiVersionEnum,
	"credit_type.go":                        ScrapeApiVersionEnum,
	"customer_type.go":                      ScrapeApiVersionEnum,
	"discount_type.go":                      ScrapeApiVersionEnum,
	"dunning_status.go":                     ScrapeApiVersionEnum,
	"duration_type.go":                      ScrapeApiVersionEnum,
	"einvoicing_method.go":                  ScrapeApiVersionEnum,
	"entity_code.go":                        ScrapeApiVersionEnum,
	"entity_type.go":                        ScrapeApiVersionEnum,
	"event_type.go":                         ScrapeApiVersionEnum,
	"forex_type.go":                         ScrapeApiVersionEnum,
	"fraud_flag.go":                         ScrapeApiVersionEnum,
	"free_period_unit.go":                   ScrapeApiVersionEnum,
	"gateway.go":                            ScrapeApiVersionEnum,
	"initiator_type.go":                     ScrapeApiVersionEnum,
	"item_applicability.go":                 ScrapeApiVersionEnum,
	"item_type.go":                          ScrapeApiVersionEnum,
	"operation_type.go":                     ScrapeApiVersionEnum,
	"order_type.go":                         ScrapeApiVersionEnum,
	"payment_method.go":                     ScrapeApiVersionEnum,
	"payment_status.go":                     ScrapeApiVersionEnum,
	"payment_voucher_type.go":               ScrapeApiVersionEnum,
	"period_unit.go":                        ScrapeApiVersionEnum,
	"pii_cleared.go":                        ScrapeApiVersionEnum,
	"price_type.go":                         ScrapeApiVersionEnum,
	"pricing_model.go":                      ScrapeApiVersionEnum,
	"proration_type.go":                     ScrapeApiVersionEnum,
	"reason_code.go":                        ScrapeApiVersionEnum,
	"resent_status.go":                      ScrapeApiVersionEnum,
	"resource_type.go":                      ScrapeApiVersionEnum,
	"schedule_type.go":                      ScrapeApiVersionEnum,
	"scheme.go":                             ScrapeApiVersionEnum,
	"shipping_period_unit.go":               ScrapeApiVersionEnum,
	"source.go":                             ScrapeApiVersionEnum,
	"status.go":                             ScrapeApiVersionEnum,
	"taxability.go":                         ScrapeApiVersionEnum,
	"trial_end_action.go":                   ScrapeApiVersionEnum,
	"trial_period_unit.go":                  ScrapeApiVersionEnum,
	"type.go":                               ScrapeApiVersionEnum,
	"usage_accumulation_reset_frequency.go": ScrapeApiVersionEnum,
	"usage_calculation.go":                  ScrapeApiVersionEnum,
	"validation_status.go":                  ScrapeApiVersionEnum,
	"vat_number.go":                         ScrapeApiVersionEnum,
	"webhook_status.go":                     ScrapeApiVersionEnum,
}
