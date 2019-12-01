{{ define "content" }}
{{ with .entity }}
<div class="rounded-box info-section">
  <h2>Base Information</h2>
  <table>
    <tr>
      <td>ID</td>
      <td>{{.ID}}</td>
    </tr>
    <tr>
      <td>Number</td>
      <td>{{.Number}}</td>
    </tr>
  </table>
</div>
<div class="rounded-box info-section">
  <h2>Additional Information</h2>
  {{ with .Meta }}
  <table>
    <tr>
      <td>Primary Group</td>
      <td>{{.PrimaryGroup}}</td>
    </tr>
    <tr>
      <td>GECOS</td>
      <td>{{.GECOS}}</td>
    </tr>
    <tr>
      <td>Display Name</td>
      <td>{{.DisplayName}}</td>
    </tr>
    <tr>
      <td>Home Directory</td>
      <td>{{.Home}}</td>
    </tr>
  </table>
  {{ end }}
</div>
{{ end }}
{{ if .keys }}
<div class="rounded-box info-section">
  <h2>Keys</h2>
  {{ range $type, $keys := .keys }}
  {{$type}}
  <ul>
    {{ range $index, $element := $keys }}
    <li class="ssh-key">{{ $element }}</li>
    {{ end }}
  </ul>
  {{ end }}
</div>
{{ end }}
{{ if .kv }}
<div class="rounded-box info-section">
  <h2>Untyped Key/Value</h2>
  <table>
    <tr>
      <th>Key</th>
      <th>Value</th>
    </tr>
    {{ range $key, $values := .kv }}
    <tr>
      <td>{{ $key }}</td>
      <td>{{ $values }}</td>
    </tr>
    {{ end}}
  </table>
</div>
{{ end }}
{{ end }}
