package enums_manual

type PaymentMethodEnum string

const (
	NoPreferredPaymentMethod                      PaymentMethodEnum = "no_preference"
	CashPreferredPaymentMethod                                      = "cash"
	CheckPreferredPaymentMethod                                     = "check"
	BankTransferPaymentMethod                                       = "bank_transfer"
	ACHCreditPreferredPaymentMethod                                 = "ach_credit"
	SepaCreditPreferredPaymentMethod                                = "sepa_credit"
	BoletoPreferredPaymentMethod                                    = "boleto"
	USAutomatedBankTransferPreferredPaymentMethod                   = "us_automated_bank_transfer"
	EUAutomatedBankTransferPreferredPaymentMethod                   = "eu_automated_bank_transfer"
	UKAutomatedBankTransferPreferredPaymentMethod                   = "uk_automated_bank_transfer"
	JPAutomatedBankTransferPreferredPaymentMethod                   = "jp_automated_bank_transfer"
	MXAutomatedBankTransferPreferredPaymentMethod                   = "mx_automated_bank_transfer"
	CustomPreferredPaymentMethod                                    = "custom"
)

type OfflinePaymentMethodEnum PaymentMethodEnum
