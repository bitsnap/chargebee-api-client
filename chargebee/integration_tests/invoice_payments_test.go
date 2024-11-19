package integration_tests

import (
	chargebee "github.com/bitsnap/chargebee-api-client/chargebee/generated"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee invoice payments test", func() {
	When("accessing invoice payments API", func() {
		It("should be able list invoice payments", func() {
			invoice, err := GetFirstInvoice()
			Expect(err).To(Not(HaveOccurred()))

			if invoice != nil {
				testInvoiceID := invoice.Id

				payments, err := chargebee.ListInvoicePayments(UseTestSite(), testInvoiceID)
				Expect(err).To(Not(HaveOccurred()))
				Expect(payments).To(goldga.Match())
			}
		})
	})
})
