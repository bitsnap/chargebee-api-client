package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
	. "github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/bitsnap/chargebee-api-client/chargebee/enums"
	"github.com/bitsnap/chargebee-api-client/chargebee/models"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

type Feature struct {
    Id string `json:"id" validate:"required"`
    Name string `json:"name" validate:"required"`
    Description string `json:"description"`
    Status enums.StatusEnum `json:"status"`
    Type enums.TypeEnum `json:"type"`
    Unit string `json:"unit"`
    ResourceVersion int64 `json:"resource_version"`
    UpdatedAt uint64 `json:"updated_at"`
    CreatedAt uint64 `json:"created_at" validate:"required"`
    Levels []models.Level `json:"levels"`
}

func ListFeaturesPageSortBy(site string, offset string, sortBy *SortBy) ([]Feature, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/features")
	if err != nil {
		return nil, "", err
	}
	
    content, err := GetQuery(client, parsedUrl, offset, sortBy)
    if err != nil {
        return nil, "", err
    }
    		
	type FeatureListItem struct {
		Feature Feature `json:"Feature"`
	}

    type FeaturePage struct {
        List       []FeatureListItem `json:"list"`
        NextOffset string `json:"next_offset,omitempty"`
    }

	entries := FeaturePage{
		List:       make([]FeatureListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}
	
	if len(entries.List) == 0 {
        return []Feature{}, "", nil
    }
	
	result := make([]Feature, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.Feature)
	}

    if len(entries.NextOffset) > 0 {
        return ResultWithOffset(result, offset, entries.NextOffset)
    }
	
	return result, "", nil
}

func ListFeaturesPage(site string, offset string) ([]Feature, string, error) {
	return ListFeaturesPageSortBy(site, offset, nil)
}

func RetrieveFeatureById(site string, id string) (*Feature, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "chargebee.com/api/v2/features/" + id)
	if err != nil {
		return nil, err
	}
	
    content, err := GetQuery(client, parsedUrl, "", nil)
    if err != nil {
        return nil, err
    }
    	
	type FeatureItem struct {
		Feature Feature `json:"Feature"`
	}

    var item FeatureItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.Feature, nil
}