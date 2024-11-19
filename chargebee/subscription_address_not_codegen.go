package chargebee

// TODO: custom code

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/bitsnap/chargebee-api-client/chargebee/models"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type subscriptionAddressLabel string

const (
	subscriptionShippingAddressLabel subscriptionAddressLabel = "shipping_address"
	subscriptionBillingAddressLabel  subscriptionAddressLabel = "billing_address"
)

func SubscriptionShippingAddress(site, subscriptionId string) (*Address, error) {
	return subscriptionAddress(site, subscriptionId, subscriptionShippingAddressLabel)
}

func SubscriptionBillingAddress(site, subscriptionId string) (*Address, error) {
	return subscriptionAddress(site, subscriptionId, subscriptionBillingAddressLabel)
}

func subscriptionAddress(site, subscriptionId string, addressLabel subscriptionAddressLabel) (*Address, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/addresses")
	if err != nil {
		return nil, err
	}

	q := parsedUrl.Query()
	q.Add("subscription_id", subscriptionId)
	q.Add("label", string(addressLabel))
	parsedUrl.RawQuery = q.Encode()

	statusCode, content, errs := client.Get(parsedUrl.String()).BasicAuth(ClientToken(), "").MaxRedirectsCount(5).Timeout(time.Second * 10).Bytes()
	if statusCode == 404 && strings.Contains(string(content), string(addressLabel)+" not found") {
		return nil, nil
	}

	if len(errs) > 0 || statusCode != 200 {
		return nil, fmt.Errorf("%v \nget failed, status code: %d, body: %s", parsedUrl.String(), statusCode, content)
	}

	type AddressListItem struct {
		Address Address `json:"Address"`
	}

	entry := &AddressListItem{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Address, nil
}

type Address struct {
	Label          string `json:"label" validate:"required"`
	SubscriptionId string `json:"subscription_id"`
	models.Address

	// TODO: left for backwards compat but should be safe to delete
	Addr          string `json:"addr"`
	ExtendedAddr  string `json:"extended_addr"`
	ExtendedAddr2 string `json:"extended_addr2"`
}
