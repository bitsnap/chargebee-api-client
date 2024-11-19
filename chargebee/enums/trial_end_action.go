package enums

type TrialEndActionEnum string

const (
	TrialEndSiteDefault TrialEndActionEnum = "site_default"
	TrialEndPlanDefault                    = "plan_default"
	TrialEndActivate                       = "activate_subscription"
	TrialEndCancel                         = "cancel_subscription"
)
