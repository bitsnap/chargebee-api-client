package chargebee_test

import (
	"strings"

	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee customer hierarchies test", func() {
	When("accessing customer hierarchies API", func() {
		It("should be able list customer hierarchies", func() {
			customer, err := GetFirstCustomer()
			Expect(err).To(Not(HaveOccurred()))

			if customer != nil {
				testCustomerID := customer.Id

				hierarchy, err := chargebee.ListCustomerHierarchies(UseTestSite(), testCustomerID)
				if !strings.Contains(err.Error(), "enable Account Hierarchy") {
					Expect(err).To(Not(HaveOccurred()))
					Expect(hierarchy).To(goldga.Match())
				}
			}
		})
	})
})
