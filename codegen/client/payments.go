package client

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// GenerateListTransactions generates chargebee client code to fetch all transactions
//
// API: https://{site}.chargebee.com/api/v2/transactions
//
// Documentation: https://apidocs.chargebee.com/docs/api/transactions?lang=curl#list_transactions
func GenerateListTransactions() string {
	return templates.GenerateList("Transactions", "Transaction", "chargebee.com/api/v2/transactions")
}

// GenerateRetrieveTransaction generates chargebee client code to retrieve specific transaction
//
// API: https://{site}.chargebee.com/api/v2/transactions/{transaction-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/transactions?lang=curl#retrieve_an_transaction
func GenerateRetrieveTransaction() string {
	return templates.GenerateRetrieve("Transaction", "chargebee.com/api/v2/transactions")
}

// GenerateTransactionsModel generates chargebee Transaction domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/transactions?lang=curl#transaction_attributes
func GenerateTransactionsModel() string {
	return templates.GenerateModel("Transaction", "Transaction attributes", "https://apidocs.chargebee.com/docs/api/transactions")
}
