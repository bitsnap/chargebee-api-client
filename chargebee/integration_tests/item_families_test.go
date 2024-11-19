package integration_tests

import (
	chargebee "github.com/bitsnap/chargebee-api-client/chargebee/generated"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee item families test", func() {
	When("accessing item families API", func() {
		It("should be able list item families", func() {
			families, err := chargebee.ListItemFamilies(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(families).To(goldga.Match())
		})
	})
})
