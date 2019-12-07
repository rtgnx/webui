{{ define "base" }}
<html>
  {{ template "head" . }}
  <body>
    <div id="container">
      <div class="row">
        <div class="rounded-box nav">
          {{ template "nav" . }}
        </div>
        <div class="main-content">
          {{ template "content" . }}
        </div>
      </div>
      {{ template "footer" . }}
    </div>
  </body>
</html>
{{ end }}
