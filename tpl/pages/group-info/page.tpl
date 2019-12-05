{{ define "content" }}
{{ with .group }}
<div class="rounded-box info-section">
  <h2>Base Information</h2>
  <table>
    <tr>
      <td>Name</td>
      <td>{{.Name}}</td>
    </tr>
    <tr>
      <td>Display Name</td>
      <td>{{.DisplayName}}</td>
    </tr>
    <tr>
      <td>Number</td>
      <td>{{.Number}}</td>
    </tr>
  </table>

  {{ if .ManagedBy }}
  <br />
  <p>This group is managed by <a href="/group/info/{{.ManagedBy}}">{{.ManagedBy}}</a></p>
  {{ end }}
</div>
{{ if .Capabilities }}
<div class="rounded-box info-section">
  <h2>Capabilities</h2>
  <ul>
    {{ range $index, $element := .Capabilities }}
    <li>{{ $element }}</li>
    {{ end }}
  </ul>
</div>
{{ end }}
{{ if .Expansions }}
<div class="rounded-box info-section">
  <h2>Rules</h2>
  <ul>
    {{ range $index, $element := .Expansions }}
    <li>{{ $element }}</li>
    {{ end }}
  </ul>
</div>
{{ end }}
{{ end }}
{{ if .manages }}
<div class="rounded-box info-section">
  <h2>Groups managed by this group</h2>
  <table>
    <tr>
      <td>Number</td>
      <td>Name</td>
      <td>Display Name</td>
    </tr>
    {{ range $index, $element := .manages }}
    <tr>
      <td>{{$element.Number}}</td>
      <td><a href="/group/info/{{$element.Name}}">{{$element.Name}}</a></td>
      <td>{{$element.DisplayName}}</td>
    </tr>
    {{ end }}
  </table>
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
