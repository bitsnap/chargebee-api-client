package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee coupons test", func() {
	When("accessing coupons API", func() {
		It("should be able list coupons", func() {
			coupons, err := chargebee.ListCoupons(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(coupons).To(goldga.Match())
		})
	})
})
