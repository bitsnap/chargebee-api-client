package integration_tests

import (
	chargebee "github.com/bitsnap/chargebee-api-client/chargebee/generated"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee customer vouchers test", func() {
	When("accessing customer vouchers API", func() {
		It("should be able list customer vouchers", func() {
			customer, err := GetFirstCustomer()
			Expect(err).To(Not(HaveOccurred()))

			if customer != nil {
				testCustomerID := customer.Id

				paymentVouchers, err := chargebee.ListCustomerPaymentVouchers(UseTestSite(), testCustomerID)
				Expect(err).To(Not(HaveOccurred()))
				Expect(paymentVouchers).To(goldga.Match())
			}
		})
	})
})
