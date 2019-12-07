{{ define "content" }}
<div class="rounded-box info-section">
  {{ if not .query }}
  <p>The search syntax used here
    is <a href="http://blevesearch.com/docs/Query-String-Query/">BleveSearch</a>.
    You can use expressions like <code>Meta.Shell:/bin/bash</code>
    or <code>ID:partial-match</code>.</p>

  <br />

  <form>
    <fieldset>
      <legend>Group Search Query</legend>

      <input type="search"
             id="query"
             name="query"
             placeholder="Name:*"
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
      <td><a href="/group/info/{{ $element.Name }}">{{ $element.Name }}</a></td>
      <td>{{ $element.DisplayName }}</td>
    </tr>
    {{ end }}
  </table>
  {{ else }}
  <p>Your search returned no results.</p>
  {{ end }}
  {{ end }}
</div>
{{ end }}
