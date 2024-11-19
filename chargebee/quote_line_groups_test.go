package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee quote line groups", func() {
	When("accessing quote line groups API", func() {
		It("should be able list quote line groups", func() {
			quote, err := GetFirstQuote()
			Expect(err).To(Not(HaveOccurred()))

			if quote != nil {
				testQuoteID := quote.Id
				lineGroups, err := chargebee.ListQuoteLineGroups(UseTestSite(), testQuoteID)
				Expect(err).To(Not(HaveOccurred()))
				Expect(lineGroups).To(goldga.Match())
			}
		})
	})
})
