package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee subscriptions", func() {
	When("accessing subscriptions API", func() {
		It("should be able list subscriptions", func() {
			subscriptions, err := chargebee.ListSubscriptions(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(subscriptions).To(goldga.Match())
		})
	})
})
