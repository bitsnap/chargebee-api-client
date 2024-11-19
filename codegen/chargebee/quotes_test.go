package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee quotes code generation", func() {
	When("generating Chargebee API client", func() {
		It("should be able to generate quotes models", func() {
			model := chargebee.GenerateQuotesModel()
			Expect(model).To(goldga.Match())
		})
	})
})
