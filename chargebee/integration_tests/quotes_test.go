package integration_tests

import (
	chargebee "github.com/bitsnap/chargebee-api-client/chargebee/generated"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee quotes", func() {
	When("accessing quotes API", func() {
		It("should be able list quotes", func() {
			quotes, err := chargebee.ListQuotes(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(quotes).To(goldga.Match())
		})
	})
})
