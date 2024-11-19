package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee subscription address", func() {
	When("accessing subscription address API", func() {
		It("should be able list subscription address", func() {
			subscription, err := GetFirstSubscription()
			Expect(err).To(Not(HaveOccurred()))

			if subscription != nil {
				testSubscriptionID := subscription.Id

				shippingAddr, err := chargebee.SubscriptionShippingAddress(UseTestSite(), testSubscriptionID)
				Expect(err).To(Not(HaveOccurred()))
				if shippingAddr != nil {
					Expect(shippingAddr).To(goldga.Match())
				}

				billingAddr, err := chargebee.SubscriptionBillingAddress(UseTestSite(), testSubscriptionID)
				Expect(err).To(Not(HaveOccurred()))
				if billingAddr != nil {
					Expect(billingAddr).To(goldga.Match())
				}
			}
		})
	})
})
