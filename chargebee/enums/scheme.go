package enums

type SchemeEnum string

const (
	ACHCreditScheme         SchemeEnum = "ach_credit"
	SepaCreditScheme                   = "sepa_credit"
	USAutomatedBankTransfer            = "us_automated_bank_transfer"
	GBAutomatedBankTransfer            = "gb_automated_bank_transfer"
	EUAutomatedBankTransfer            = "eu_automated_bank_transfer"
	JPAutomatedBankTransfer            = "jp_automated_bank_transfer"
	MXAutomatedBankTransfer            = "mx_automated_bank_transfer"
)
