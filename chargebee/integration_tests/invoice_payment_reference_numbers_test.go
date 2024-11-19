package integration_tests

import (
	chargebee "github.com/bitsnap/chargebee-api-client/chargebee/generated"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee payment reference numbers test", func() {
	When("accessing payment reference numbers API", func() {
		It("should be able list payment reference numbers", func() {
			invoice, err := GetFirstInvoice()
			Expect(err).To(Not(HaveOccurred()))

			if invoice != nil {
				testInvoiceID := invoice.Id

				numbers, err := chargebee.ListPaymentReferenceNumbers(UseTestSite(), testInvoiceID)
				Expect(err).To(Not(HaveOccurred()))
				Expect(numbers).To(goldga.Match())
			}
		})
	})
})
