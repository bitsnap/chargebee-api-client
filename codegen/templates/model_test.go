package templates_test

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Model attributes code generation", func() {
	When("generating models from scraped documentation pages", func() {
		It("should be able to generate models from templates", func() {
			output := templates.GenerateModel("Entitlement", "", "https://apidocs.chargebee.com/docs/api/entitlements")
			Expect(output).To(goldga.Match())
		})
	})
})
