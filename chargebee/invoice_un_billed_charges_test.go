package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee un-billed charges test", func() {
	When("accessing un-billed charges API", func() {
		It("should be able list un-billed charges", func() {
			payments, err := chargebee.ListUnBilledCharges(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(payments).To(goldga.Match())
		})
	})
})
