package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee gifts", func() {
	When("accessing gifts API", func() {
		It("should be able list gifts", func() {
			gifts, err := chargebee.ListGifts(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(gifts).To(goldga.Match())
		})
	})
})
