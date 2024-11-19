package scraper

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	. "github.com/bitsnap/go-util"
	"github.com/gocolly/colly/v2"
)

type Attribute struct {
	Name     string
	Tag      string
	Required bool
	Type     string
	MinChars int
	MaxChars int
}

func (a *Attribute) String() string {
	req := " required "
	if !a.Required {
		req = " "
	}

	return fmt.Sprintf("%v%v%v", a.Name, req, a.Type)
}

var validTypes = []string{
	"string",
	"long",
	"boolean",
	"double",
	"bigdecimal",
	"in cents",
	"integer",
	"jsonarray",
	"jsonobject",
	"enumerated string",
	"timestamp(UTC) in seconds",
}

var models = []string{
	// Customer
	"relationship",
	"parent_account_access",
	"child_account_access",
	"webhook",
	// Transactions
	"gateway_error_detail",
	// Subscriptions
	"referral_info",
	// Invoices
	"einvoice",
	"site_details_at_creation",
	"contract_term",
	"tax_origin",
	// Impacted subs
	"download",
	// Item Price
	"tax_detail",
	"accounting_detail",
	// Payment source
	"card",
	"bank_account",
	"cust_voucher_source",
	"amazon_payment",
	"upi",
	"paypal",
	"venmo",
	"klarna_pay_now",
	"quoted_contract_term",
	"component",
	// Gift
	"gifter",
	"gift_receiver",
	// Ramp
	"status_transition_reason",
	// Advanced Invoice Schedule
	"fixed_interval_schedule",
	"specific_dates_schedule",
	"webhooks",
	"statement_descriptor",
	"payment_method",
}

func ValidModel(name string) bool {
	if strings.HasSuffix(name, "address") {
		return true
	}

	return slices.Contains(models, name)
}

var enums = []string{
	// TODO: both model and an enum, needs disambiguation
	//"payment_method",
}

func ValidEnum(name string) bool {
	return slices.Contains(enums, name)
}

func NewAttribute(name, attrString string) (*Attribute, error) {
	isOptional := false
	maxChars := 0
	minChars := 0
	attributeType := "string"

	for _, component := range strings.Split(attrString, ",") {
		trimmedComponent := strings.TrimSpace(component)

		if strings.HasPrefix(trimmedComponent, "max chars") {
			maxCharsValue := strings.Replace(trimmedComponent, "max chars", "", -1)
			maxCharsValue = strings.Replace(maxCharsValue, "=", "", -1)
			maxCharsValue = strings.TrimSpace(maxCharsValue)
			maxCharsInt, err := strconv.Atoi(strings.TrimSpace(maxCharsValue))
			if err == nil {
				maxChars = maxCharsInt
			}

			continue
		}

		if strings.HasPrefix(trimmedComponent, "min chars") {
			minCharsValue := strings.Replace(trimmedComponent, "min chars", "", -1)
			minCharsValue = strings.Replace(minCharsValue, "=", "", -1)
			minCharsValue = strings.TrimSpace(minCharsValue)
			minCharsInt, err := strconv.Atoi(strings.TrimSpace(minCharsValue))
			if err == nil {
				minChars = minCharsInt
			}

			continue
		}

		if strings.HasPrefix(trimmedComponent, "optional") {
			isOptional = true

			continue
		}

		if strings.HasPrefix(trimmedComponent, "required") {
			isOptional = false

			continue
		}

		// TODO: handle defaults
		if strings.HasPrefix(trimmedComponent, "default=") {
			continue
		}

		// TODO: handle min
		if strings.HasPrefix(trimmedComponent, "min=") {
			continue
		}

		// TODO: handle max
		if strings.HasPrefix(trimmedComponent, "max=") {
			continue
		}

		if strings.HasSuffix(trimmedComponent, "filter") {
			trimmedComponent = strings.TrimSuffix(trimmedComponent, " filter")
		}

		if slices.Contains(validTypes, trimmedComponent) {
			attributeType = formatType(trimmedComponent)

			continue
		}

		if strings.HasPrefix(trimmedComponent, "list of") {
			attributeType = formatType(trimmedComponent)

			continue
		}

		if ValidModel(trimmedComponent) {
			attributeType = "models." + SnakeToCamel(trimmedComponent, true)

			continue
		}

		if ValidEnum(trimmedComponent) {
			attributeType = "enums." + SnakeToCamel(trimmedComponent, true) + "Enum"

			continue
		}

		return nil, fmt.Errorf("unknown attribute type: %s", trimmedComponent)
	}

	return &Attribute{
		Name:     SnakeToCamel(name, true),
		Tag:      formatTag(name, !isOptional),
		Required: !isOptional,
		Type:     formatEnumName(name, attributeType),
		MinChars: minChars,
		MaxChars: maxChars,
	}, nil
}

func EnsureParentLevel(root *goquery.Selection, elements *goquery.Selection, level int) (ok bool) {
	el := elements.First()
	for i := 0; i < level; i++ {
		el = el.Parent()
	}

	return root.Index() == el.Index() && root.Children().Length() == el.Children().Length()
}

func ScrapeAttributes(header string, page string) ([]*Attribute, error) {
	collector := colly.NewCollector(
		colly.AllowedDomains("apidocs.chargebee.com", "chargebee.com"),
		colly.CacheDir("./.cache"),
	)

	attributes := make([]*Attribute, 0, 10)

	collector.OnHTML("div.page-header h2", func(headerElement *colly.HTMLElement) {
		headerElementText := strings.TrimSpace(strings.Replace(headerElement.Text, "¶", "", -1))
		if headerElementText == header {
			headerElement.DOM.ParentsFiltered("div.page-header").NextAllFiltered("div.cb-list-group").Each(func(_ int, list *goquery.Selection) {
				toString := func(_ int, i *goquery.Selection) string {
					return strings.TrimSpace(i.Text())
				}

				filterLevel := func(_ int, i *goquery.Selection) bool {
					return EnsureParentLevel(list, i, 4)
				}

				attributeNames := list.Find("div.cb-list-item > samp").FilterFunction(filterLevel).Map(toString)
				attributeTypes := list.Find("div.cb-list-desc > dfn").FilterFunction(filterLevel).Map(toString)

				for i, attributeString := range attributeTypes {
					attr, err := NewAttribute(attributeNames[i], attributeString)
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
		Logger().Errorw(response.Request.URL.String(), "error", err)
	})

	err := collector.Visit(page)
	if err != nil {
		return nil, err
	}

	return attributes, nil
}
