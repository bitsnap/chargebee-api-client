package integration_tests

import (
	chargebee "github.com/bitsnap/chargebee-api-client/chargebee/generated"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Chargebee invoices test", func() {
	When("accessing invoices API", func() {
		It("should be able list invoices", func() {
			payments, err := chargebee.ListInvoices(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(payments).NotTo(BeNil())
		})
	})
})
