package enums

type ChargeOnEventEnum string

const (
	ChargeOnSubscriptionCreation   ChargeOnEventEnum = "subscription_creation"
	ChargeOnTrialStart                               = "subscription_trial_start"
	ChargeOnPlanActivation                           = "plan_activation"
	ChargeOnSubscriptionActivation                   = "subscription_activation"
	ChargeOnContractTermination                      = "contract_termination"
	ChargeOnDemand                                   = "on_demand"
)
