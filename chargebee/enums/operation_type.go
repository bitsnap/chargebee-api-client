package enums

type OperationTypeEnum string

const (
	CreateSubscriptionForCustomer OperationTypeEnum = "create_subscription_for_customer"
	ChangeSubscription                              = "change_subscription"
	OneTimeInvoice                                  = "onetime_invoice"
)
