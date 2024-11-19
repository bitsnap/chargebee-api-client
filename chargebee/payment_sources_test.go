package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee payment sources", func() {
	When("accessing payment sources API", func() {
		It("should be able list payment sources", func() {
			sources, err := chargebee.ListPaymentSources(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(sources).To(goldga.Match())
		})
	})
})
