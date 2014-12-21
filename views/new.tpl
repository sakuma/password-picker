<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>New Password</title>
  </head>
  <body>
    <div id="content">
      <form action="/passwords" method="POST">
      <h1 class="title">New Password</h1>
      <p>
        Title:
        <input id="title" name="title" cols="80" rows="20" type="text" value="{{ password.Title }}" />
      </p>
      <p>
        Password:
        <input id="body" name="body" cols="80" rows="20" type="text" value="{{ password.Body }}" />
      </p>
      <p>
        <input type="submit" value="Save">
      </p>
      </form>
    </div>
  </body>
</html>
