package enums_manual

type ChargeEventEnum string

const (
	ChargeImmediate            ChargeEventEnum = "immediate"
	ChargeSubscriptionCreation                 = "subscription_creation"
	ChargeTrialStart                           = "trial_start"
	ChargeSubscriptionChange                   = "subscription_change"
	ChargeSubscriptionRenewal                  = "subscription_renewal"
	ChargeSubscriptionCancel                   = "subscription_cancel"
)
