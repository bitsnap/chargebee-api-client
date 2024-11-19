package templates_test

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("List client endpoint code generation", func() {
	When("generating list endpoints client", func() {
		It("should be able to generate list methods", func() {
			output := templates.GenerateList("Entitlements", "Entitlement", "")
			Expect(output).To(goldga.Match())
		})

		It("should be able to generate list children methods", func() {
			output := templates.GenerateListChildren("Entitlements", "Entitlement", "", "children")
			Expect(output).To(goldga.Match())
		})
	})
})
