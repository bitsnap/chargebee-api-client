{{- range $method := .Methods }}

func List{{ $method.Name }}(site{{ if .WithID }}, id{{end}} string) ([]{{ $method.NameSingular }}, error) {
	return {{ if .WithID }}ListAllOfId{{ else }}ListAll{{ end }}(site,{{ if .WithID }} id,{{ end }} DefaultSortBy(), List{{ $method.Name }}PageSortBy)
}

func List{{ $method.Name }}CreatedSince(site{{ if .WithID }}, id{{end}} string, createdSince *time.Time) ([]{{ $method.NameSingular }}, error) {
	return {{ if .WithID }}ListAllOfId{{ else }}ListAll{{ end }}(site,{{ if .WithID }} id,{{ end }} &SortBy{
		Name:      "created_at",
		Asc:       true,
		AfterDate: createdSince,
	}, List{{ $method.Name }}PageSortBy)
}

func List{{ $method.Name }}UpdatedSince(site{{ if .WithID }}, id{{end}} string, updatedSince *time.Time) ([]{{ $method.NameSingular }}, error) {
	return {{ if .WithID }}ListAllOfId{{ else }}ListAll{{ end }}(site,{{ if .WithID }} id,{{ end }} &SortBy{
		Name:         "updated_at",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, List{{ $method.Name }}PageSortBy)
}
{{end}}
