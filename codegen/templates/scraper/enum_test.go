package scraper_test

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates/scraper"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Enums scraping", func() {
	When("traversing documentation pages", func() {
		It("should be able to find enum values", func() {
			v, err := scraper.ScrapeEnum("apply_on", "https://apidocs.chargebee.com/docs/api/coupons")
			Expect(err).To(Not(HaveOccurred()))
			Expect(v).To(goldga.Match())
		})
	})
})
