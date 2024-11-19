package client

import (
	"embed"
	"strings"

	"github.com/bitsnap/chargebee-api-client/codegen/scraper"
	. "github.com/bitsnap/go-util"
)

//go:embed model.go.tpl
var modelTemplateFS embed.FS

var modelTemplate = MustParseTemplate(modelTemplateFS, "model.go.tpl")

type Model struct {
	Name       string
	Attributes []*scraper.Attribute
}

func GenerateModel(name, header, page string) string {
	attributes, err := scraper.ScrapeAttributes(header, page)
	if err != nil {
		Logger().Fatalln(err)
	}

	output := strings.Builder{}
	err = modelTemplate.Execute(&output, Model{name, attributes})
	if err != nil {
		panic(err)
	}

	return output.String()
}
