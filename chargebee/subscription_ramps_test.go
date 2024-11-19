package chargebee_test

import (
	"strings"

	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee ramps", func() {
	When("accessing ramps API", func() {
		It("should be able list ramps", func() {
			ramps, err := chargebee.ListRamps(UseTestSite())

			if !strings.Contains(err.Error(), "is not enabled") {
				Expect(err).To(Not(HaveOccurred()))
				Expect(ramps).To(goldga.Match())
			}
		})
	})
})
