package enums

type VatNumberStatusEnum string

const (
	VATValid        VatNumberStatusEnum = "valid"
	VATInvalid                          = "invalid"
	VATNotValidated                     = "not_validated"
	VATUndetermined                     = "undetermined"
)
