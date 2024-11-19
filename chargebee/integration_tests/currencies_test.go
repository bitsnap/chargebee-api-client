package integration_tests

import (
	chargebee "github.com/bitsnap/chargebee-api-client/chargebee/generated"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee client test", func() {
	When("accessing currencies entities API", func() {
		currencies, err := chargebee.ListCurrencies(UseTestSite())

		It("should be able list currencies", func() {
			Expect(err).To(Not(HaveOccurred()))
			Expect(currencies).To(goldga.Match())
		})

		It("should be able to retrieve specific currency", func() {
			Expect(err).To(Not(HaveOccurred()))
			Expect(currencies).To(goldga.Match())
		})
	})
})
