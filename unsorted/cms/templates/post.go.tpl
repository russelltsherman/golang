{{ define "post" }}
  <h1>{{ .Title }}</h1>
  <p>{{ .Content }}</p>
  {{ if .Comments }}
    {{ range .Comments }}
      {{ template "comment" . }}
    {{ end }}
  {{ end }}
{{ end }}
