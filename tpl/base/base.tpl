{{ define "base" }}
<html>
  {{ template "head" . }}
  <body>
    <div>
      {{ template "content" . }}
    </div>
    {{ template "footer" . }}
  </body>
</html>
{{ end }}
