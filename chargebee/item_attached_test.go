package chargebee_test

import (
	"github.com/bitsnap/chargebee-api-client/chargebee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

var _ = Describe("Chargebee items attached test", func() {
	When("accessing items attached API", func() {
		It("should be able list items attached", func() {
			item, err := GetFirstItem()
			Expect(err).To(Not(HaveOccurred()))

			if item != nil {
				testItemId := item.Id

				itemsAttached, err := chargebee.ListItemsAttached(UseTestSite(), testItemId)
				Expect(err).To(Not(HaveOccurred()))
				Expect(itemsAttached).To(goldga.Match())
			}
		})
	})
})
