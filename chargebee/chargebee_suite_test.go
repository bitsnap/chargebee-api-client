package chargebee_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestChargebeeSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chargebee Client Suite")
}
