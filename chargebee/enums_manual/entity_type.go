package enums_manual

type EntityTypeEnum string

const (
	PlanEntity       EntityTypeEnum = "plan"
	AddonEntity                     = "addon"
	ChargeEntity                    = "charge"
	PlanPriceEntity                 = "plan_price"
	AddonPriceEntity                = "addon_price"
)
