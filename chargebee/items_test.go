package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee items", func() {
	When("accessing items API", func() {
		It("should be able list items", func() {
			items, err := chargebee.ListItems(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(items).To(goldga.Match())
		})
	})
})
