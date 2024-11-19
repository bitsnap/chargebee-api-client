package integration_tests

import (
	chargebee "github.com/bitsnap/chargebee-api-client/chargebee/generated"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee promotional credits test", func() {
	When("accessing promotional credits API", func() {
		It("should be able list promotional credits", func() {
			payments, err := chargebee.ListPromotionalCredits(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(payments).To(goldga.Match())
		})
	})
})
