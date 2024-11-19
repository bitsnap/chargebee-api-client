package client_test

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
	"github.com/bitsnap/chargebee-api-client/codegen/scraper"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Model attributes code generation", func() {
	When("generating models from scraped documentation pages", func() {
		It("should be able to generate models from templates", func() {
			attrs, err := scraper.ScrapeAttributes("Entitlement attributes", "https://apidocs.chargebee.com/docs/api/entitlements")
			Expect(err).To(Not(HaveOccurred()))

			output := client.GenerateModel("Entitlement", attrs)
			Expect(output).To(goldga.Match())
		})
	})
})
