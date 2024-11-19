package client_test

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee payment transactions code generation", func() {
	When("generating Chargebee API client", func() {
		It("should be able to generate payment transactions models", func() {
			model := client.GenerateTransactionsModel()
			Expect(model).To(goldga.Match())
		})
	})
})