package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee coupon codes test", func() {
	When("accessing coupon codes API", func() {
		It("should be able list coupon codes", func() {
			codes, err := chargebee.ListCouponCodes(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(codes).To(goldga.Match())
		})
	})
})
