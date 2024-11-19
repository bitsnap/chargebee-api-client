type {{ .Name }} struct {
{{- range $attr := .Attributes }}
    {{ $attr.Name }} {{ $attr.Type }} {{ $attr.Tag }}{{ end }}
}
