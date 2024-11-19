package scraper_test

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates/scraper"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Endpoint parameters scraping", func() {
	When("traversing documentation pages", func() {
		It("should be able to find endpoint parameters tables with specified header text", func() {
			attrs, err := scraper.ScrapeEndpointParameters("List items", "https://apidocs.chargebee.com/docs/api/items")
			Expect(err).To(Not(HaveOccurred()))
			Expect(attributes(attrs)).To(goldga.Match())
		})
	})
})
