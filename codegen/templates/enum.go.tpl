
type {{ .Name }} {{ .Type }}

{{- if .Values }}

const (
    {{- $first := true }}
    {{- range $k, $v := .Values }}
        {{- if $first }}
        {{ $.Prefix }}{{ $k }} {{ $.Name }} {{ if $.IsInt }}= {{ $v }}{{ else }}= "{{ $v }}"{{ end }}
        {{- $first = false }}
        {{- else }}
        {{ $.Prefix }}{{ $k }} {{ if $.IsInt }}= {{ $v }}{{ else }}= "{{ $v }}"{{ end }}
        {{- end }}
    {{- end }}
)
{{- end }}