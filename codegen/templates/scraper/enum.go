package scraper

import (
	"github.com/PuerkitoBio/goquery"
	. "github.com/bitsnap/go-util"
	"github.com/gocolly/colly/v2"
)

func ScrapeEnum(abbr string, page string) (map[string]string, error) {
	collector := colly.NewCollector(
		colly.AllowedDomains("apidocs.chargebee.com", "chargebee.com"),
		colly.CacheDir("./.cache"),
	)

	values := make(map[string]string)

	collector.OnHTML("div.cb-list-group", func(lists *colly.HTMLElement) {
		filterAbbr := func(_ int, elem *goquery.Selection) bool {
			return elem.Find("div.cb-list-item samp").First().Text() == abbr
		}

		elem := lists.DOM.Find("div.cb-list-action").FilterFunction(filterAbbr).First()
		snakeCaseValues := elem.Find("samp.enum").Map(func(i int, v *goquery.Selection) string {
			return v.Text()
		})

		for _, v := range snakeCaseValues {
			values[SnakeToCamel(v, true)] = v
		}
	})

	collector.OnError(func(response *colly.Response, err error) {
		Logger().Errorw(response.Request.URL.String(), "error", err)
	})

	err := collector.Visit(page)
	if err != nil {
		return nil, err
	}

	return values, nil
}
