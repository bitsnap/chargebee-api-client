package integration_tests

import (
	chargebee "github.com/bitsnap/chargebee-api-client/chargebee/generated"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee un-billed charges test", func() {
	When("accessing un-billed charges API", func() {
		It("should be able list un-billed charges", func() {
			payments, err := chargebee.ListInvoiceUnbilledCharges(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(payments).To(goldga.Match())
		})
	})
})
