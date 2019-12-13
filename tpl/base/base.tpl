{{ define "base" }}
<html>
  {{ template "head" . }}
  <body>
    <div id="top-bar">
      <div class="right">
        {{ if .token }}
        <a class="button" href="/auth/logout">Log Out ({{.token.EntityID}})</a>
        {{ end }}
      </div>
    </div>
    <div id="container">
      <div class="row">
        <div>
          {{ block "nav" . }}
          {{ end }}
        </div>
        <div class="main-content">
          {{ block "content" .data }}
          <div class="rounded-box">
            This is where some page content would be if it loaded.
          </div>
          {{ end }}
        </div>
      </div>
      {{ block "footer" . }}
      {{ end }}
    </div>
  </body>
</html>
{{ end }}
