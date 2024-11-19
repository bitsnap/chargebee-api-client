package client

import (
	"embed"
	"strings"

	. "github.com/bitsnap/go-util"
)

//go:embed list_children.go.tpl
var listChildrenTemplateFS embed.FS
var listChildrenTemplate = MustParseTemplate(listChildrenTemplateFS, "list_children.go.tpl")

func GenerateListChildren(name, nameSingular, url, urlSuffix string) string {
	return GenerateListChildrenWithId(name, nameSingular, url, urlSuffix, true, "")
}

func GenerateListChildrenWithId(name string, nameSingular string, url string, urlSuffix string, addIdURLSuffix bool, renameId string) string {
	type ListChildrenEndpoint struct {
		Name         string
		NameSingular string
		Url          string
		UrlSuffix    string
		MaxRedirects uint
		Timeout      uint
		AddSuffix    bool
		RenameId     string
	}

	if renameId == "" {
		renameId = "id"
	}

	output := strings.Builder{}
	err := listChildrenTemplate.Execute(&output, ListChildrenEndpoint{
		Name:         name,
		NameSingular: nameSingular,
		Url:          url,
		UrlSuffix:    urlSuffix,
		MaxRedirects: 5,
		Timeout:      10,
		AddSuffix:    addIdURLSuffix,
		RenameId:     renameId,
	})
	if err != nil {
		panic(err)
	}

	return output.String()
}
