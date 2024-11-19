package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
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
