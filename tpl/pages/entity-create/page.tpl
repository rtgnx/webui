{{ define "content" }}
<div class="rounded-box">
  <h1>Create New Entity</h1>
  <br />
  <p>On this page you can create a new entity.</p>
  <form method="post">
    <fieldset>
      <legend>Set Entity ID</legend>
      <p>Specify the ID for the new entity</p>
      <p>This is the ID that an entity will need to use to identify
        itself.  While NetAuth will accept any unicode string here,
        there are many reasons you should pick some an ID with some
        formatting rules.</p>
      <p>A good recommendation if you don't have your own rules is to
        adopt the rules for names on POSIX systems:</p>
      <ul>
        <li><strong>Must</strong> start with a lower case letter</li>
        <li>Composed with lower case letters, digits, undercores, or dashes</li>
        <li><strong>May</strong> end with a dollar sign</li>
        <li>Conform to the following regex <code>[a-z_][a-z0-9_-]*[$]?</code></li>
      </ul>
      <br />

      <input type="text" name="id" placeholder="Entity ID" />
    </fieldset><br />

    <fieldset>
      <legend>Set Entity Number</legend>
      <p>Specify a Number for the entity</p>
      <p>The entity must have a number for certain use cases.  If you
        are using NetAuth's nsscache implementation, make sure you are
        using a number here that is above your minimum range.  When in
        doubt, leave this number at -1 and NetAuth will select the
        next available number.  In general, you shouldn't be assigning
        these by hand anyway.</p>

      <input type="number" name="number" value="-1"
      placeholder="Number" />
    </fieldset><br />

    <fieldset>
      <legend>Set Entity Secret</legend>
      <p>Specify the entity secret</p>
      <p>This is the authentication secret for the entity.  This can
        be any unicode string that can later be supplied to NetAuth to
        authenticate itself.  While you can provide any string you want
        here, its recommended that you use an ASCII string here.  If you
        need to supply this for a human user, selecting a passphrase
        composed of plain words is a good pick.</p>

      <input type="secret"
             name="secret"
             placeholder="Secret" />
    </fieldset>
    <br />

    <div class="center">
      <button>Create Entity</button>
      <button type="reset">Reset</button>
    </div>
  </form>
</div>
{{ end }}
