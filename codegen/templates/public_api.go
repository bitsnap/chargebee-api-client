package templates

import (
	"embed"
	"strings"

	. "github.com/bitsnap/go-util"
)

//go:embed public_api.go.tpl
var publicAPITemplateFS embed.FS

var publicAPITemplate = MustParseTemplate(publicAPITemplateFS, "public_api.go.tpl")

func GeneratePublicAPIContent() map[string][]func() string {
	output := strings.Builder{}
	err := publicAPITemplate.Execute(&output, convenience)
	if err != nil {
		panic(err)
	}

	return map[string][]func() string{
		"chargebee.go": {
			output.String,
		},
	}
}
