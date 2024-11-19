package enums

import "slices"

type GatewayEnum string

const (
	TestChargeBeeGateway          GatewayEnum = "chargebee"
	ChargeBeeGateway                          = "chargebee_payments"
	StripeGateway                             = "stripe"
	WepayGateway                              = "wepay"
	BrainTreeGateway                          = "braintree"
	AuthorizeNetGateway                       = "authorize_net"
	PaypalProGateway                          = "paypal_pro"
	PinGateway                                = "pin"
	EWayGateway                               = "eway"
	EWayRapidGateway                          = "eway_rapid"
	WorldpayGateway                           = "worldpay"
	BalancedPaymentsGateway                   = "balanced_payments"
	BeanStreamGateway                         = "beanstream"
	BamboraGateway                            = "Bambora"
	BluePayGateway                            = "bluepay"
	ElavonGateway                             = "elavon"
	FirstDataGlobalGateway                    = "first_data_global"
	HDFCGateway                               = "hdfc"
	MIGSGateway                               = "migs"
	NMIGateway                                = "nmi"
	OGoneGateway                              = "ogone"
	PayMillGateway                            = "paymill"
	PayPalPayflowPro                          = "paypal_payflow_pro"
	SagePayGateway                            = "sage_pay"
	TCOGateway                                = "tco"
	WireCardGateway                           = "wirecard"
	AmazonPaymentsGateway                     = "amazon_payments"
	PaypallExpressCheckoutGateway             = "paypal_express_checkout"
	GocardlessGateway                         = "gocardless"
	AdyenGateway                              = "adyen"
	OrbitalGateway                            = "orbital"
	MonerisUSGateway                          = "moneris_us"
	MonerisGateway                            = "moneris_"
	BluesnapGateway                           = "bluesnap"
	CyberSourceGateway                        = "cybersource"
	VantivGateway                             = "vantiv"
	CheckoutComGateway                        = "checkout_com"
	PaypalGateway                             = "paypal"
	IngenicoDirectGateway                     = "ingenico_direct"
	ExactGateway                              = "exact"
	MollieGateway                             = "mollie"
	QuickBooksGateway                         = "quickbooks"
	RazorpayGateway                           = "razorpay"
	GlobalPaymentsGateway                     = "global_payments"
	BankOfAmericaGateway                      = "bank_of_america"
	EcentricGateway                           = "ecentric"
	MetricsGlobalGateway                      = "metrics_global"
	WindcaveGateway                           = "windcave"
	PayCom                                    = "pay_com"
	EbanxGateway                              = "ebanx"
	DLocalGateway                             = "dlocal"
	NuveiGateway                              = "nuvei"
	NotApplicableGateway                      = "not_applicable"
)

func KnownPaymentGateway(name string) bool {
	return slices.Contains([]GatewayEnum{
		TestChargeBeeGateway,
		ChargeBeeGateway,
		StripeGateway,
		WepayGateway,
		BrainTreeGateway,
		AuthorizeNetGateway,
		PaypalProGateway,
		PinGateway,
		EWayGateway,
		EWayRapidGateway,
		WorldpayGateway,
		BalancedPaymentsGateway,
		BeanStreamGateway,
		BamboraGateway,
		BluePayGateway,
		ElavonGateway,
		FirstDataGlobalGateway,
		HDFCGateway,
		MIGSGateway,
		NMIGateway,
		OGoneGateway,
		PayMillGateway,
		PayPalPayflowPro,
		SagePayGateway,
		TCOGateway,
		WireCardGateway,
		AmazonPaymentsGateway,
		PaypallExpressCheckoutGateway,
		GocardlessGateway,
		AdyenGateway,
		OrbitalGateway,
		MonerisUSGateway,
		MonerisGateway,
		BluesnapGateway,
		CyberSourceGateway,
		VantivGateway,
		CheckoutComGateway,
		PaypalGateway,
		IngenicoDirectGateway,
		ExactGateway,
		MollieGateway,
		QuickBooksGateway,
		RazorpayGateway,
		GlobalPaymentsGateway,
		BankOfAmericaGateway,
		EcentricGateway,
		MetricsGlobalGateway,
		WindcaveGateway,
		PayCom,
		EbanxGateway,
		DLocalGateway,
		NuveiGateway,
		NotApplicableGateway,
	}, GatewayEnum(name))
}
