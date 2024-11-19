package enums_manual

type CustomerTypeEnum string

const (
	ResidentialCustomer   CustomerTypeEnum = "residential"
	BusinessCustomer                       = "business"
	SeniorCitizenCustomer                  = "senior_citizen"
	IndustrialCustomer                     = "industrial"
)
