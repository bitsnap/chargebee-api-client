package integration_tests

import (
	chargebee "github.com/bitsnap/chargebee-api-client/chargebee/generated"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee coupon sets test", func() {
	When("accessing coupon sets API", func() {
		It("should be able list coupon sets", func() {
			sets, err := chargebee.ListCouponSets(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(sets).To(goldga.Match())
		})
	})
})
