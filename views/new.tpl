<!DOCTYPE html>
<html lang="ja">
  <head>
    <meta charset="UTF-8">
    <title>New Password</title>
    <link href="/assets/stylesheets/normalize.css" media="all" rel="stylesheet" />
    <link href="/assets/stylesheets/skeleton.css " media="all" rel="stylesheet" />
  </head>
  <body>
    <div class="container">

      <h1 class="title">New Password</h1>
      <form action="/passwords" method="POST">

        <div class="row">
          <div class="two columns">
            <label for="title"> Title: </label>
          </div>
          <input id="title" name="title" cols="80" rows="20" type="text" value="{{ password.Title }}" />
        </div>

        <div class="row">
          <div class="two columns">
            <label for="body"> Password: </label>
          </div>
          <input id="body" name="body" cols="80" rows="20" type="text" value="{{ password.Body }}" />
        </div>

        <div class="row">
          <div class="two columns">&nbsp;</div>
          <input type="submit" value="Save">
        </div>

      </form>

    </div>
  </body>

</html>
