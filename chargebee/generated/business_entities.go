package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	"net/url"

	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/bitsnap/chargebee-api-client/chargebee/generated/enums"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type BusinessEntity struct {
	Id              string           `json:"id" validate:"required"`
	Name            string           `json:"name" validate:"required"`
	Status          enums.StatusEnum `json:"status" validate:"required"`
	Deleted         bool             `json:"deleted" validate:"required"`
	CreatedAt       uint64           `json:"created_at" validate:"required"`
	ResourceVersion int64            `json:"resource_version"`
	UpdatedAt       uint64           `json:"updated_at"`
}

type BusinessEntityTransfer struct {
	Id                          string                 `json:"id" validate:"required"`
	ResourceType                enums.ResourceTypeEnum `json:"resource_type" validate:"required"`
	ResourceId                  string                 `json:"resource_id" validate:"required"`
	ActiveResourceId            string                 `json:"active_resource_id" validate:"required"`
	DestinationBusinessEntityId string                 `json:"destination_business_entity_id" validate:"required"`
	SourceBusinessEntityId      string                 `json:"source_business_entity_id" validate:"required"`
	ReasonCode                  enums.ReasonCodeEnum   `json:"reason_code" validate:"required"`
	CreatedAt                   uint64                 `json:"created_at" validate:"required"`
}

func ListBusinessEntityTransfersPageSortBy(site string, offset string, sortBy *SortBy) ([]BusinessEntityTransfer, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/business_entities/transfers")
	if err != nil {
		return nil, "", err
	}

	content, err := GetQuery(client, parsedUrl, offset, sortBy)
	if err != nil {
		return nil, "", err
	}

	type BusinessEntityTransferListItem struct {
		BusinessEntityTransfer BusinessEntityTransfer `json:"BusinessEntityTransfer"`
	}

	type BusinessEntityTransferPage struct {
		List       []BusinessEntityTransferListItem `json:"list"`
		NextOffset string                           `json:"next_offset,omitempty"`
	}

	entries := BusinessEntityTransferPage{
		List: make([]BusinessEntityTransferListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}

	if len(entries.List) == 0 {
		return []BusinessEntityTransfer{}, "", nil
	}

	result := make([]BusinessEntityTransfer, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.BusinessEntityTransfer)
	}

	if len(entries.NextOffset) > 0 {
		return ResultWithOffset(result, offset, entries.NextOffset)
	}

	return result, "", nil
}

func ListBusinessEntityTransfersPage(site string, offset string) ([]BusinessEntityTransfer, string, error) {
	return ListBusinessEntityTransfersPageSortBy(site, offset, nil)
}
