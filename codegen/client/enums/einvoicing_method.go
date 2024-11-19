package enums

type EinvoicingMethodEnum string

const (
	EInvoicingAutomatic   EinvoicingMethodEnum = "automatic"
	EInvoicingManual                           = "manual"
	EInvoicingSiteDefault                      = "site_default"
)
