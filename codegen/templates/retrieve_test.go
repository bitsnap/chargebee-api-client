package templates_test

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Model attributes code generation", func() {
	When("generating retrieve endpoints client from scraped documentation pages", func() {
		It("should be able to generate retrieve methods from templates", func() {
			output := templates.GenerateRetrieve("Entitlement", "")
			Expect(output).To(goldga.Match())
		})
	})
})
