package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee invoice schedules", func() {
	When("accessing invoice schedules API", func() {
		It("should be able list invoice schedules", func() {
			subscription, err := GetFirstSubscription()
			Expect(err).To(Not(HaveOccurred()))

			if subscription != nil {
				testSubscriptionID := subscription.Id

				invoiceSchedules, err := chargebee.ListAdvancedInvoiceSchedules(UseTestSite(), testSubscriptionID)
				Expect(err).To(Not(HaveOccurred()))
				Expect(invoiceSchedules).To(goldga.Match())
			}
		})
	})
})
