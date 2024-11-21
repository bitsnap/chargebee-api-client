package enums_manual

type PricingModelEnum string

const (
	FlatFeePricing   PriceTypeEnum = "flat_fee"
	PerUnitPricing                 = "per_unit"
	TieredPricing                  = "tiered"
	VolumePricing                  = "volume"
	StairStepPricing               = "stair_step"
)
