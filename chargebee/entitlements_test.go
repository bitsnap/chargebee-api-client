package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee entitlements test", func() {
	When("accessing entitlements API", func() {
		It("should be able list entitlements", func() {
			entitlements, err := chargebee.ListEntitlements(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(entitlements).To(goldga.Match())
		})
	})
})
