package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee events test", func() {
	When("accessing events API", func() {
		It("should be able list events", func() {
			events, err := chargebee.ListEvents(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(events).To(goldga.Match())
		})
	})
})
