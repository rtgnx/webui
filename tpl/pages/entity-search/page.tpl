{{ define "content" }}
{{ if not .query }}
<p>The search syntax used here
  is <a href="http://blevesearch.com/docs/Query-String-Query/">BleveSearch</a>.
  You can use expressions like <code>Meta.Shell:/bin/bash</code>
  or <code>ID:partial-match</code>.</p>

<br />

<form>
  <fieldset>
    <legend>Entity Search Query</legend>

    <input type="search"
           id="query"
           name="query"
           placeholder="ID:*"
           required="yes"
           autofocus="yes" />
    <input type="submit" />
  </fieldset>
</form>
{{ else }}
{{ if .result }}
<table>
  <tr>
    <th>Number</th>
    <th>ID</th>
    <th>Display Name</th>
  </tr>
  {{ range $index, $element := .result }}
  <tr>
    <td>{{ $element.Number }}</td>
    <td><a href="/entity/info/{{ $element.ID }}">{{ $element.ID }}</a></td>
    <td>{{ $element.Meta.DisplayName }}</td>
  </tr>
  {{ end }}
</table>
{{ else }}
<p>Your search returned no results.</p>
{{ end }}
{{ end }}
{{ end }}
