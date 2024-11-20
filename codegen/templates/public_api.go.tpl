package chargebee

// THIS IS GENERATED CODE. DO NOT EDIT.

import (
    "time"

	"github.com/bitsnap/chargebee-api-client/chargebee/client"
	"github.com/bitsnap/chargebee-api-client/chargebee/generated"
)

// UseSite sets custom Chargebee site name overriding CHARGEBEE_SITE environment variable value
func UseSite(site string) {
	client.UseSite(site)
}

// UseToken sets custom Chargebee API token, overriding CHARGEBEE_API_TOKEN environment variable value
func UseToken(token string) {
	client.UseToken(token)
}
{{- range $method := .Methods }}

// {{ $method.FileName}}

func ListAll{{ $method.Name }}({{ if .WithID }}id string{{end}}) ([]chargebee.{{ $method.NameSingular }}, error) {
	return ListAllSite{{ $method.Name }}(client.Site(){{ if .WithID }}, id{{ end }})
}

func ListAllSite{{ $method.Name }}(site{{ if .WithID }}, id{{end}} string) ([]chargebee.{{ $method.NameSingular }}, error) {
	return chargebee.List{{ $method.Name }}(site{{ if .WithID }}, id{{ end }})
}

func List{{ $method.Name }}CreatedSince({{ if .WithID }}id string, {{end}}createdSince *time.Time) ([]chargebee.{{ $method.NameSingular }}, error) {
    return ListSite{{ $method.Name }}CreatedSince(client.Site(){{ if .WithID }}, id{{end}}, createdSince)
}

func ListSite{{ $method.Name }}CreatedSince(site{{ if .WithID }}, id{{end}} string, createdSince *time.Time) ([]chargebee.{{ $method.NameSingular }}, error) {
    return chargebee.List{{ $method.Name }}CreatedSince(site{{ if .WithID }}, id{{ end }}, createdSince)
}

func List{{ $method.Name }}UpdatedSince({{ if .WithID }}id string, {{end}}updatedSince *time.Time) ([]chargebee.{{ $method.NameSingular }}, error) {
    return ListSite{{ $method.Name }}UpdatedSince(client.Site(){{ if .WithID }}, id{{ end }}, updatedSince)
}

func ListSite{{ $method.Name }}UpdatedSince(site{{ if .WithID }}, id{{end}} string, updatedSince *time.Time) ([]chargebee.{{ $method.NameSingular }}, error) {
    return chargebee.List{{ $method.Name }}UpdatedSince(site{{ if .WithID }}, id{{ end }}, updatedSince)
}

func List{{ $method.Name }}({{ if .WithID }}id string{{end}}) ([]chargebee.{{ $method.NameSingular }}, string, error) {
	return ListSite{{ $method.Name }}(client.Site(){{ if .WithID }}, id{{end}})
}

func ListSite{{ $method.Name }}(site{{ if .WithID }}, id{{end}} string) ([]chargebee.{{ $method.NameSingular }}, string, error) {
	return chargebee.List{{ $method.Name }}Page(site{{ if .WithID }}, id{{end}}, "")
}

func List{{ $method.Name }}Page({{ if .WithID }}id, {{end}}offset string) ([]chargebee.{{ $method.NameSingular }}, string, error) {
	return ListSite{{ $method.Name }}Page(client.Site(){{ if .WithID }}, id{{end}}, offset)
}

func ListSite{{ $method.Name }}Page(site{{ if .WithID }}, id{{end}} string, offset string) ([]chargebee.{{ $method.NameSingular }}, string, error) {
	return chargebee.List{{ $method.Name }}Page(site{{ if .WithID }}, id{{end}}, offset)
}

func List{{ $method.Name }}PageSortBy({{ if .WithID }}id, {{end}}offset string, field string) ([]chargebee.{{ $method.NameSingular }}, string, error) {
	return ListSite{{ $method.Name }}PageSortBy(client.Site(){{ if .WithID }}, id{{end}}, offset, field)
}

func ListSite{{ $method.Name }}PageSortBy(site{{ if .WithID }}, id{{end}} string, offset string, field string) ([]chargebee.{{ $method.NameSingular }}, string, error) {
	sortBy := client.DefaultSortBy()
	sortBy.Name = field
	return chargebee.List{{ $method.Name }}PageSortBy(site{{ if .WithID }}, id{{end}}, offset, sortBy)
}{{ if .ByID }}

func Retrieve{{ $method.NameSingular }}ById(id string) (*chargebee.{{ $method.NameSingular }}, error) {
	return RetrieveSite{{ $method.NameSingular }}ById(client.Site(), id)
}

func RetrieveSite{{ $method.NameSingular }}ById(site string, id string) (*chargebee.{{ $method.NameSingular }}, error) {
	return chargebee.Retrieve{{ $method.NameSingular }}ById(site, id)
}{{ end }}{{ end }}
