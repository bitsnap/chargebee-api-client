package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee subscription usages", func() {
	When("accessing subscription usages API", func() {
		It("should be able list subscription usages", func() {
			usages, err := chargebee.ListUsages(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(usages).To(goldga.Match())
		})
	})
})
