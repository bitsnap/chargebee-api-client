package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee customer contracts test", func() {
	When("accessing customer contracts API", func() {
		It("should be able list customer contracts", func() {
			customer, err := GetFirstCustomer()
			Expect(err).To(Not(HaveOccurred()))
			testCustomerID := customer.Id

			if customer != nil {
				contacts, err := chargebee.ListCustomerContacts(UseTestSite(), testCustomerID)
				Expect(err).To(Not(HaveOccurred()))
				Expect(contacts).To(goldga.Match())
			}
		})
	})
})
