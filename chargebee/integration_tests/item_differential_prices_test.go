package integration_tests

import (
	chargebee "github.com/bitsnap/chargebee-api-client/chargebee/generated"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee item differential prices test", func() {
	When("accessing item differential prices API", func() {
		It("should be able list item differential prices", func() {
			diffPrices, err := chargebee.ListDifferentialPrices(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(diffPrices).To(goldga.Match())
		})
	})
})
