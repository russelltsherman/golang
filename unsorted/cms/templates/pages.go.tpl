{{ define "pages" }}
<!DOCTYPE html>
<html>
<head>
  <title>Latest Pages</title>
</head>
<body>
  <h1>Latest Pages</h1>
  {{ range . }}
    <h2>
      <a href="/page/{{ .ID }}">{{ .Title }}</a>
    </h2>
    <p>{{ .Content }}</p>
  {{ end }}
</body>
</html>
{{ end }}
