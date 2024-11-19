package integration_tests

import (
	"os"
	"testing"

	chargebee "github.com/bitsnap/chargebee-api-client/chargebee/generated"
	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tommy351/goldga"
)

func TestChargebeeSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chargebee Client Suite")
}

func init() {
	conf := spew.NewDefaultConfig()
	conf.SortKeys = true
	conf.DisableCapacities = true
	conf.DisablePointerAddresses = true
	conf.DisablePointerMethods = true

	goldga.DefaultSerializer = &goldga.DumpSerializer{
		Config: conf,
	}
}

func UseTestSite() string {
	if site, ok := os.LookupEnv("CHARGEBEE_TEST_SITE"); ok {
		return site
	}

	return "bitsnap-test"
}

func GetFirst[T any](f func(site string) ([]T, error)) (*T, error) {
	items, err := f(UseTestSite())
	if len(items) == 0 {
		return nil, err
	}

	return &items[0], err
}

func GetFirstCustomer() (*chargebee.Customer, error) {
	return GetFirst(chargebee.ListCustomers)
}

func GetFirstInvoice() (*chargebee.Invoice, error) {
	return GetFirst(chargebee.ListInvoices)
}

func GetFirstQuote() (*chargebee.Quote, error) {
	return GetFirst(chargebee.ListQuotes)
}

func GetFirstSubscription() (*chargebee.Subscription, error) {
	return GetFirst(chargebee.ListSubscriptions)
}

func GetFirstItem() (*chargebee.Item, error) {
	return GetFirst(chargebee.ListItems)
}
