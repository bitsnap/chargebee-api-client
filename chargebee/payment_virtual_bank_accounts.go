package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	"github.com/bitsnap/chargebee-api-client/chargebee/enums"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type VirtualBankAccount struct {
	Id               string            `json:"id" validate:"required"`
	CustomerId       string            `json:"customer_id" validate:"required"`
	Email            string            `json:"email" validate:"required"`
	Scheme           enums.SchemeEnum  `json:"scheme"`
	BankName         string            `json:"bank_name"`
	AccountNumber    string            `json:"account_number" validate:"required"`
	RoutingNumber    string            `json:"routing_number"`
	SwiftCode        string            `json:"swift_code" validate:"required"`
	Gateway          enums.GatewayEnum `json:"gateway" validate:"required"`
	GatewayAccountId string            `json:"gateway_account_id" validate:"required"`
	ResourceVersion  int64             `json:"resource_version"`
	UpdatedAt        uint64            `json:"updated_at"`
	CreatedAt        uint64            `json:"created_at" validate:"required"`
	ReferenceId      string            `json:"reference_id" validate:"required"`
	Deleted          bool              `json:"deleted" validate:"required"`
}

func ListVirtualBankAccountsPageSortBy(site string, offset string, sortBy *SortBy) ([]VirtualBankAccount, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/virtual_bank_accounts")
	if err != nil {
		return nil, "", err
	}

	content, err := getQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type VirtualBankAccountListItem struct {
		VirtualBankAccount VirtualBankAccount `json:"VirtualBankAccount"`
	}

	type VirtualBankAccountPage struct {
		List       []VirtualBankAccountListItem `json:"list"`
		NextOffset string                       `json:"next_offset,omitempty"`
	}

	entries := VirtualBankAccountPage{
		List: make([]VirtualBankAccountListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []VirtualBankAccount{}, "", nil
	}

	result := make([]VirtualBankAccount, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.VirtualBankAccount)
	}

	if len(entries.NextOffset) > 0 {
		return resultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListVirtualBankAccountsPage(site string, offset string) ([]VirtualBankAccount, string, error) {
	return ListVirtualBankAccountsPageSortBy(site, offset, nil)
}

func RetrieveVirtualBankAccountById(site string, id string) (*VirtualBankAccount, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/virtual_bank_accounts/" + id)
	if err != nil {
		return nil, err
	}

	content, err := getQuery(client, parsedUrl, "", nil)
	if err != nil {
		return nil, err
	}

	type VirtualBankAccountItem struct {
		VirtualBankAccount VirtualBankAccount `json:"VirtualBankAccount"`
	}

	var item VirtualBankAccountItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.VirtualBankAccount, nil
}
