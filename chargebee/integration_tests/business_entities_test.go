package integration_tests

import (
	chargebee "github.com/bitsnap/chargebee-api-client/chargebee/generated"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee client test", func() {
	When("accessing business entities API", func() {
		It("should be able list business entity transfers", func() {
			transfers, err := chargebee.ListBusinessEntityTransfers(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(transfers).To(goldga.Match())
		})
	})
})
