package client

import (
	"embed"
	"strings"

	. "github.com/bitsnap/go-util"
)

//go:embed retrieve.go.tpl
var retrieveTemplateFS embed.FS

var retrieveTemplate = MustParseTemplate(retrieveTemplateFS, "retrieve.go.tpl")

func GenerateRetrieve(nameSingular string, url string) string {
	type RetrieveEndpoint struct {
		NameSingular string
		Url          string
		MaxRedirects uint
		Timeout      uint
	}

	output := strings.Builder{}
	err := retrieveTemplate.Execute(&output, RetrieveEndpoint{
		NameSingular: nameSingular,
		Url:          url,
		MaxRedirects: 5,
		Timeout:      10,
	})
	if err != nil {
		panic(err)
	}

	return output.String()
}
