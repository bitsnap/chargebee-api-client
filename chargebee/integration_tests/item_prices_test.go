package integration_tests

import (
	chargebee "github.com/bitsnap/chargebee-api-client/chargebee/generated"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee item prices test", func() {
	When("accessing item prices API", func() {
		It("should be able list item prices", func() {
			prices, err := chargebee.ListItemPrices(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(prices).To(goldga.Match())
		})
	})
})
