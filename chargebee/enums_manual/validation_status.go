package enums_manual

type ValidationStatusEnum string

const (
	AddressNotValidated   ValidationStatusEnum = "not_validated"
	AddressValid                               = "valid"
	AddressPartiallyValid                      = "partially_valid"
	AddressInvalid                             = "invalid"
)
