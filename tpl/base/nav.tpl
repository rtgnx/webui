{{ define "nav" }}
<div class="rounded-box nav">
  <nav>
    <ul>
      <li class="nav-section">Entity</li>
      <li>
        <ul>
          <a href="/entity/search"><li class="nav-page">Search</li></a>
          <a href="/entity/create"><li class="nav-page">Create</li></a>
        </ul>
      </li>
      <li class="nav-section">Group</li>
      <li>
        <ul>
          <a href="/group/search"><li class="nav-page">Search</li></a>
        </ul>
      </li>
      <li class="nav-section">System</li>
      <li>
        <ul>
          <a href="/system/status"><li class="nav-page">Status</li></a>
        </ul>
      </li>
      <li class="nav-section">Meta</li>
      <li>
        <ul>
          <a href="/meta/about"><li class="nav-page">About</li></a>
        </ul>
      </li>
    </ul>
  </nav>
</div>
{{ end }}
