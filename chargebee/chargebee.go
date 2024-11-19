package chargebee

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/goccy/go-json"
)

//go:generate go run ../../../cmd/bitsnap-accounting-chargebee-codegen/main.go

var defaultToken = ""

func ClientToken() string {
	if tok, ok := os.LookupEnv("CHARGEBEE_APIKEY"); ok {
		return tok
	}

	return defaultToken
}

func getQuery(client *fiber.Client, parsedUrl *url.URL, offset string, sortBy *SortBy) ([]byte, error) {
	q := parsedUrl.Query()
	if offset != "" {
		q.Add("offset", offset)
	}

	if sortBy != nil {
		if sortBy.Name != "" {
			if sortBy.Asc {
				q.Add("sort_by[asc]", sortBy.Name)
			} else {
				q.Add("sort_by[desc]", sortBy.Name)
			}
		}

		if sortBy.AfterDate != nil {
			if sortBy.UpdatedAfter {
				if sortBy.Asc {
					q.Add("sort_by[asc]", "updated_at")
				} else {
					q.Add("sort_by[desc]", "updated_at")
				}

				q.Add("updated_at[after]", strconv.Itoa(int(sortBy.AfterDate.Unix())))
				q.Add("updated_at[before]", strconv.Itoa(int(time.Now().Unix())))
			} else {
				if sortBy.Asc {
					q.Add("sort_by[asc]", "created_at")
				} else {
					q.Add("sort_by[desc]", "created_at")
				}

				q.Add("created_at[after]", strconv.Itoa(int(sortBy.AfterDate.Unix())))
				q.Add("created_at[before]", strconv.Itoa(int(time.Now().Unix())))
			}
		}
	}

	parsedUrl.RawQuery = q.Encode()

	statusCode, content, errs := client.Get(parsedUrl.String()).BasicAuth(ClientToken(), "").MaxRedirectsCount(5).Timeout(time.Second * 10).Bytes()
	if len(errs) > 0 || statusCode != 200 {
		return nil, fmt.Errorf("%v \nget failed, status code: %d, body: %s", parsedUrl.String(), statusCode, content)
	}

	return content, nil
}

func resultWithOffset[R any](result []R, offset, nextOffset string) ([]R, string, error) {
	nextOffsetVal := make([]string, 0, 10)

	// NOTE: some offsets are represented as Numbers while others as stringified JSON Arrays of string offsets :\
	err := json.Unmarshal([]byte(nextOffset), &nextOffsetVal)
	if err != nil {
		nextOffsetValNum := 1
		err := json.Unmarshal([]byte(nextOffset), &nextOffsetValNum)
		if err != nil {
			return nil, "", err
		}

		nextOffsetVal = append(nextOffsetVal, strconv.Itoa(nextOffsetValNum))
	}

	if offset != "" {
		for i := range nextOffsetVal {
			if nextOffsetVal[i] == offset {
				if i+1 != len(nextOffsetVal) {
					return result, nextOffsetVal[i+1], nil
				}

				break
			}
		}

		return result, "", nil
	}

	return result, nextOffsetVal[0], nil
}
