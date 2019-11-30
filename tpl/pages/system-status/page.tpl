{{ define "content" }}
The system is on fire.
<br />
<ul>
  {{ range .SubSystems }}
  <li>{{.}}</li>
  {{ end }}
</ul>
{{ end }}
