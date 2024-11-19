package enums_manual

type VatNumberStatusEnum string

const (
	VATValid        VatNumberStatusEnum = "valid"
	VATInvalid                          = "invalid"
	VATNotValidated                     = "not_validated"
	VATUndetermined                     = "undetermined"
)
