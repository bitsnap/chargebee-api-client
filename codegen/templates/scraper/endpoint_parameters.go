package scraper

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	. "github.com/bitsnap/go-util"
	"github.com/gocolly/colly/v2"
)

var replaceBrackets = regexp.MustCompile(`\[.*]`)

// TODO: integrate filtering for the API client ???

func ScrapeEndpointParameters(header string, page string) ([]*Attribute, error) {
	collector := colly.NewCollector()
	attributes := make([]*Attribute, 0, 10)

	collector.OnHTML("div.page-header h2", func(headerElement *colly.HTMLElement) {
		headerElementText := strings.TrimSpace(strings.Replace(headerElement.Text, "¶", "", -1))
		if headerElementText == header {
			headerElement.DOM.ParentsFiltered("div.page-header").NextAllFiltered("div.cb-list-group").Each(func(_ int, list *goquery.Selection) {
				attributeNames := make([]string, 0, 10)
				list.Find("div.cb-list-action").Each(func(_ int, e *goquery.Selection) {
					name := e.Find("div.cb-list-action.cb-list > div > div.cb-list-item > samp").First().Text()
					name = strings.TrimSpace(replaceBrackets.ReplaceAllString(name, ""))
					attributeNames = append(attributeNames, name)
				})

				attributeTypes := make([]string, 0, 10)
				list.Find("div.cb-list-action").Each(func(_ int, e *goquery.Selection) {
					name := e.Find("div.cb-list-action.cb-list > div > div.cb-list-desc > dfn").First().Text()
					name = strings.TrimSpace(replaceBrackets.ReplaceAllString(name, ""))
					attributeTypes = append(attributeTypes, name)
				})

				if len(attributeTypes) != len(attributeNames) {
					Logger().Fatalw("attribute names len != attribute types len", "page", page, "attributeTypes", len(attributeTypes), "attributeNames", len(attributeNames))
				}

				for i, attributeName := range attributeNames {
					attr, err := NewAttribute(attributeName, attributeTypes[i])
					if err != nil {
						Logger().Errorw(err.Error(), "header", header, "attr", attributeNames[i])
					} else {
						attributes = append(attributes, attr)
					}
				}
			})
		}
	})

	collector.OnError(func(response *colly.Response, err error) {
		Logger().Errorw(response.Request.URL.String(), err)
	})

	err := collector.Visit(page)
	if err != nil {
		return nil, err
	}

	return attributes, nil
}
