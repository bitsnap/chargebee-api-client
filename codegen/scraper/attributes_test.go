package scraper_test

import (
	"strings"

	"github.com/bitsnap/chargebee-api-client/codegen/scraper"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

// golden files helper
type attributes []*scraper.Attribute

func (a attributes) String() string {
	sb := strings.Builder{}
	for _, attr := range a {
		sb.WriteString(attr.String())
		sb.WriteRune(',')
		sb.WriteRune(' ')
	}

	return sb.String()
}

var _ = Describe("Model attributes scraping", func() {
	When("traversing documentation pages", func() {
		It("should be able to find attribute tables with specified header text", func() {
			attrs, err := scraper.ScrapeAttributes("Customer attributes", "https://apidocs.chargebee.com/docs/api/customers#customer_attributes")
			Expect(err).To(Not(HaveOccurred()))
			Expect(attributes(attrs)).To(goldga.Match())

			//for _, attr := range attrs {
			//	for _, otherAttr := range attrs {
			//		Expect(attr.Name).ToNot(Equal(otherAttr.Name))
			//	}
			//}
		})
	})
})
