package client_test

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee items code generation", func() {
	When("generating Chargebee API client", func() {
		It("should be able to generate items models", func() {
			model := client.GenerateItemsModel()
			Expect(model).To(goldga.Match())
		})
	})
})
