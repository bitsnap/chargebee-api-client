
{{- if .Name }}
{{ if not (eq .Name .Type) }}
type {{ .Name }} {{ .Type }}
{{- end }}
{{- end }}
{{- if .Values }}

const (
    {{- range $k, $v := .Values }}
        {{ $.Prefix }}{{ $k }} {{ $.Name }} {{ if $.IsInt }}= {{ $v }}{{ else }}= "{{ $v }}"{{ end }}
    {{- end }}
){{- end }}