package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee contract terms", func() {
	When("accessing contract terms API", func() {
		It("should be able list contract terms", func() {
			subscription, err := GetFirstSubscription()
			Expect(err).To(Not(HaveOccurred()))

			if subscription != nil {
				testSubscriptionID := subscription.Id

				addresses, err := chargebee.ListContractTerms(UseTestSite(), testSubscriptionID)
				Expect(err).To(Not(HaveOccurred()))
				Expect(addresses).To(goldga.Match())
			}
		})
	})
})
