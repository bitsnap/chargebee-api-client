package client_test

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee entitlements code generation", func() {
	When("generating Chargebee API client", func() {
		It("should be able to generate entitlements models", func() {
			model := client.GenerateEntitlementsModel()
			Expect(model).To(goldga.Match())
		})
	})
})
