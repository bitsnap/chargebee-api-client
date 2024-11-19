package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee credit notes test", func() {
	When("accessing credit notes API", func() {
		It("should be able list credit notes", func() {
			notes, err := chargebee.ListCreditNotes(UseTestSite())
			Expect(err).To(Not(HaveOccurred()))
			Expect(notes).To(goldga.Match())
		})
	})
})
