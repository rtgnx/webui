{{ define "content" }}
<div class="rounded-box info-section">
  <h2> Members of {{.name}}</h2>
  <table>
    <tr>
      <th>Number</th>
      <th>ID</th>
      <th>Display Name</th>
    </tr>
    {{ range $index, $element := .members }}
    <tr>
      <td>{{$element.Number}}</td>
      <td><a href="/entity/info/{{$element.ID}}">{{$element.ID}}</a></td>
      <td>{{$element.Meta.DisplayName}}</td>
    </tr>
    {{ end }}
  </table>
</div>
{{ end }}
