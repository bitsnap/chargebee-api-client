package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee virtual bank accounts", func() {
	When("accessing virtual bank accounts API", func() {
		It("should be able list virtual bank accounts", func() {
			accounts, err := chargebee.ListVirtualBankAccounts(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(accounts).To(goldga.Match())
		})
	})
})
