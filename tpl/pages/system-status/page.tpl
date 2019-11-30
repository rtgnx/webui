{{ define "content" }}
{{ if .SystemOK }}
<h1>NetAuth is operating normally.</h1>
{{ else }}
<h1>There are problems!</h1>

<p>The first system reporting a problem is {{.FirstFailure.Name}}</p>
{{ end }}
<br />

<table>
  <tr>
    <th>Status</th>
    <th>System</th>
    <th>Message</th>
  </tr>
  {{ range .SubSystems }}
  <tr>
    {{ if .OK }}
    <td class="status-good">PASS</td>
    {{ else }}
    <td class="status-bad">FAIL</td>
    {{ end }}
    <td>{{.Name}}</td>
    <td>{{.FaultMessage}}</td>
  </tr>
  {{ end }}
</table>
{{ end }}
