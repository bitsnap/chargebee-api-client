package enums

type OrderTypeEnum string

const (
	OrderManual          OrderTypeEnum = "manual"
	OrderSystemGenerated               = "system_generated"
)
