package enums_manual

type CancellationReasonEnum string

const (
	OrderShippingCutOff         CancellationReasonEnum = "shipping_cut_off_passed"
	OrderProductUnsatisfactory                         = "product_unsatisfactory"
	OrderThirdPartyCancellation                        = "third_party_cancellation"
	OrderProductNotRequired                            = "product_not_required"
)
