package client_test

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee business entities code generation", func() {
	When("generating Chargebee API client", func() {
		It("should be able to generate business entities model", func() {
			model := client.GenerateBusinessEntitiesModel()
			Expect(model).To(goldga.Match())
		})

		It("should be able to generate business entities transfer model", func() {
			model := client.GenerateBusinessEntityTransferModel()
			Expect(model).To(goldga.Match())
		})
	})
})
