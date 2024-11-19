package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee features test", func() {
	When("accessing features API", func() {
		It("should be able list features", func() {
			features, err := chargebee.ListFeatures(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(features).To(goldga.Match())
		})
	})
})
