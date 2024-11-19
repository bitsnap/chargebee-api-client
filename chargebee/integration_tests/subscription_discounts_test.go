package integration_tests

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Chargebee discounts", func() {
	When("accessing discounts API", func() {
		It("should be able list discounts", func() {
			subscription, err := GetFirstSubscription()
			Expect(err).To(Not(HaveOccurred()))

			if subscription != nil {
				// testSubscriptionID := subscription.Id
				//
				// addresses, err := chargebee.ListSubscriptionDiscounts(UseTestSite(), testSubscriptionID)
				// Expect(err).To(Not(HaveOccurred()))
				// Expect(addresses).To(goldga.Match())
			}
		})
	})
})
