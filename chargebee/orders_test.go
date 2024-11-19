package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee orders", func() {
	When("accessing orders API", func() {
		It("should be able list orders", func() {
			orders, err := chargebee.ListOrders(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(orders).To(goldga.Match())
		})
	})
})
