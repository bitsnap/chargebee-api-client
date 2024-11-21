package enums_manual

type DunningStatusEnum string

const (
	InvoiceInProgress DunningStatusEnum = "in_progress"
	InvoiceExhausted                    = "exhausted"
	InvoiceStopped                      = "stopped"
	InvoiceSuccess                      = "success"
)