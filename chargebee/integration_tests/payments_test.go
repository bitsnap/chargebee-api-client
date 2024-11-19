package integration_tests

import (
	chargebee "github.com/bitsnap/chargebee-api-client/chargebee/generated"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee payments", func() {
	When("accessing payments API", func() {
		It("should be able list payments", func() {
			transactions, err := chargebee.ListTransactions(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(transactions).To(goldga.Match())
		})
	})
})
