{{ define "nav" }}
<div></div>
{{ end }}

{{ define "content" }}
<div class="flex-center">
  <div class="rounded-box">
    <h1>Authentication Required:</h1>
    <form method="POST">
      <input type="text" name="id" placeholder="ID" /> <br />
      <input type="password" name="secret" placeholder="secret" /> <br />
      <div class="center">
        <button>Login</button>
      </div>
    </form>
  </div>
</div>
{{ end }}

{{ define "footer" }}
<div></div>
{{ end }}
