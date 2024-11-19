package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee discounts", func() {
	When("accessing discounts API", func() {
		It("should be able list discounts", func() {
			subscription, err := GetFirstSubscription()
			Expect(err).To(Not(HaveOccurred()))

			if subscription != nil {
				testSubscriptionID := subscription.Id

				entitlements, err := chargebee.ListSubscriptionEntitlements(UseTestSite(), testSubscriptionID)
				Expect(err).To(Not(HaveOccurred()))
				Expect(entitlements).To(goldga.Match())
			}
		})
	})
})
