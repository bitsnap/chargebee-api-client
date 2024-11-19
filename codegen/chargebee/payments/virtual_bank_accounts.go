package payments

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateListVirtualBankAccounts generates chargebee client code to fetch all virtual bank accounts
//
// API: https://{site}.chargebee.com/api/v2/virtual_bank_accounts
//
// Documentation: https://apidocs.chargebee.com/docs/api/virtual_bank_accounts?lang=curl#list_virtual_bank_accounts
func GenerateListVirtualBankAccounts() string {
	return client.GenerateList("VirtualBankAccounts", "VirtualBankAccount", "chargebee.com/api/v2/virtual_bank_accounts")
}

// GenerateRetrieveVirtualBankAccount generates chargebee client code to retrieve specific virtual bank account
//
// API: https://{site}.chargebee.com/api/v2/virtual_bank_accounts/{virtual-bank-account-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/virtual_bank_accounts?lang=curl#retrieve_a_virtual_bank_account
func GenerateRetrieveVirtualBankAccount() string {
	return client.GenerateRetrieve("VirtualBankAccount", "chargebee.com/api/v2/virtual_bank_accounts")
}

// GenerateVirtualBankAccountsModel generates chargebee virtual bank account domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/virtual_bank_accounts?lang=curl#virtual_bank_account_attributes
func GenerateVirtualBankAccountsModel() string {
	return client.GenerateModel("VirtualBankAccount", "Virtual bank account attributes", "https://apidocs.chargebee.com/docs/api/virtual_bank_accounts")
}
