{{ define "base" }}
<html>
  {{ template "head" . }}
  <body>
    <div class="row">
      <div>
        <h1 class="logo">NetAuth</h1>
      </div>
    </div>
    <div class="row">
      <div class="nav">
        {{ template "nav" . }}
      </div>
      <div class="main-content">
        {{ template "content" . }}
      </div>
    </div>
    {{ template "footer" . }}
  </body>
</html>
{{ end }}
