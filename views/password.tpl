<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>{{ password.Title }}</title>
  </head>

  <body>
    <h1 class="title">「{{ password.Title }}」</h1>
    <form action="{{ password.UpdateURL }}" method="POST">
      <p>
        Subject:
        <input id="title" name="title" type="text" value="{{ password.Title }}">
      </p>
      <p>
        Password:
        <input id="body" name="body" type="text" value="{{ password.Body }}">
      </p>
      <input type="submit" value="Save">
    </form>
    <hr>
    <a href="/passwords">To Index</a>
  </body>

</html>
