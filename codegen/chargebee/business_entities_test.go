package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee business entities code generation", func() {
	When("generating Chargebee API client", func() {
		It("should be able to generate business entities model", func() {
			model := chargebee.GenerateBusinessEntitiesModel()
			Expect(model).To(goldga.Match())
		})

		It("should be able to generate business entities transfer model", func() {
			model := chargebee.GenerateBusinessEntityTransferModel()
			Expect(model).To(goldga.Match())
		})
	})
})