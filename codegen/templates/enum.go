package templates

import (
	"embed"
	"github.com/bitsnap/chargebee-api-client/codegen/templates/scraper"
	"strings"

	. "github.com/bitsnap/go-util"
)

//go:embed enum.go.tpl
var enumTemplateFS embed.FS

var enumTemplate = MustParseTemplate(enumTemplateFS, "enum.go.tpl")

func GenerateStringEnum(name, abbr, page string) string {
	return GenerateEnum(name, abbr, "", "", page)
}

func GeneratePreffixedStringEnum(name, prefix, abbr, page string) string {
	return GenerateEnum(name, abbr, prefix, "", page)
}

func GenerateAliasEnum(name, t string) string {
	return GenerateEnum(name, "", "", t, "")
}

func GenerateEnum(name, abbr, prefix, t, page string) string {
	var (
		values = make(map[string]string)
		err    error
	)

	if abbr != "" && page != "" {
		values, err = scraper.ScrapeEnum(abbr, page)
		if err != nil {
			Logger().Fatalln(err)
		}
	}

	if t == "" {
		t = "string"
	}

	type EnumValues struct {
		Name   string
		Type   string
		IsInt  bool
		Prefix string
		Values map[string]string
	}

	output := strings.Builder{}
	err = enumTemplate.Execute(&output, EnumValues{
		Name:   name,
		Type:   t,
		IsInt:  strings.HasSuffix(t, "int"),
		Prefix: prefix,
		Values: values,
	})

	if err != nil {
		panic(err)
	}

	return output.String()
}
