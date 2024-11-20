package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

type ContractTerm struct {
}

func ListContractTermsPageSortBy(site string, id string, offset string, sortBy *SortBy) ([]ContractTerm, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/subscriptions/" + id + "/contract_terms")
	if err != nil {
		return nil, "", err
	}
		
    content, err := GetQuery(client, parsedUrl, offset, sortBy)
    if err != nil {
        return nil, "", err
    }
    	
	type ContractTermListItem struct {
		ContractTerm ContractTerm `json:"ContractTerm"`
	}

    type ContractTermPage struct {
        List       []ContractTermListItem `json:"list"`
        NextOffset string `json:"next_offset,omitempty"`
    }

	entries := ContractTermPage{
		List:       make([]ContractTermListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}
	
	if len(entries.List) == 0 {
        return []ContractTerm{}, "", nil
    }
	
	result := make([]ContractTerm, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.ContractTerm)
	}

    if len(entries.NextOffset) > 0 {
        return ResultWithOffset(result, offset, entries.NextOffset)
    }
	
	return result, "", nil
}

func ListContractTermsPage(site string, id string, offset string) ([]ContractTerm, string, error) {
	return ListContractTermsPageSortBy(site, id, offset, nil)
}