package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee customers test", func() {
	When("accessing customers API", func() {
		It("should be able list customers", func() {
			customers, err := chargebee.ListCustomers(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(customers).To(goldga.Match())
		})
	})
})
